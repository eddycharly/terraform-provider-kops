package schemas

import (
	"time"

	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

type Expander func(interface{}) interface{}

func ExpandPtr(next Expander) Expander {
	return func(in interface{}) interface{} {
		if in == nil {
			return nil
		}
		ret := next(in)
		return &ret
	}
}

func ExpandInt() Expander {
	return func(in interface{}) interface{} {
		return in.(int)
	}
}

func ExpandString() Expander {
	return func(in interface{}) interface{} {
		return in.(string)
	}
}

func ExpandBool() Expander {
	return func(in interface{}) interface{} {
		return in.(bool)
	}
}

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

func ExpandStruct(next func(map[string]interface{}) interface{}) Expander {
	return func(in interface{}) interface{} {
		return next(in.(map[string]interface{}))
	}
}

func ExpandList(next Expander) Expander {
	return func(in interface{}) interface{} {
		if in == nil {
			return nil
		}
		out := []interface{}{}
		for _, in := range in.([]interface{}) {
			out = append(out, next(in))
		}
		return out
	}
}
