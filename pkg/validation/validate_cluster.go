/*
Copyright 2019 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package validation

import (
	"context"
	"fmt"
	"net"
	"net/url"
	"strings"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/pager"
	"k8s.io/kops/upup/pkg/fi"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/klog"
	"k8s.io/kops/pkg/apis/kops"
	"k8s.io/kops/pkg/cloudinstances"
	"k8s.io/kops/pkg/dns"
	kopsValidation "k8s.io/kops/pkg/validation"
)

type clusterValidatorImpl struct {
	cluster        *kops.Cluster
	cloud          fi.Cloud
	instanceGroups []*kops.InstanceGroup
	k8sClient      kubernetes.Interface
	config         *rest.Config
}

func addError(v *kopsValidation.ValidationCluster, failure *kopsValidation.ValidationError) {
	v.Failures = append(v.Failures, failure)
}

// hasPlaceHolderIP checks if the API DNS has been updated.
func hasPlaceHolderIP(config *rest.Config) (bool, error) {
	apiAddr, err := url.Parse(config.Host)
	if err != nil {
		return true, fmt.Errorf("unable to parse Kubernetes cluster API URL: %v", err)
	}
	hostAddrs, err := net.LookupHost(apiAddr.Hostname())
	if err != nil {
		return true, fmt.Errorf("unable to resolve Kubernetes cluster API URL dns: %v", err)
	}

	for _, h := range hostAddrs {
		if h == "203.0.113.123" {
			return true, nil
		}
	}

	return false, nil
}

func NewClusterValidator(cluster *kops.Cluster, cloud fi.Cloud, instanceGroupList *kops.InstanceGroupList, config *rest.Config, k8sClient kubernetes.Interface) (kopsValidation.ClusterValidator, error) {
	var instanceGroups []*kops.InstanceGroup

	for i := range instanceGroupList.Items {
		ig := &instanceGroupList.Items[i]
		instanceGroups = append(instanceGroups, ig)
	}

	if len(instanceGroups) == 0 {
		return nil, fmt.Errorf("no InstanceGroup objects found")
	}

	return &clusterValidatorImpl{
		cluster:        cluster,
		cloud:          cloud,
		instanceGroups: instanceGroups,
		k8sClient:      k8sClient,
		config:         config,
	}, nil
}

func (v *clusterValidatorImpl) Validate() (*kopsValidation.ValidationCluster, error) {
	ctx := context.TODO()

	clusterName := v.cluster.Name

	validation := &kopsValidation.ValidationCluster{}

	// Do not use if we are running gossip
	if !dns.IsGossipHostname(clusterName) {
		hasPlaceHolderIPAddress, err := hasPlaceHolderIP(v.config)
		if err != nil {
			return nil, err
		}

		if hasPlaceHolderIPAddress {
			message := "Validation Failed\n\n" +
				"The dns-controller Kubernetes deployment has not updated the Kubernetes cluster's API DNS entry to the correct IP address." +
				"  The API DNS IP address is the placeholder address that kops creates: 203.0.113.123." +
				"  Please wait about 5-10 minutes for a master to start, dns-controller to launch, and DNS to propagate." +
				"  The protokube container and dns-controller deployment logs may contain more diagnostic information." +
				"  Etcd and the API DNS entries must be updated for a kops Kubernetes cluster to start."
			addError(validation, &kopsValidation.ValidationError{
				Kind:    "dns",
				Name:    "apiserver",
				Message: message,
			})
			return validation, nil
		}
	}

	nodeList, err := v.k8sClient.CoreV1().Nodes().List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, fmt.Errorf("error listing nodes: %v", err)
	}

	warnUnmatched := false
	cloudGroups, err := v.cloud.GetCloudGroups(v.cluster, v.instanceGroups, warnUnmatched, nodeList.Items)
	if err != nil {
		return nil, err
	}
	readyNodes := validateNodes(validation, cloudGroups, v.instanceGroups)

	if err := collectComponentFailures(ctx, validation, v.k8sClient); err != nil {
		return nil, fmt.Errorf("cannot get component status for %q: %v", clusterName, err)
	}

	if err := collectPodFailures(ctx, validation, v.k8sClient, readyNodes); err != nil {
		return nil, fmt.Errorf("cannot get pod health for %q: %v", clusterName, err)
	}

	return validation, nil
}

func collectComponentFailures(ctx context.Context, v *kopsValidation.ValidationCluster, client kubernetes.Interface) error {
	componentList, err := client.CoreV1().ComponentStatuses().List(ctx, metav1.ListOptions{})
	if err != nil {
		return fmt.Errorf("error listing ComponentStatuses: %v", err)
	}

	for _, component := range componentList.Items {
		for _, condition := range component.Conditions {
			if condition.Status != v1.ConditionTrue {
				addError(v, &kopsValidation.ValidationError{
					Kind:    "ComponentStatus",
					Name:    component.Name,
					Message: fmt.Sprintf("component %q is unhealthy", component.Name),
				})
			}
		}
	}
	return nil
}

var masterStaticPods = []string{
	"kube-apiserver",
	"kube-controller-manager",
	"kube-scheduler",
}

func collectPodFailures(ctx context.Context, v *kopsValidation.ValidationCluster, client kubernetes.Interface, nodes []v1.Node) error {
	masterWithoutPod := map[string]map[string]bool{}
	nodeByAddress := map[string]string{}
	for _, node := range nodes {
		labels := node.GetLabels()
		if labels != nil && labels["kubernetes.io/role"] == "master" {
			masterWithoutPod[node.Name] = map[string]bool{}
			for _, pod := range masterStaticPods {
				masterWithoutPod[node.Name][pod] = true
			}
		}
		for _, nodeAddress := range node.Status.Addresses {
			nodeByAddress[nodeAddress.Address] = node.Name
		}
	}

	err := pager.New(pager.SimplePageFunc(func(opts metav1.ListOptions) (runtime.Object, error) {
		return client.CoreV1().Pods(metav1.NamespaceAll).List(ctx, opts)
	})).EachListItem(context.TODO(), metav1.ListOptions{}, func(obj runtime.Object) error {
		pod := obj.(*v1.Pod)

		app := pod.GetLabels()["k8s-app"]
		if pod.Namespace == "kube-system" && masterWithoutPod[nodeByAddress[pod.Status.HostIP]][app] {
			delete(masterWithoutPod[nodeByAddress[pod.Status.HostIP]], app)
		}

		priority := pod.Spec.PriorityClassName
		if priority != "system-cluster-critical" && priority != "system-node-critical" {
			return nil
		}
		if pod.Status.Phase == v1.PodSucceeded {
			return nil
		}
		if pod.Status.Phase == v1.PodPending {
			addError(v, &kopsValidation.ValidationError{
				Kind:    "Pod",
				Name:    pod.Namespace + "/" + pod.Name,
				Message: fmt.Sprintf("%s pod %q is pending", priority, pod.Name),
			})
			return nil
		}
		if pod.Status.Phase == v1.PodUnknown {
			addError(v, &kopsValidation.ValidationError{
				Kind:    "Pod",
				Name:    pod.Namespace + "/" + pod.Name,
				Message: fmt.Sprintf("%s pod %q is unknown phase", priority, pod.Name),
			})
			return nil
		}
		var notready []string
		for _, container := range pod.Status.ContainerStatuses {
			if !container.Ready {
				notready = append(notready, container.Name)
			}
		}
		if len(notready) != 0 {
			addError(v, &kopsValidation.ValidationError{
				Kind:    "Pod",
				Name:    pod.Namespace + "/" + pod.Name,
				Message: fmt.Sprintf("%s pod %q is not ready (%s)", priority, pod.Name, strings.Join(notready, ",")),
			})

		}
		return nil
	})
	if err != nil {
		return fmt.Errorf("error listing Pods: %v", err)
	}

	for node, nodeMap := range masterWithoutPod {
		for app := range nodeMap {
			addError(v, &kopsValidation.ValidationError{
				Kind:    "Node",
				Name:    node,
				Message: fmt.Sprintf("master %q is missing %s pod", node, app),
			})
		}
	}

	return nil
}

func validateNodes(v *kopsValidation.ValidationCluster, cloudGroups map[string]*cloudinstances.CloudInstanceGroup, groups []*kops.InstanceGroup) []v1.Node {
	var readyNodes []v1.Node
	groupsSeen := map[string]bool{}

	for _, cloudGroup := range cloudGroups {
		var allMembers []*cloudinstances.CloudInstance
		allMembers = append(allMembers, cloudGroup.Ready...)
		allMembers = append(allMembers, cloudGroup.NeedUpdate...)

		groupsSeen[cloudGroup.InstanceGroup.Name] = true
		numNodes := 0
		for _, m := range allMembers {
			if m.Status != cloudinstances.CloudInstanceStatusDetached {
				numNodes++
			}
		}
		if numNodes < cloudGroup.TargetSize {
			addError(v, &kopsValidation.ValidationError{
				Kind: "InstanceGroup",
				Name: cloudGroup.InstanceGroup.Name,
				Message: fmt.Sprintf("InstanceGroup %q did not have enough nodes %d vs %d",
					cloudGroup.InstanceGroup.Name,
					numNodes,
					cloudGroup.TargetSize),
			})
		}

		for _, member := range allMembers {
			node := member.Node

			if node == nil {
				nodeExpectedToJoin := true
				if cloudGroup.InstanceGroup.Spec.Role == kops.InstanceGroupRoleBastion {
					// bastion nodes don't join the cluster
					nodeExpectedToJoin = false
				}

				if nodeExpectedToJoin {
					addError(v, &kopsValidation.ValidationError{
						Kind:    "Machine",
						Name:    member.ID,
						Message: fmt.Sprintf("machine %q has not yet joined cluster", member.ID),
					})
				}
				continue
			}

			role := strings.ToLower(string(cloudGroup.InstanceGroup.Spec.Role))
			if role == "" {
				role = "node"
			}

			n := &kopsValidation.ValidationNode{
				Name:     node.Name,
				Zone:     node.ObjectMeta.Labels["failure-domain.beta.kubernetes.io/zone"],
				Hostname: node.ObjectMeta.Labels["kubernetes.io/hostname"],
				Role:     role,
				Status:   getNodeReadyStatus(node),
			}

			ready := isNodeReady(node)
			if ready {
				readyNodes = append(readyNodes, *node)
			}

			if n.Role == "master" {
				if !ready {
					addError(v, &kopsValidation.ValidationError{
						Kind:    "Node",
						Name:    node.Name,
						Message: fmt.Sprintf("master %q is not ready", node.Name),
					})
				}

				v.Nodes = append(v.Nodes, n)
			} else if n.Role == "node" {
				if !ready {
					addError(v, &kopsValidation.ValidationError{
						Kind:    "Node",
						Name:    node.Name,
						Message: fmt.Sprintf("node %q is not ready", node.Name),
					})
				}

				v.Nodes = append(v.Nodes, n)
			} else {
				klog.Warningf("ignoring node with role %q", n.Role)
			}
		}
	}

	for _, ig := range groups {
		if !groupsSeen[ig.Name] {
			addError(v, &kopsValidation.ValidationError{
				Kind:    "InstanceGroup",
				Name:    ig.Name,
				Message: fmt.Sprintf("InstanceGroup %q is missing from the cloud provider", ig.Name),
			})
		}
	}

	return readyNodes
}
