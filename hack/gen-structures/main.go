package main

import (
	"fmt"
	"os"
	"reflect"
	"strings"
	"text/template"

	"github.com/Masterminds/sprig"
	"github.com/eddycharly/terraform-provider-kops/pkg/api"
	"k8s.io/kops/pkg/apis/kops"
)

func schemaType(t reflect.Type) string {
	switch t.Kind() {
	case reflect.String:
		return "String"
	case reflect.Bool:
		return "Bool"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return "Int"
	case reflect.Float32, reflect.Float64:
		return "Float"
	default:
		panic(fmt.Sprintf("unknown kind %v", t.Kind()))
	}
}

func buildStructure(t reflect.Type, exclude ...string) {
	tmplString := `
package {{ .Package }}

import (
	"encoding/json"
	"log"
	"strings"
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/kops/pkg/apis/kops"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/eddycharly/terraform-provider-kops/pkg/api"
)

{{- define "expandElem" -}}
{{- if isPtr . -}}
func (in interface{}) {{ .String }} {
	if in == nil {
		return nil
	}
	if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
		return nil
	}
	tmp := func (in {{ .Elem.String }}) {{ .String }} {
		// if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
		// 	return nil
		// }
		return &in
	}({{ template "expandElem" .Elem }})
	return tmp
}(in)
{{- else if isDuration . -}}
ExpandDuration(in)
{{- else if isQuantity . -}}
ExpandQuantity(in)
{{- else if isIntOrString . -}}
ExpandIntOrString(in)
{{- else if isStruct . -}}
func (in interface{}) {{ .String }} {
	if in == nil {
		return {{ .String }}{}
	}
	return (Expand{{ .Name }}(in.(map[string]interface{})))
}(in)
{{- else -}}
{{ template "expand" . }}
{{- end -}}
{{- end -}}

{{- define "expand" -}}
{{- if isPtr . -}}
func (in interface{}) {{ .String }} {
	if in == nil {
		return nil
	}
	if slice, ok := in.([]interface{}); ok && len(slice) == 0 {
		return nil
	}
	tmp := func (in {{ .Elem.String }}) {{ .String }} {
		// if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
		// 	return nil
		// }
		return &in
	}({{ template "expand" .Elem }})
	return tmp
}(in)
{{- else if isList . -}}
func (in interface{}) {{ .String }} {
	var out {{ .String }}
	for _, in := range in.([]interface{}) {
		out = append(out, {{ template "expandElem" .Elem }})
	}
	return out
}(in)
{{- else if isMap . -}}
func (in interface{}) map[string]{{ .Elem.String }} {
	if in == nil {
		return nil
	}
	out := {{ .String }}{}
	for key, in := range in.(map[string]interface{}) {
		out[key] = {{ template "expand" .Elem }}
	}
	return out
}(in)
{{- else if isDuration . -}}
ExpandDuration(in)
{{- else if isQuantity . -}}
ExpandQuantity(in)
{{- else if isIntOrString . -}}
ExpandIntOrString(in)
{{- else if isStruct . -}}
func (in interface{}) {{ .String }} {
	if in.([]interface{})[0] == nil {
		return {{ .String }}{}
	}
	return (Expand{{ .Name }}(in.([]interface{})[0].(map[string]interface{})))
}(in)
{{- else -}}
{{ .String }}(Expand{{ schemaType . }}(in))
{{- end -}}
{{- end -}}

{{- define "flattenElem" -}}
{{- if isPtr . -}}
func (in {{ .String }}) interface{} {
	if in == nil {
		return nil
	}
	return func (in {{ .Elem.String }}) interface{} {
		return {{ template "flattenElem" .Elem }}
	}(*in)
}(in)
{{- else if isDuration . -}}
FlattenDuration(in)
{{- else if isQuantity . -}}
FlattenQuantity(in)
{{- else if isIntOrString . -}}
FlattenIntOrString(in)
{{- else if isStruct . -}}
func (in {{ .String }}) interface{} {
	return Flatten{{ .Name }}(in)
}(in)
{{- else -}}
{{ template "flatten" . }}
{{- end -}}
{{- end -}}

{{- define "flatten" -}}
{{- if isPtr . -}}
func (in {{ .String }}) interface{} {
	if in == nil {
		return nil
	}
	return func (in {{ .Elem.String }}) interface{} {
		return {{ template "flatten" .Elem }}
	}(*in)
}(in)
{{- else if isList . -}}
func (in {{ .String }}) []interface{} {
	var out []interface{}
	for _, in := range in {
		out = append(out, {{ template "flattenElem" .Elem }})
	}
	return out
}(in)
{{- else if isMap . -}}
func (in map[string]{{ .Elem.String }}) map[string]interface{} {
	if in == nil {
		return nil
	}
	// TODO
	return nil
}(in)
{{- else if isDuration . -}}
FlattenDuration(in)
{{- else if isQuantity . -}}
FlattenQuantity(in)
{{- else if isIntOrString . -}}
FlattenIntOrString(in)
{{- else if isStruct . -}}
func (in {{ .String }}) []map[string]interface{} {
	return []map[string]interface{}{Flatten{{ .Name }}(in)}
}(in)
{{- else -}}
Flatten{{ schemaType . }}({{ schemaType . | lower }}(in))
{{- end -}}
{{- end -}}

{{- with .Type }}

func Expand{{ .Name }}(in map[string]interface{}) {{ .String }} {
	if in == nil {
		panic("expand {{ .Name }} failure, in is nil")
	}
	return {{ .String }}{
	{{- range (fields .) }}
	{{- if not (has .Name $.Exclude) }}
	{{ .Name }}: func (in interface{}) {{ .Type.String }} {
		value := {{ template "expand" .Type }}
		return value
	}(in[{{ fieldName .Name | snakecase | quote }}]),
	{{- end }}
	{{- end }}
	}
}

func Flatten{{ .Name }}(in {{ .String }}) map[string]interface{} {
	return map[string]interface{}{
	{{- range (fields .) }}
	{{- if not (has .Name $.Exclude) }}
	{{ fieldName .Name | snakecase | quote }}: func (in {{ .Type.String }}) interface{} {
		value := {{ template "flatten" .Type }}
		return value
	}(in.{{ .Name }}),
	{{- end }}
	{{- end }}
	}
}

{{- end }}
`

	tmpl := template.New("doc").Funcs(template.FuncMap{
		"fields": func(t reflect.Type) []reflect.StructField {
			var ret []reflect.StructField
			for i := 0; i < t.NumField(); i++ {
				ret = append(ret, t.Field(i))
			}
			return ret
		},
		"schemaType": schemaType,
		"fieldName": func(in string) string {
			in = strings.ReplaceAll(in, "CIDR", "Cidr")
			in = strings.ReplaceAll(in, "DNS", "Dns")
			in = strings.ReplaceAll(in, "IP", "Ip")
			return in
		},
		"isPtr": func(t reflect.Type) bool {
			return t.Kind() == reflect.Ptr
		},
		"isList": func(t reflect.Type) bool {
			return t.Kind() == reflect.Slice
		},
		"isStruct": func(t reflect.Type) bool {
			return t.Kind() == reflect.Struct
		},
		"isMap": func(t reflect.Type) bool {
			return t.Kind() == reflect.Map
		},
		"isDuration": func(t reflect.Type) bool {
			return t.Kind() == reflect.Struct && t.String() == "v1.Duration"
		},
		"isQuantity": func(t reflect.Type) bool {
			return t.Kind() == reflect.Struct && t.String() == "resource.Quantity"
		},
		"isIntOrString": func(t reflect.Type) bool {
			return t.Kind() == reflect.Struct && t.String() == "intstr.IntOrString"
		},
	}).Funcs(sprig.TxtFuncMap())

	if _, err := tmpl.Parse(tmplString); err != nil {
		panic(err)
	}

	f, _ := os.Create("pkg/structures/" + t.Name() + ".generated.go")
	defer f.Close()

	if err := tmpl.Execute(f, map[string]interface{}{
		"Package": "structures",
		"Type":    t,
		"Exclude": exclude,
	}); err != nil {
		panic(err)
	}
}

func build(t reflect.Type, exclude ...string) {
	buildStructure(t, exclude...)
}

func main() {
	build(reflect.TypeOf(api.Cluster{}))
	build(reflect.TypeOf(api.InstanceGroup{}))
	build(reflect.TypeOf(kops.AccessSpec{}))
	build(reflect.TypeOf(kops.DNSAccessSpec{}))
	build(reflect.TypeOf(kops.LoadBalancerAccessSpec{}))
	build(reflect.TypeOf(kops.EtcdClusterSpec{}))
	build(reflect.TypeOf(kops.EtcdBackupSpec{}))
	build(reflect.TypeOf(kops.EtcdManagerSpec{}))
	build(reflect.TypeOf(kops.EtcdMemberSpec{}))
	build(reflect.TypeOf(kops.EnvVar{}))
	build(reflect.TypeOf(kops.ClusterSubnetSpec{}))
	build(reflect.TypeOf(kops.TopologySpec{}))
	build(reflect.TypeOf(kops.BastionSpec{}))
	build(reflect.TypeOf(kops.BastionLoadBalancerSpec{}))
	build(reflect.TypeOf(kops.DNSSpec{}))
	build(reflect.TypeOf(kops.NetworkingSpec{}))
	build(reflect.TypeOf(kops.ClassicNetworkingSpec{}))
	build(reflect.TypeOf(kops.KubenetNetworkingSpec{}))
	build(reflect.TypeOf(kops.ExternalNetworkingSpec{}))
	build(reflect.TypeOf(kops.CNINetworkingSpec{}))
	build(reflect.TypeOf(kops.KopeioNetworkingSpec{}))
	build(reflect.TypeOf(kops.WeaveNetworkingSpec{}))
	build(reflect.TypeOf(kops.FlannelNetworkingSpec{}))
	build(reflect.TypeOf(kops.CalicoNetworkingSpec{}))
	build(reflect.TypeOf(kops.CanalNetworkingSpec{}))
	build(reflect.TypeOf(kops.KuberouterNetworkingSpec{}))
	build(reflect.TypeOf(kops.RomanaNetworkingSpec{}))
	build(reflect.TypeOf(kops.AmazonVPCNetworkingSpec{}))
	build(reflect.TypeOf(kops.CiliumNetworkingSpec{}))
	build(reflect.TypeOf(kops.LyftVPCNetworkingSpec{}))
	build(reflect.TypeOf(kops.GCENetworkingSpec{}))
	build(reflect.TypeOf(kops.VolumeSpec{}))
	build(reflect.TypeOf(kops.VolumeMountSpec{}))
	build(reflect.TypeOf(kops.MixedInstancesPolicySpec{}))
	build(reflect.TypeOf(kops.UserData{}))
	build(reflect.TypeOf(kops.LoadBalancer{}))
	build(reflect.TypeOf(kops.IAMProfileSpec{}))
	build(reflect.TypeOf(kops.HookSpec{}))
	build(reflect.TypeOf(kops.ExecContainerAction{}))
	build(reflect.TypeOf(kops.FileAssetSpec{}))
	build(reflect.TypeOf(kops.KubeletConfigSpec{}))
	build(reflect.TypeOf(kops.RollingUpdate{}))
}
