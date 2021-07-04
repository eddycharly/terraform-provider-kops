package schemas

import (
	"time"

	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

func ExpandQuantity(in string) resource.Quantity {
	q, _ := resource.ParseQuantity(in)
	return q
}

func ExpandDuration(in string) metav1.Duration {
	d, _ := time.ParseDuration(in)
	return metav1.Duration{
		Duration: d,
	}
}

func ExpandIntOrString(in string) intstr.IntOrString {
	return intstr.Parse(in)
}
