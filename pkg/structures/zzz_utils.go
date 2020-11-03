package structures

import (
	"time"

	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

// IntOrString

func ExpandIntOrString(in interface{}) intstr.IntOrString {
	return intstr.Parse(in.(string))
}

func FlattenIntOrString(in intstr.IntOrString) interface{} {
	return in.String()
}

// Duration

func ExpandDuration(in interface{}) metav1.Duration {
	d, _ := time.ParseDuration(in.(string))
	return metav1.Duration{
		Duration: d,
	}
}

func FlattenDuration(in metav1.Duration) interface{} {
	return in.String()
}

// Quantity

func ExpandQuantity(in interface{}) resource.Quantity {
	q, _ := resource.ParseQuantity(in.(string))
	return q
}

func FlattenQuantity(in resource.Quantity) interface{} {
	return in.String()
}

// Int

func ExpandInt(in interface{}) int {
	return in.(int)
}

func FlattenInt(in int) interface{} {
	return in
}

// String

func ExpandString(in interface{}) string {
	return in.(string)
}

func FlattenString(in string) interface{} {
	return in
}

// Bool

func ExpandBool(in interface{}) bool {
	return in.(bool)
}

func FlattenBool(in bool) interface{} {
	return in
}

// Float

func ExpandFloat(in interface{}) float64 {
	return in.(float64)
}

func FlattenFloat(in float64) interface{} {
	return in
}
