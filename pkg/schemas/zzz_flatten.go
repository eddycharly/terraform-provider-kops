package schemas

import (
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

func FlattenQuantity(in resource.Quantity) string {
	return in.String()
}

func FlattenDuration(in metav1.Duration) string {
	return in.String()
}

func FlattenIntOrString(in intstr.IntOrString) string {
	return in.String()
}
