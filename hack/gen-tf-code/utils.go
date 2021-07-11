package main

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"
)

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

func toSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

func fieldName(in string) string {
	in = strings.ReplaceAll(in, "EBS", "Ebs")
	in = strings.ReplaceAll(in, "CSI", "Csi")
	in = strings.ReplaceAll(in, "CIDR", "Cidr")
	in = strings.ReplaceAll(in, "DNS", "Dns")
	in = strings.ReplaceAll(in, "IP", "Ip")
	in = strings.ReplaceAll(in, "SSH", "Ssh")
	in = strings.ReplaceAll(in, "API", "Api")
	in = strings.ReplaceAll(in, "SAN", "San")
	in = strings.ReplaceAll(in, "SQS", "Sqs")
	in = strings.ReplaceAll(in, "AWS", "Aws")
	return in
}

func schemaType(in reflect.Type) string {
	switch in.Kind() {
	case reflect.String:
		return "String"
	case reflect.Bool:
		return "Bool"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return "Int"
	case reflect.Float32, reflect.Float64:
		return "Float"
	default:
		panic(fmt.Sprintf("unknown kind %v", in.Kind()))
	}
}
