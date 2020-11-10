package main

import (
	"fmt"
	"go/ast"
	"log"
	"os"
	"reflect"
	"regexp"
	"strings"
	"text/template"

	"github.com/Masterminds/sprig"
	"github.com/eddycharly/terraform-provider-kops/pkg/api"
	"golang.org/x/tools/go/packages"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/kops/pkg/apis/kops"
)

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

func toSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

type options struct {
	exclude      sets.String
	required     sets.String
	computed     sets.String
	computedOnly sets.String
	sensitive    sets.String
	suppressDiff sets.String
}

func newOptions() *options {
	return &options{
		exclude:      sets.NewString(),
		required:     sets.NewString(),
		computed:     sets.NewString(),
		computedOnly: sets.NewString(),
		sensitive:    sets.NewString(),
		suppressDiff: sets.NewString(),
	}
}

func excluded(excluded ...string) func(o *options) {
	return func(o *options) {
		o.exclude.Insert(excluded...)
	}
}

func required(required ...string) func(o *options) {
	return func(o *options) {
		o.required.Insert(required...)
	}
}

func computed(computed ...string) func(o *options) {
	return func(o *options) {
		o.computed.Insert(computed...)
	}
}

func computedOnly(computedOnly ...string) func(o *options) {
	return func(o *options) {
		o.computedOnly.Insert(computedOnly...)
	}
}

func sensitive(sensitive ...string) func(o *options) {
	return func(o *options) {
		o.sensitive.Insert(sensitive...)
	}
}

func suppressDiff(suppressDiff ...string) func(o *options) {
	return func(o *options) {
		o.suppressDiff.Insert(suppressDiff...)
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

func getSubResources(t reflect.Type, seen map[reflect.Type]bool) []reflect.Type {
	if t.Kind() == reflect.Array || t.Kind() == reflect.Map || t.Kind() == reflect.Slice || t.Kind() == reflect.Ptr {
		return getSubResources(t.Elem(), seen)
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
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		ret = append(ret, getSubResources(f.Type, seen)...)
	}
	return ret
}

func getResourceComment(packName, structName string, c map[string]map[string]map[string]string) string {
	if p, ok := c[packName]; ok {
		if s, ok := p[structName]; ok {
			ret := strings.ReplaceAll(strings.TrimSpace(s[""]), "\n", "<br />")
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

func getAttributeComment(packName, structName, fieldName string, c map[string]map[string]map[string]string) string {
	if p, ok := c[packName]; ok {
		if s, ok := p[structName]; ok {
			if a, ok := s[fieldName]; ok {
				ret := strings.ReplaceAll(strings.TrimSpace(a), "\n", "<br />")
				if ret != "" {
					if !strings.HasSuffix(ret, ".") {
						ret = ret + "."
					}
					return ret
				}
			}
		}
	}
	return ""
}

func funcMap(op map[reflect.Type]*options, c map[string]map[string]map[string]string) template.FuncMap {
	return template.FuncMap{
		"fields": func(t reflect.Type) []reflect.StructField {
			var ret []reflect.StructField
			for i := 0; i < t.NumField(); i++ {
				ret = append(ret, t.Field(i))
			}
			return ret
		},
		"schemaType": schemaType,
		"resourceComment": func(t reflect.Type) string {
			return getResourceComment(t.PkgPath(), t.Name(), c)
		},
		"attributeComment": func(t reflect.Type, f reflect.StructField) string {
			return getAttributeComment(t.PkgPath(), t.Name(), f.Name, c)
		},
		"subResources": func(t reflect.Type) []reflect.Type {
			return getSubResources(t, map[reflect.Type]bool{})[1:]
		},
		"isExcluded": func(in string, t reflect.Type) bool {
			return op[t].exclude.Has(in)
		},
		"isRequired": func(in string, t reflect.Type) bool {
			return op[t].required.Has(in)
		},
		"isOptional": func(in string, t reflect.Type) bool {
			return !op[t].required.Has(in) && !op[t].computedOnly.Has(in)
		},
		"isComputed": func(in string, t reflect.Type) bool {
			return op[t].computed.Has(in) || op[t].computedOnly.Has(in)
		},
		"isSensitive": func(in string, t reflect.Type) bool {
			return op[t].sensitive.Has(in)
		},
		"suppressDiff": func(in string, t reflect.Type) bool {
			return op[t].suppressDiff.Has(in)
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

func buildDoc(i interface{}, o map[reflect.Type]*options, c map[string]map[string]map[string]string) {
	t := reflect.TypeOf(i)
	tmplString := `# kops_{{ fieldName .Name | snakecase }}

{{- $comment := resourceComment . }}
{{ if $comment }}
{{ $comment }}
{{- end }}
	
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
[{{ .Name | snakecase }}](#{{ .Name | snakecase }})
{{- else -}}
{{- schemaType . -}}
{{- end -}}
{{- end }}

## Argument Reference

The following arguments are supported:

{{- range (fields .) -}}
{{- if not (isExcluded .Name $) }}
- {{ fieldName .Name | snakecase | code }}
{{- if isOptional .Name $ }} - (Optional){{ end }}
{{- if isRequired .Name $ }} - (Required){{ end }}
{{- if isComputed .Name $ }} - (Computed){{ end }} - {{ template "type" .Type }}{{ if (attributeComment $ .) }} - {{ attributeComment $ . }}{{ end }}
{{- end }}
{{- end }}

## Nested resources
{{ range $type := (subResources .) }}
### {{ fieldName .Name | snakecase }}
{{- $comment := resourceComment $type }}
{{ if $comment }}
{{ $comment }}
{{- end }}
{{ $fields := fields . -}}
{{ if eq 0 (len $fields) }}
This resource has no attributes.
{{- else -}}

#### Argument Reference

The following arguments are supported:
{{ range $fields }}
- {{ fieldName .Name | snakecase | code }}
{{- if isOptional .Name $type }} - (Optional){{ end }}
{{- if isRequired .Name $type }} - (Required){{ end }}
{{- if isComputed .Name $type }} - (Computed){{ end }} - {{ template "type" .Type }}{{ if (attributeComment $type .) }} - {{ attributeComment $type . }}{{ end }}
{{- end -}}
{{- end }}
{{ end }}
`
	tmpl := template.New("doc").Funcs(funcMap(o, c)).Funcs(sprig.TxtFuncMap())

	if _, err := tmpl.Parse(tmplString); err != nil {
		panic(err)
	}

	fileName := toSnakeCase(fieldName(t.Name())) + ".md"
	f, _ := os.Create("docs/resources/" + fileName)
	defer f.Close()

	if err := tmpl.Execute(f, t); err != nil {
		panic(err)
	}
}

func buildSchema(t reflect.Type, o map[reflect.Type]*options) {
	tmplString := `
package schemas

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
{{- end }}

func {{ .Name }}() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			{{- range (fields .) }}
			{{- if not (isExcluded .Name $) }}
			{{ fieldName .Name | snakecase | quote }}:
			{{- $sensitive := isSensitive .Name $ -}}
			{{- $suppressDiff := suppressDiff .Name $ -}}
			{{- if $suppressDiff -}}SuppressDiff({{- end -}}
			{{- if $sensitive -}}Sensitive({{- end -}}
			{{- if isRequired .Name $ -}}
			Required
			{{- else if isOptional .Name $ -}}
			Optional
			{{- end -}}
			{{- if isComputed .Name $ -}}
			Computed
			{{- end -}}
			{{ template "schema" .Type }}
			{{- if $suppressDiff -}}){{- end -}}
			{{- if $sensitive -}}){{- end -}}
			,
			{{- end -}}
			{{- end }}
		},
	}
}
`
	tmpl := template.New("doc").Funcs(funcMap(o, nil)).Funcs(sprig.TxtFuncMap())

	if _, err := tmpl.Parse(tmplString); err != nil {
		panic(err)
	}

	f, _ := os.Create("pkg/schemas/" + t.Name() + ".generated.go")
	defer f.Close()

	if err := tmpl.Execute(f, t); err != nil {
		panic(err)
	}
}

func buildStructure(t reflect.Type, o map[reflect.Type]*options) {
	tmplString := `
package structures

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
{{- end }}

func Expand{{ .Name }}(in map[string]interface{}) {{ .String }} {
	if in == nil {
		panic("expand {{ .Name }} failure, in is nil")
	}
	return {{ .String }}{
	{{- range (fields .) }}
	{{- if not (isExcluded .Name $) }}
	{{ .Name }}: func (in interface{}) {{ .Type.String }} {
		{{- if and (isPtr .Type) (isValueType .Type) (not (isRequired .Name $)) }}
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
	{{- if not (isExcluded .Name $) }}
	{{ fieldName .Name | snakecase | quote }}: func (in {{ .Type.String }}) interface{} {
		return {{ template "flatten" .Type }}
	}(in.{{ .Name }}),
	{{- end }}
	{{- end }}
	}
}
`
	tmpl := template.New("doc").Funcs(funcMap(o, nil)).Funcs(sprig.TxtFuncMap())

	if _, err := tmpl.Parse(tmplString); err != nil {
		panic(err)
	}

	f, _ := os.Create("pkg/structures/" + t.Name() + ".generated.go")
	defer f.Close()

	if err := tmpl.Execute(f, t); err != nil {
		panic(err)
	}
}

type generated struct {
	t reflect.Type
	o *options
}

func generate(i interface{}, opts ...func(o *options)) generated {
	o := newOptions()
	for _, opt := range opts {
		opt(o)
	}
	t := reflect.TypeOf(i)
	return generated{
		t: t,
		o: o,
	}
}

func build(g ...generated) map[reflect.Type]*options {
	o := map[reflect.Type]*options{}
	for _, gen := range g {
		o[gen.t] = gen.o
	}
	for _, gen := range g {
		buildStructure(gen.t, o)
		buildSchema(gen.t, o)
	}
	return o
}

func (p *parser) parseStruct(typeSpec *ast.TypeSpec) map[string]string {
	ret := make(map[string]string)
	structure, ok := typeSpec.Type.(*ast.StructType)
	if ok {
		for _, field := range structure.Fields.List {
			for _, name := range field.Names {
				ret[name.Name] = field.Doc.Text()
			}
		}
	}
	return ret
}

func (p *parser) parsePackage(pack *packages.Package) {
	if _, ok := p.packs[pack.ID]; !ok {
		p.packs[pack.ID] = make(map[string]map[string]string)
		for _, v := range pack.Imports {
			p.parsePackage(v)
		}
		for _, node := range pack.Syntax {
			for _, decl := range node.Decls {
				genDecl, ok := decl.(*ast.GenDecl)
				if ok {
					for _, spec := range genDecl.Specs {
						typeSpec, ok := spec.(*ast.TypeSpec)
						if ok {
							p.packs[pack.ID][typeSpec.Name.Name] = p.parseStruct(typeSpec)
							p.packs[pack.ID][typeSpec.Name.Name][""] = genDecl.Doc.Text()
						}
					}
				}
			}
		}
	}
}

type parser struct {
	packages []*packages.Package
	packs    map[string]map[string]map[string]string
}

func main() {
	log.Println("loading packages...")
	cfg := packages.Config{
		Mode: packages.NeedName | packages.NeedFiles | packages.NeedCompiledGoFiles | packages.NeedImports | packages.NeedDeps | packages.NeedTypes | packages.NeedSyntax | packages.NeedTypesInfo,
	}
	packs, err := packages.Load(
		&cfg,
		"github.com/eddycharly/terraform-provider-kops/pkg/api",
	)
	if err != nil {
		panic(err)
	}
	parser := parser{
		packages: packs,
		packs:    make(map[string]map[string]map[string]string),
	}
	for _, pack := range packs {
		parser.parsePackage(pack)
	}
	log.Println("generating schemas, expanders and flatteners...")
	typeMap := build(
		generate(api.ProviderConfig{},
			required("StateStore"),
		),
		generate(api.AwsConfig{}),
		generate(api.AwsAssumeRole{}),
		generate(api.Cluster{},
			required("Name", "AdminSshKey", "CloudProvider", "Subnet", "NetworkID", "Topology", "EtcdCluster", "Networking", "InstanceGroup"),
			computed("MasterPublicName", "MasterInternalName", "ConfigBase", "NetworkCIDR", "NonMasqueradeCIDR", "IAM"),
			computedOnly("KubeConfig"),
			sensitive("AdminSshKey", "KubeConfig"),
			suppressDiff("RollingUpdateOptions", "ValidateOptions"),
		),
		generate(api.RollingUpdateOptions{}),
		generate(api.ValidateOptions{}),
		generate(api.KubeConfig{},
			computedOnly("Server", "Context", "Namespace", "KubeBearerToken", "KubeUser", "KubePassword", "CaCert", "ClientCert", "ClientKey"),
			sensitive("KubeBearerToken", "KubeUser", "KubePassword", "CaCert", "ClientCert", "ClientKey"),
		),
		generate(kops.AddonSpec{},
			required("Manifest"),
		),
		generate(kops.EgressProxySpec{},
			required("HTTPProxy"),
		),
		generate(kops.HTTPProxy{},
			required("Host", "Port"),
		),
		generate(kops.ContainerdConfig{}),
		generate(kops.DockerConfig{}),
		generate(kops.KubeDNSConfig{}),
		generate(kops.KubeAPIServerConfig{}),
		generate(kops.KubeControllerManagerConfig{}),
		generate(kops.CloudControllerManagerConfig{}),
		generate(kops.KubeSchedulerConfig{}),
		generate(kops.KubeProxyConfig{}),
		generate(kops.KubeletConfigSpec{}),
		generate(kops.CloudConfiguration{}),
		generate(kops.ExternalDNSConfig{}),
		generate(kops.OpenstackConfiguration{}),
		generate(kops.OpenstackLoadbalancerConfig{}),
		generate(kops.OpenstackMonitor{}),
		generate(kops.OpenstackRouter{}),
		generate(kops.OpenstackBlockStorageConfig{}),
		generate(kops.LeaderElectionConfiguration{}),
		generate(kops.NodeLocalDNSConfig{}),
		generate(kops.AuthenticationSpec{}),
		generate(kops.AuthorizationSpec{}),
		generate(kops.NodeAuthorizationSpec{}),
		generate(kops.Assets{}),
		generate(kops.IAMSpec{}),
		generate(kops.KopeioAuthenticationSpec{}),
		generate(kops.AwsAuthenticationSpec{}),
		generate(kops.AlwaysAllowAuthorizationSpec{}),
		generate(kops.RBACAuthorizationSpec{}),
		generate(kops.NodeAuthorizerSpec{}),
		generate(api.InstanceGroup{},
			required("Name", "Role", "MinSize", "MaxSize", "MachineType", "Subnets"),
			computed("Image"),
		),
		generate(kops.AccessSpec{}),
		generate(kops.DNSAccessSpec{}),
		generate(kops.LoadBalancerAccessSpec{},
			required("Type"),
		),
		generate(kops.EtcdClusterSpec{},
			required("Name", "Members"),
		),
		generate(kops.EtcdBackupSpec{},
			required("BackupStore", "Image"),
		),
		generate(kops.EtcdManagerSpec{},
			required("Image"),
		),
		generate(kops.EtcdMemberSpec{},
			required("Name", "InstanceGroup"),
		),
		generate(kops.EnvVar{},
			required("Name"),
		),
		generate(kops.ClusterSubnetSpec{},
			required("Name", "ProviderID", "Type", "Zone"),
			computed("CIDR"),
		),
		generate(kops.TopologySpec{},
			required("Masters", "Nodes", "DNS"),
		),
		generate(kops.BastionSpec{},
			required("BastionPublicName"),
		),
		generate(kops.BastionLoadBalancerSpec{},
			required("AdditionalSecurityGroups"),
		),
		generate(kops.DNSSpec{},
			required("Type"),
		),
		generate(kops.NetworkingSpec{}),
		generate(kops.ClassicNetworkingSpec{}),
		generate(kops.KubenetNetworkingSpec{}),
		generate(kops.ExternalNetworkingSpec{}),
		generate(kops.CNINetworkingSpec{}),
		generate(kops.KopeioNetworkingSpec{}),
		generate(kops.WeaveNetworkingSpec{}),
		generate(kops.FlannelNetworkingSpec{}),
		generate(kops.CalicoNetworkingSpec{}),
		generate(kops.CanalNetworkingSpec{}),
		generate(kops.KuberouterNetworkingSpec{}),
		generate(kops.RomanaNetworkingSpec{}),
		generate(kops.AmazonVPCNetworkingSpec{}),
		generate(kops.CiliumNetworkingSpec{}),
		generate(kops.LyftVPCNetworkingSpec{}),
		generate(kops.GCENetworkingSpec{}),
		generate(kops.VolumeSpec{},
			required("Device"),
		),
		generate(kops.VolumeMountSpec{},
			required("Device", "Filesystem", "Path"),
		),
		generate(kops.MixedInstancesPolicySpec{}),
		generate(kops.UserData{},
			required("Name", "Type", "Content"),
		),
		generate(kops.LoadBalancer{}),
		generate(kops.IAMProfileSpec{},
			required("Profile"),
		),
		generate(kops.HookSpec{},
			required("Name"),
		),
		generate(kops.ExecContainerAction{},
			required("Image"),
		),
		generate(kops.FileAssetSpec{},
			required("Name", "Path", "Content"),
		),
		generate(kops.RollingUpdate{}),
	)
	log.Println("generating docs...")
	buildDoc(api.Cluster{}, typeMap, parser.packs)
}
