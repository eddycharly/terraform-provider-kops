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

func isSlice(t reflect.Type) bool {
	return t.Kind() == reflect.Slice
}

func isStruct(t reflect.Type) bool {
	return t.Kind() == reflect.Struct
}

func isMap(t reflect.Type) bool {
	return t.Kind() == reflect.Map
}

func isDuration(t reflect.Type) bool {
	return t.Kind() == reflect.Struct && t.String() == "v1.Duration"
}

func isQuantity(t reflect.Type) bool {
	return t.Kind() == reflect.Struct && t.String() == "resource.Quantity"
}

func isIntOrString(t reflect.Type) bool {
	return t.Kind() == reflect.Struct && t.String() == "intstr.IntOrString"
}

func isPtr(t reflect.Type) bool {
	return t.Kind() == reflect.Ptr
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
