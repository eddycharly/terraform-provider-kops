package main

import (
	"fmt"
	"reflect"
)

func isBool(in reflect.Type) bool {
	return in.Kind() == reflect.Bool
}

func isInt(in reflect.Type) bool {
	switch in.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return true
	}
	return false
}

func isString(in reflect.Type) bool {
	return in.Kind() == reflect.String
}

func isFloat(in reflect.Type) bool {
	switch in.Kind() {
	case reflect.Float32, reflect.Float64:
		return true
	}
	return false
}

func isSlice(in reflect.Type) bool {
	return in.Kind() == reflect.Slice
}

func isStruct(in reflect.Type) bool {
	return in.Kind() == reflect.Struct
}

func isMap(in reflect.Type) bool {
	return in.Kind() == reflect.Map
}

func isDuration(in reflect.Type) bool {
	return in.Kind() == reflect.Struct && in.String() == "v1.Duration"
}

func isQuantity(in reflect.Type) bool {
	return in.Kind() == reflect.Struct && in.String() == "resource.Quantity"
}

func isIntOrString(in reflect.Type) bool {
	return in.Kind() == reflect.Struct && in.String() == "intstr.IntOrString"
}

func isPtr(in reflect.Type) bool {
	return in.Kind() == reflect.Ptr
}

func isValueType(in reflect.Type) bool {
	switch in.Kind() {
	case reflect.Ptr:
		return isValueType(in.Elem())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Float32, reflect.Float64, reflect.String, reflect.Bool:
		return true
	case reflect.Slice, reflect.Map:
		return false
	case reflect.Struct:
		return in.String() == "v1.Duration" || in.String() == "resource.Quantity" || in.String() == "intstr.IntOrString"
	default:
		panic(fmt.Sprintf("unknown kind %v", in.Kind()))
	}
}

func getFields(t reflect.Type, flatten bool) []_field {
	var ret []_field
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if f.Anonymous && flatten {
			ret = append(ret, getFields(f.Type, flatten)...)
		} else {
			ret = append(ret, _field{
				StructField: t.Field(i),
				Owner:       t,
			})
		}
	}
	return ret
}

func verifyFields(t reflect.Type, fields ...string) error {
	for _, field := range fields {
		valid := false
		for i := 0; i < t.NumField(); i++ {
			f := t.Field(i)
			if f.Name == field {
				valid = true
			}
		}
		if !valid {
			return fmt.Errorf("field %s is not part of struct %s", field, t.Name())
		}
	}
	return nil
}

func resolve(in reflect.Type, m map[string]string) string {
	switch in.Kind() {
	case reflect.Ptr:
		return "*" + resolve(in.Elem(), m)
	case reflect.Slice:
		return "[]" + resolve(in.Elem(), m)
	case reflect.Map:
		return "map[" + resolve(in.Key(), m) + "]" + resolve(in.Elem(), m)
	default:
		p := in.PkgPath()
		if p == "" {
			return in.String()
		} else {
			p, ok := m[p]
			if ok {
				return p + "." + in.Name()
			} else {
				return in.String()
			}
		}
	}
}
