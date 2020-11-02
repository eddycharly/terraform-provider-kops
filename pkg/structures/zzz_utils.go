package structures

import (
	"k8s.io/apimachinery/pkg/api/resource"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

// IntOrString

func ExpandIntOrString(in interface{}) intstr.IntOrString {
	// TODO
	return intstr.IntOrString{}
}

func FlattenIntOrString(in intstr.IntOrString) interface{} {
	// TODO
	return nil
}

// Duration

func ExpandDuration(in interface{}) v1.Duration {
	// TODO
	return v1.Duration{}
}

func FlattenDuration(in v1.Duration) interface{} {
	// TODO
	return nil
}

// Quantity

func ExpandQuantity(in interface{}) resource.Quantity {
	// TODO
	return resource.Quantity{}
}

func FlattenQuantity(in resource.Quantity) interface{} {
	// TODO
	return nil
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
