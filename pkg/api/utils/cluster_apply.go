package utils

import (
	"context"

	"k8s.io/kops/pkg/client/simple"
	"k8s.io/kops/upup/pkg/fi/cloudup"
)

func ClusterApply(clientset simple.Clientset, clusterName string, allowKopsDowngrade bool) error {
	kc, err := clientset.GetCluster(context.Background(), clusterName)
	if err != nil {
		return err
	}
	cloud, err := cloudup.BuildCloud(kc)
	if err != nil {
		return err
	}
	apply := &cloudup.ApplyClusterCmd{
		Cloud:              cloud,
		Cluster:            kc,
		Clientset:          clientset,
		TargetName:         cloudup.TargetDirect,
		AllowKopsDowngrade: allowKopsDowngrade,
	}
	return apply.Run(context.Background())
}
