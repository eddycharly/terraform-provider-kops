package main

import (
	"fmt"
	"os"
	"reflect"
	"strings"
)

func readFile(path string) string {
	if out, err := os.ReadFile(path); err != nil {
		panic(err)
	} else {
		return string(out)
	}
}

var nullableSection = readFile("hack/gen-tf-code/docs/nullable-section.md")

func readHeader(path string, nullable bool) string {
	if !nullable {
		return readFile(path)
	}
	return fmt.Sprintf("%s\n\n%s", readFile(path), nullableSection)
}

var (
	resourceClusterHeader        = readHeader("hack/gen-tf-code/docs/resource-cluster-header.md", true)
	resourceClusterFooter        = readFile("hack/gen-tf-code/docs/resource-cluster-footer.md")
	resourceClusterUpdaterHeader = readHeader("hack/gen-tf-code/docs/resource-cluster-updater-header.md", false)
	resourceInstanceGroupHeader  = readHeader("hack/gen-tf-code/docs/resource-instance-group-header.md", true)
	resourceInstanceGroupFooter  = readFile("hack/gen-tf-code/docs/resource-instance-group-footer.md")
	dataClusterHeader            = readHeader("hack/gen-tf-code/docs/data-cluster-header.md", true)
	dataClusterStatusHeader      = readHeader("hack/gen-tf-code/docs/data-cluster-status-header.md", false)
	dataInstanceGroupHeader      = readHeader("hack/gen-tf-code/docs/data-instance-group-header.md", true)
	dataKubeConfigHeader         = readHeader("hack/gen-tf-code/docs/data-kube-config-header.md", false)
	configProviderHeader         = readHeader("hack/gen-tf-code/docs/config-provider-header.md", true)
)

func getSubResources(t reflect.Type, seen map[reflect.Type]bool, isExcluded func(in _field) bool) []reflect.Type {
	if t.Kind() == reflect.Array || t.Kind() == reflect.Map || t.Kind() == reflect.Slice || t.Kind() == reflect.Ptr {
		return getSubResources(t.Elem(), seen, isExcluded)
	}
	if t.Kind() != reflect.Struct {
		return nil
	}
	if isValueType(t) {
		return nil
	}
	if _, ok := seen[t]; ok {
		return nil
	}
	seen[t] = true
	ret := []reflect.Type{t}
	for _, f := range getFields(t, true) {
		if !isExcluded(f) {
			ret = append(ret, getSubResources(f.Type, seen, isExcluded)...)
		}
	}
	return ret
}

func getResourceComment(packName, structName string, c map[string]map[string]_struct) string {
	if p, ok := c[packName]; ok {
		if s, ok := p[structName]; ok {
			ret := strings.ReplaceAll(strings.TrimSpace(s.doc), "\n", "<br />")
			if ret != "" {
				if !strings.HasSuffix(ret, ".") {
					ret = ret + "."
				}
				return ret
			}
		}
	}
	return ""
}

func getAttributeComment(packName, structName, fieldName string, c map[string]map[string]_struct) string {
	if p, ok := c[packName]; ok {
		if s, ok := p[structName]; ok {
			ret := strings.ReplaceAll(strings.TrimSpace(s.lookup(fieldName, c)), "\n", "<br />")
			if ret != "" {
				if !strings.HasSuffix(ret, ".") {
					ret = ret + "."
				}
				return ret
			}
		}
	}
	return ""
}
