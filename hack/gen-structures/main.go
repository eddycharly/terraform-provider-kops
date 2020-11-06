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
	exclude      []string
	required     []string
	computed     []string
	computedOnly []string
	sensitive    []string
}

func excluded(excluded ...string) func(o options) options {
	return func(o options) options {
		o.exclude = append(o.exclude, excluded...)
		return o
	}
}

func required(required ...string) func(o options) options {
	return func(o options) options {
		o.required = append(o.required, required...)
		return o
	}
}

func computed(computed ...string) func(o options) options {
	return func(o options) options {
		o.computed = append(o.computed, computed...)
		return o
	}
}

func computedOnly(computedOnly ...string) func(o options) options {
	return func(o options) options {
		o.computedOnly = append(o.computedOnly, computedOnly...)
		return o
	}
}

func sensitive(sensitive ...string) func(o options) options {
	return func(o options) options {
		o.sensitive = append(o.sensitive, sensitive...)
		return o
	}
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

func funcMap(o options) template.FuncMap {
	required := sets.NewString(o.required...)
	computed := sets.NewString(o.computed...)
	computedOnly := sets.NewString(o.computedOnly...)
	sensitive := sets.NewString(o.sensitive...)
	return template.FuncMap{
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
		"isOptional": func(in string) bool {
			return !required.Has(in) && !computedOnly.Has(in)
		},
		"isComputed": func(in string) bool {
			return computed.Has(in) || computedOnly.Has(in)
		},
		"isSensitive": func(in string) bool {
			return sensitive.Has(in)
		},
		"fieldName":   fieldName,
		"isValueType": isValueType,
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
{{ if isRequired .Name }}:white_check_mark:{{ end }}
{{- end }}

{{- define "optional" -}}
{{ if isOptional .Name }}:white_check_mark:{{ end }}
{{- end }}

{{- define "computed" -}}
{{ if isComputed .Name }}:white_check_mark:{{ end }}
{{- end }}

| attribute | type | optional | required | computed |
| --- | --- | --- | --- | --- |
{{- with .Type -}}
{{- range (fields .) -}}
{{- if not (has .Name $.Exclude) }}
| {{ fieldName .Name | snakecase | code }} | {{ template "type" .Type }} | {{ template "required" . }} | {{ template "optional" . }} | {{ template "computed" . }} |
{{- end -}}
{{- end -}}
{{- end }}
`
	tmpl := template.New("doc").Funcs(funcMap(o)).Funcs(sprig.TxtFuncMap())

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

{{- define "modifier" -}}
{{- if isRequired . -}}
Required
{{- else -}}
{{- if isOptional . -}}
Optional
{{- end -}}
{{- if isComputed . -}}
Computed
{{- end -}}
{{- end -}}
{{- end -}}

{{- define "prop" -}}
{{- if isSensitive .Name -}}
Sensitive({{ template "modifier" .Name }}{{ template "schema" .Type }})
{{- else -}}
{{ template "modifier" .Name }}{{ template "schema" .Type }}
{{- end -}}
{{- end -}}

{{- with .Type }}

func {{ .Name }}() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			{{- range (fields .) }}
			{{- if not (has .Name $.Exclude) }}
			{{ fieldName .Name | snakecase | quote }}: {{ template "prop" . }},
			{{- end }}
			{{- end }}
		},
	}
}

{{- end }}
`
	tmpl := template.New("doc").Funcs(funcMap(o)).Funcs(sprig.TxtFuncMap())

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
	if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
		return nil
	}
	return func (in {{ .Elem.String }}) {{ .String }} {
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
	if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
		return nil
	}
	return func (in {{ .Elem.String }}) {{ .String }} {
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
	if len(in.([]interface{})) == 0 || in.([]interface{})[0] == nil {
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
	out := map[string]interface{}{}
	for key, in := range in {
		out[key] = {{ template "flattenElem" .Elem }}
	}
	return out
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
		{{- if and (isPtr .Type) (isValueType .Type) (not (isRequired .Name)) }}
		if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
			return nil
		}
		{{- end }}
		return {{ template "expand" .Type }}
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
		return {{ template "flatten" .Type }}
	}(in.{{ .Name }}),
	{{- end }}
	{{- end }}
	}
}

{{- end }}
`
	tmpl := template.New("doc").Funcs(funcMap(o)).Funcs(sprig.TxtFuncMap())

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
	build(api.Cluster{},
		required("Name", "AdminSshKey", "CloudProvider", "Subnet", "NetworkID", "Topology", "EtcdCluster", "Networking", "InstanceGroup"),
		computed("MasterPublicName", "MasterInternalName", "ConfigBase", "NetworkCIDR", "NonMasqueradeCIDR", "IAM"),
		computedOnly("KubeConfig"),
		sensitive("AdminSshKey", "KubeConfig"),
	)
	build(api.RollingUpdateOptions{})
	build(api.ValidateOptions{})
	build(api.KubeConfig{},
		computedOnly("Server", "Context", "Namespace", "KubeBearerToken", "KubeUser", "KubePassword", "CaCert", "ClientCert", "ClientKey"),
		sensitive("KubeBearerToken", "KubeUser", "KubePassword", "CaCert", "ClientCert", "ClientKey"),
	)
	build(kops.AddonSpec{},
		required("Manifest"),
	)
	build(kops.EgressProxySpec{},
		required("HTTPProxy"),
	)
	build(kops.HTTPProxy{},
		required("Host", "Port"),
	)
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
	build(api.InstanceGroup{},
		required("Name", "Role", "MinSize", "MaxSize", "MachineType", "Subnets"),
		computed("Image"),
	)
	build(kops.AccessSpec{})
	build(kops.DNSAccessSpec{})
	build(kops.LoadBalancerAccessSpec{},
		required("Type"),
	)
	build(kops.EtcdClusterSpec{},
		required("Name", "Members"),
	)
	build(kops.EtcdBackupSpec{},
		required("BackupStore", "Image"),
	)
	build(kops.EtcdManagerSpec{},
		required("Image"),
	)
	build(kops.EtcdMemberSpec{},
		required("Name", "InstanceGroup"),
	)
	build(kops.EnvVar{},
		required("Name"),
	)
	build(kops.ClusterSubnetSpec{},
		required("Name", "ProviderID", "Type", "Zone"),
		computed("CIDR"),
	)
	build(kops.TopologySpec{},
		required("Masters", "Nodes", "DNS"),
	)
	build(kops.BastionSpec{},
		required("BastionPublicName"),
	)
	build(kops.BastionLoadBalancerSpec{},
		required("AdditionalSecurityGroups"),
	)
	build(kops.DNSSpec{},
		required("Type"),
	)
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
	build(kops.VolumeSpec{},
		required("Device"),
	)
	build(kops.VolumeMountSpec{},
		required("Device", "Filesystem", "Path"),
	)
	build(kops.MixedInstancesPolicySpec{})
	build(kops.UserData{},
		required("Name", "Type", "Content"),
	)
	build(kops.LoadBalancer{})
	build(kops.IAMProfileSpec{},
		required("Profile"),
	)
	build(kops.HookSpec{},
		required("Name"),
	)
	build(kops.ExecContainerAction{},
		required("Image"),
	)
	build(kops.FileAssetSpec{},
		required("Name", "Path", "Content"),
	)
	build(kops.RollingUpdate{})
}
