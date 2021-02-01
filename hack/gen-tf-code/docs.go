package main

import "io/ioutil"

func readFile(path string) string {
	if out, err := ioutil.ReadFile(path); err != nil {
		panic(err)
	} else {
		return string(out)
	}
}

const genericHeader = `
{{- $comment := resourceComment . }}
{{ if $comment }}
{{ $comment }}
{{- end }}
`

var (
	resourceClusterHeader        = readFile("hack/gen-tf-code/resource-cluster-header.md")
	resourceClusterFooter        = readFile("hack/gen-tf-code/resource-cluster-footer.md")
	resourceClusterUpdaterHeader = readFile("hack/gen-tf-code/resource-cluster-updater-header.md")
	resourceInstanceGroupHeader  = readFile("hack/gen-tf-code/resource-instance-group-header.md")
	resourceInstanceGroupFooter  = readFile("hack/gen-tf-code/resource-instance-group-footer.md")
	dataClusterHeader            = readFile("hack/gen-tf-code/data-cluster-header.md")
	dataClusterStatusHeader      = readFile("hack/gen-tf-code/data-cluster-status-header.md")
	dataInstanceGroupHeader      = readFile("hack/gen-tf-code/data-instance-group-header.md")
	dataKubeConfigHeader         = readFile("hack/gen-tf-code/data-kube-config-header.md")
	configProviderHeader         = readFile("hack/gen-tf-code/config-provider-header.md")
)
