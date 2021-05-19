module github.com/eddycharly/terraform-provider-kops

go 1.14

// Version kubernetes-1.19.0 => tag v0.19.5

replace k8s.io/api => k8s.io/api v0.19.5

replace k8s.io/apimachinery => k8s.io/apimachinery v0.19.5

replace k8s.io/client-go => k8s.io/client-go v0.19.5

replace k8s.io/cloud-provider => k8s.io/cloud-provider v0.19.5

replace k8s.io/kubectl => k8s.io/kubectl v0.19.5

replace k8s.io/apiserver => k8s.io/apiserver v0.19.5

replace k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.19.5

replace k8s.io/kube-scheduler => k8s.io/kube-scheduler v0.19.5

replace k8s.io/kube-proxy => k8s.io/kube-proxy v0.19.5

replace k8s.io/cri-api => k8s.io/cri-api v0.19.5

replace k8s.io/csi-translation-lib => k8s.io/csi-translation-lib v0.19.5

replace k8s.io/legacy-cloud-providers => k8s.io/legacy-cloud-providers v0.19.5

replace k8s.io/component-base => k8s.io/component-base v0.19.5

replace k8s.io/cluster-bootstrap => k8s.io/cluster-bootstrap v0.19.5

replace k8s.io/metrics => k8s.io/metrics v0.19.5

replace k8s.io/sample-apiserver => k8s.io/sample-apiserver v0.19.5

replace k8s.io/kube-aggregator => k8s.io/kube-aggregator v0.19.5

replace k8s.io/kubelet => k8s.io/kubelet v0.19.5

replace k8s.io/cli-runtime => k8s.io/cli-runtime v0.19.5

replace k8s.io/kube-controller-manager => k8s.io/kube-controller-manager v0.19.5

replace k8s.io/code-generator => k8s.io/code-generator v0.19.5

replace github.com/gophercloud/gophercloud => github.com/gophercloud/gophercloud v0.11.0

// replace golang.org/x/sys => golang.org/x/sys v0.0.0-20200826173525-f9321e4c35a6

require (
	github.com/Masterminds/goutils v1.1.0 // indirect
	github.com/Masterminds/sprig v2.22.0+incompatible
	github.com/aws/aws-sdk-go v1.38.42
	github.com/hashicorp/terraform-plugin-sdk/v2 v2.5.0
	golang.org/x/tools v0.1.0
	google.golang.org/grpc/examples v0.0.0-20201121004645-9da74c039bbf // indirect
	k8s.io/api v0.20.4
	k8s.io/apimachinery v0.21.0
	k8s.io/client-go v1.5.2
	k8s.io/klog v1.0.0
	k8s.io/kops v1.19.3
)
