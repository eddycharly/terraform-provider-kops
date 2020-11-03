package main

import (
	"fmt"
	"os"
	"reflect"
	"strings"
	"text/template"

	"github.com/Masterminds/sprig"
	"github.com/eddycharly/terraform-provider-kops/pkg/api"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/kops/pkg/apis/kops"
)

type options struct {
	exclude  []string
	required []string
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

func fieldName(in string) string {
	in = strings.ReplaceAll(in, "CIDR", "Cidr")
	in = strings.ReplaceAll(in, "DNS", "Dns")
	in = strings.ReplaceAll(in, "IP", "Ip")
	in = strings.ReplaceAll(in, "SSH", "Ssh")
	in = strings.ReplaceAll(in, "API", "Api")
	in = strings.ReplaceAll(in, "SAN", "San")
	return in
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

func buildDoc(t reflect.Type, o options) {
	tmplString := `# kops_{{ fieldName .Type.Name | snakecase }}

{{- define "type" -}}
{{- if isPtr . -}}
{{- template "type" .Elem }}
{{- else if isList . -}}
List({{- template "type" .Elem }})
{{- else if isMap . -}}
Map({{- template "type" .Elem }})
{{- else if isDuration . -}}
Duration
{{- else if isQuantity . -}}
Quantity
{{- else if isIntOrString . -}}
IntOrString
{{- else if isStruct . -}}
[{{ .Name }}](./{{ .Name }}.generated.md)
{{- else -}}
{{- schemaType . -}}
{{- end -}}
{{- end }}

{{- define "required" -}}
{{ if isRequired .Name }}Required{{ else }}Optional{{ end }}
{{- end }}

{{- define "computed" -}}
{{ if isComputed .Name }}Computed{{ end }}
{{- end }}

| attribute | type | optional/required | computed |
| --- | --- | --- | --- |
{{- with .Type -}}
{{- range (fields .) -}}
{{- if not (has .Name $.Exclude) }}
| {{ fieldName .Name | snakecase | code }} | {{ template "type" .Type }} | {{ template "required" . }} | {{ template "computed" . }} |
{{- end -}}
{{- end -}}
{{- end }}
`
	required := sets.NewString(o.required...)

	tmpl := template.New("doc").Funcs(template.FuncMap{
		"fields": func(t reflect.Type) []reflect.StructField {
			var ret []reflect.StructField
			for i := 0; i < t.NumField(); i++ {
				ret = append(ret, t.Field(i))
			}
			return ret
		},
		"schemaType": schemaType,
		"isRequired": func(in string) bool {
			return required.Has(in)
		},
		"isComputed": func(in string) bool {
			// TODO
			return false
		},
		"fieldName": fieldName,
		"code": func(in string) string {
			return fmt.Sprintf("`%s`", in)
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

	f, _ := os.Create("docs/" + t.Name() + ".generated.md")
	defer f.Close()

	if err := tmpl.Execute(f, map[string]interface{}{
		"Type":    t,
		"Exclude": o.exclude,
	}); err != nil {
		panic(err)
	}
}
func buildSchema(t reflect.Type, o options) {
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
)

{{- define "schemaElem" -}}
{{- if isPtr . -}}
{{- template "schemaElem" .Elem }}
{{- else if isStruct . -}}
{{ .Name }}()
{{- else -}}
{{- template "schema" . -}}
{{- end -}}
{{- end -}}

{{- define "schema" -}}
{{- if isPtr . -}}
{{- template "schema" .Elem }}
{{- else if isList . -}}
List({{ template "schemaElem" .Elem }})
{{- else if isMap . -}}
Map({{ template "schemaElem" .Elem }})
{{- else if isDuration . -}}
Duration()
{{- else if isQuantity . -}}
Quantity()
{{- else if isIntOrString . -}}
IntOrString()
{{- else if isStruct . -}}
Struct({{ .Name }}())
{{- else -}}
{{- schemaType . -}}()
{{- end -}}
{{- end -}}

{{- with .Type }}

func {{ .Name }}() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			{{- range (fields .) }}
			{{- if not (has .Name $.Exclude) }}
			{{ fieldName .Name | snakecase | quote }}: {{ schemaModifier .Name }}{{- template "schema" .Type }},
			{{- end }}
			{{- end }}
		},
	}
}

{{- end }}
`
	required := sets.NewString(o.required...)

	tmpl := template.New("doc").Funcs(template.FuncMap{
		"fields": func(t reflect.Type) []reflect.StructField {
			var ret []reflect.StructField
			for i := 0; i < t.NumField(); i++ {
				ret = append(ret, t.Field(i))
			}
			return ret
		},
		"schemaType": schemaType,
		"schemaModifier": func(in string) string {
			out := ""
			if required.Has(in) {
				out += "Required"
			} else {
				out += "Optional"
			}
			return out
		},
		"fieldName": fieldName,
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

	f, _ := os.Create("pkg/schemas/" + t.Name() + ".generated.go")
	defer f.Close()

	if err := tmpl.Execute(f, map[string]interface{}{
		"Package": "schemas",
		"Type":    t,
		"Exclude": o.exclude,
	}); err != nil {
		panic(err)
	}
}

func buildStructure(t reflect.Type, o options) {
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
	return func (in {{ .Elem.String }}) {{ .String }} {
		{{- if isValueType . }}
		if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
			return nil
		}
		{{- end }}
		return &in
	}({{ template "expandElem" .Elem }})
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
	return func (in {{ .Elem.String }}) {{ .String }} {
		{{- if isValueType . }}
		if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
			return nil
		}
		{{- end }}
		return &in
	}({{ template "expand" .Elem }})
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
		"schemaType":  schemaType,
		"fieldName":   fieldName,
		"isValueType": isValueType,
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
		"Exclude": o.exclude,
	}); err != nil {
		panic(err)
	}
}

func required(required ...string) func(o options) options {
	return func(o options) options {
		o.required = append(o.required, required...)
		return o
	}
}

func excluded(excluded ...string) func(o options) options {
	return func(o options) options {
		o.exclude = append(o.exclude, excluded...)
		return o
	}
}

func build(i interface{}, o ...func(o options) options) {
	opts := options{}
	for _, opt := range o {
		opts = opt(opts)
	}
	t := reflect.TypeOf(i)
	buildDoc(t, opts)
	buildStructure(t, opts)
	buildSchema(t, opts)
}

func main() {
	build(api.Cluster{}, required("Name", "CloudProvider", "Subnet", "NetworkID", "Topology", "EtcdCluster", "Networking", "InstanceGroup"))
	build(kops.AddonSpec{}, required("Manifest"))
	build(kops.EgressProxySpec{}, required("HTTPProxy"))
	build(kops.HTTPProxy{}, required("Host", "Port"))
	build(kops.ContainerdConfig{})
	build(kops.DockerConfig{})
	build(kops.KubeDNSConfig{})
	build(kops.KubeAPIServerConfig{})
	build(kops.KubeControllerManagerConfig{})
	build(kops.CloudControllerManagerConfig{})
	build(kops.KubeSchedulerConfig{})
	build(kops.KubeProxyConfig{})
	build(kops.KubeletConfigSpec{})
	build(kops.CloudConfiguration{})
	build(kops.ExternalDNSConfig{})
	build(kops.OpenstackConfiguration{})
	build(kops.OpenstackLoadbalancerConfig{})
	build(kops.OpenstackMonitor{})
	build(kops.OpenstackRouter{})
	build(kops.OpenstackBlockStorageConfig{})
	build(kops.LeaderElectionConfiguration{})
	build(kops.NodeLocalDNSConfig{})
	build(kops.AuthenticationSpec{})
	build(kops.AuthorizationSpec{})
	build(kops.NodeAuthorizationSpec{})
	build(kops.Assets{})
	build(kops.IAMSpec{})
	build(kops.KopeioAuthenticationSpec{})
	build(kops.AwsAuthenticationSpec{})
	build(kops.AlwaysAllowAuthorizationSpec{})
	build(kops.RBACAuthorizationSpec{})
	build(kops.NodeAuthorizerSpec{})
	build(api.InstanceGroup{}, required("Name", "Role", "MinSize", "MaxSize", "MachineType", "Subnets"))
	build(kops.AccessSpec{})
	build(kops.DNSAccessSpec{})
	build(kops.LoadBalancerAccessSpec{}, required("Type"))
	build(kops.EtcdClusterSpec{}, required("Name", "Members"))
	build(kops.EtcdBackupSpec{}, required("BackupStore", "Image"))
	build(kops.EtcdManagerSpec{}, required("Image"))
	build(kops.EtcdMemberSpec{}, required("Name", "InstanceGroup"))
	build(kops.EnvVar{}, required("Name"))
	build(kops.ClusterSubnetSpec{}, required("Name", "ProviderID", "Type"))
	build(kops.TopologySpec{}, required("Masters", "Nodes", "DNS"))
	build(kops.BastionSpec{}, required("BastionPublicName"))
	build(kops.BastionLoadBalancerSpec{}, required("AdditionalSecurityGroups"))
	build(kops.DNSSpec{}, required("Type"))
	build(kops.NetworkingSpec{})
	build(kops.ClassicNetworkingSpec{})
	build(kops.KubenetNetworkingSpec{})
	build(kops.ExternalNetworkingSpec{})
	build(kops.CNINetworkingSpec{})
	build(kops.KopeioNetworkingSpec{})
	build(kops.WeaveNetworkingSpec{})
	build(kops.FlannelNetworkingSpec{})
	build(kops.CalicoNetworkingSpec{})
	build(kops.CanalNetworkingSpec{})
	build(kops.KuberouterNetworkingSpec{})
	build(kops.RomanaNetworkingSpec{})
	build(kops.AmazonVPCNetworkingSpec{})
	build(kops.CiliumNetworkingSpec{})
	build(kops.LyftVPCNetworkingSpec{})
	build(kops.GCENetworkingSpec{})
	build(kops.VolumeSpec{}, required("Device"))
	build(kops.VolumeMountSpec{}, required("Device", "Filesystem", "Path"))
	build(kops.MixedInstancesPolicySpec{})
	build(kops.UserData{}, required("Name", "Type", "Content"))
	build(kops.LoadBalancer{})
	build(kops.IAMProfileSpec{}, required("Profile"))
	build(kops.HookSpec{}, required("Name"))
	build(kops.ExecContainerAction{}, required("Image"))
	build(kops.FileAssetSpec{}, required("Name", "Path", "Content"))
	build(kops.RollingUpdate{})
}
