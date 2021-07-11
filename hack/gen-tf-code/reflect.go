package main

import "reflect"

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
