package main

import (
	"fmt"
	"io/ioutil"
)

func readFile(path string) string {
	if out, err := ioutil.ReadFile(path); err != nil {
		panic(err)
	} else {
		return string(out)
	}
}

var nullableSection = readFile("hack/gen-tf-code/nullable-section.md")

func readHeader(path string, nullable bool) string {
	if !nullable {
		return readFile(path)
	}
	return fmt.Sprintf("%s\n\n%s", readFile(path), nullableSection)
}

var (
	resourceClusterHeader        = readHeader("hack/gen-tf-code/resource-cluster-header.md", true)
	resourceClusterFooter        = readFile("hack/gen-tf-code/resource-cluster-footer.md")
	resourceClusterUpdaterHeader = readHeader("hack/gen-tf-code/resource-cluster-updater-header.md", false)
	resourceInstanceGroupHeader  = readHeader("hack/gen-tf-code/resource-instance-group-header.md", true)
	resourceInstanceGroupFooter  = readFile("hack/gen-tf-code/resource-instance-group-footer.md")
	dataClusterHeader            = readHeader("hack/gen-tf-code/data-cluster-header.md", true)
	dataClusterStatusHeader      = readHeader("hack/gen-tf-code/data-cluster-status-header.md", false)
	dataInstanceGroupHeader      = readHeader("hack/gen-tf-code/data-instance-group-header.md", true)
	dataKubeConfigHeader         = readHeader("hack/gen-tf-code/data-kube-config-header.md", false)
	configProviderHeader         = readHeader("hack/gen-tf-code/config-provider-header.md", true)
)
