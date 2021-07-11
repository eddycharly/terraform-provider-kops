package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"reflect"
	"regexp"
	"strings"
	"text/template"

	"github.com/Masterminds/sprig"
	"github.com/eddycharly/terraform-provider-kops/pkg/api/config"
	"github.com/eddycharly/terraform-provider-kops/pkg/api/datasources"
	"github.com/eddycharly/terraform-provider-kops/pkg/api/kube"
	"github.com/eddycharly/terraform-provider-kops/pkg/api/resources"
	"github.com/eddycharly/terraform-provider-kops/pkg/api/utils"
	"k8s.io/kops/pkg/apis/kops"
)

var mappings = map[string]string{
	"github.com/eddycharly/terraform-provider-kops/pkg/api/config":      "config",
	"github.com/eddycharly/terraform-provider-kops/pkg/api/datasources": "datasources",
	"github.com/eddycharly/terraform-provider-kops/pkg/api/kube":        "kube",
	"github.com/eddycharly/terraform-provider-kops/pkg/api/resources":   "resources",
	"github.com/eddycharly/terraform-provider-kops/pkg/api/utils":       "utils",
	"k8s.io/kops/pkg/apis/kops":                                         "kops",
}

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

func toSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
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

func funcMap(baseType reflect.Type, optionsMap map[reflect.Type]*options, scope string, parser *parser) template.FuncMap {
	return template.FuncMap{
		"scope": func() string {
			return scope
		},
		"imports": func() map[string]string {
			out := make(map[string]string)
			for i, v := range parser.packages[baseType.PkgPath()].Imports {
				if v, ok := mappings[i]; ok {
					out["github.com/eddycharly/terraform-provider-kops/pkg/schemas/"+v] = v + "schemas"
				}
				for i2 := range parser.packages[v.PkgPath].Imports {
					if v, ok := mappings[i2]; ok {
						out["github.com/eddycharly/terraform-provider-kops/pkg/schemas/"+v] = v + "schemas"
					}
				}
			}
			return out
		},
		"mapping": func(t reflect.Type) string {
			if baseType.PkgPath() == t.PkgPath() {
				return ""
			}
			return mappings[t.PkgPath()] + "schemas."
		},
		"schemaType": schemaType,
		"fieldName": func(in _field) string {
			if optionsMap[in.Owner].rename[in.Name] != "" {
				return fieldName(optionsMap[in.Owner].rename[in.Name])
			}
			return fieldName(in.Name)
		},
		"schemaName": func(in string) string {
			return fieldName(in)
		},
		"code": func(in string) string {
			return fmt.Sprintf("`%s`", in)
		},
	}
}

func buildDoc(i interface{}, p string, optionsMap map[reflect.Type]*options, scope string, parser *parser, header, footer string) {
	t := reflect.TypeOf(i)
	tmplString := fmt.Sprintf(`# kops_{{ schemaName .Name | snakecase }}

%s

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
[{{ schemaName .Name | snakecase }}](#{{ schemaName .Name | snakecase }})
{{- else -}}
{{- schemaType . -}}
{{- end -}}
{{- end }}

## Argument Reference

The following arguments are supported:

{{- range (fields . true) -}}
{{- if not (isExcluded .) }}
- {{ fieldName . | snakecase | code }}
{{- if isOptional . }} - (Optional){{ end }}
{{- if isRequired . }} - (Required){{ end }}
{{- if forceNew . }} - (Force new){{ end }}
{{- if isSensitive . }} - (Sensitive){{ end }}
{{- if isComputed . }} - (Computed){{ end }} - {{ if isNullable . -}}
{{ template "type" .Type }}([Nullable](#nullable-arguments))
{{- else -}}
{{ template "type" .Type }}
{{- end -}}{{ if (attributeComment .) }} - {{ attributeComment . }}{{ end }}
{{- end }}
{{- end }}

{{ if (subResources .) -}}
## Nested resources
{{ range $type := (subResources .) }}
### {{ schemaName .Name | snakecase }}
{{- $comment := resourceComment $type }}
{{ if $comment }}
{{ $comment }}
{{ end }}
{{ $fields := fields . true -}}
{{ if eq 0 (len $fields) }}
This resource has no attributes.
{{- else -}}

#### Argument Reference

The following arguments are supported:
{{ range $fields }}
{{- if not (isExcluded .) }}
- {{ fieldName . | snakecase | code }}
{{- if isOptional . }} - (Optional){{ end }}
{{- if isRequired . }} - (Required){{ end }}
{{- if forceNew . }} - (Force new){{ end }}
{{- if isSensitive . }} - (Sensitive){{ end }}
{{- if isComputed . }} - (Computed){{ end }} - {{ if isNullable . -}}
{{ template "type" .Type }}([Nullable](#nullable-arguments))
{{- else -}}
{{ template "type" .Type }}
{{- end -}}{{ if (attributeComment .) }} - {{ attributeComment . }}{{ end }}
{{- end -}}
{{- end }}
{{- end }}
{{ end }}
{{- end }}

%s
`, header, footer)

	tmpl := template.New("doc")
	funcMaps := []template.FuncMap{
		reflectFuncs(),
		docFuncs(parser, optionsMap),
		optionFuncs(scope == "DataSource", optionsMap),
		funcMap(t, optionsMap, scope, parser),
		sprig.TxtFuncMap(),
	}
	for _, funcMap := range funcMaps {
		tmpl = tmpl.Funcs(funcMap)
	}
	if _, err := tmpl.Parse(tmplString); err != nil {
		panic(err)
	}
	if err := os.MkdirAll(p, 0755); err != nil {
		panic(fmt.Sprintf("Failed to create directories for %s", p))
	}
	fileName := toSnakeCase(fieldName(t.Name())) + ".md"
	f, err := os.Create(path.Join(p, fileName))
	if err != nil {
		panic(fmt.Sprintf("Failed to create file %s", path.Join(p, fileName)))
	}
	defer f.Close()
	if err := tmpl.Execute(f, t); err != nil {
		panic(err)
	}
}

func buildSchema(t reflect.Type, p string, scope string, funcMaps ...template.FuncMap) {
	tmplString := `
package schemas

import (
	"encoding/json"
	"log"
	"strings"
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/kops/pkg/apis/kops"
	"github.com/eddycharly/terraform-provider-kops/pkg/api/config"
	"github.com/eddycharly/terraform-provider-kops/pkg/api/datasources"
	"github.com/eddycharly/terraform-provider-kops/pkg/api/kube"
	"github.com/eddycharly/terraform-provider-kops/pkg/api/resources"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	{{ range $k, $v := imports -}}
	{{ $v }} {{ $k | quote }}
	{{ end -}}
)

var _ = Schema

{{- define "schemaElem" -}}
{{- if isPtr . -}}
{{- template "schemaElem" .Elem }}
{{- else if isStruct . -}}
{{ mapping . }}{{ scope }}{{ .Name }}()
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
Struct({{ mapping . }}{{ scope }}{{ .Name }}())
{{- else -}}
{{- schemaType . -}}()
{{- end -}}
{{- end }}

{{ if needsSchema . }}
func {{ scope }}{{ .Name }}() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			{{- range (fields . true) }}
			{{- if not (isExcluded .) }}
			{{ fieldName . | snakecase | quote }}:
			{{- $forceNew := forceNew . -}}
			{{- $sensitive := isSensitive . -}}
			{{- $suppressDiff := suppressDiff . -}}
			{{- if $suppressDiff -}}SuppressDiff({{- end -}}
			{{- if $forceNew -}}ForceNew({{- end -}}
			{{- if $sensitive -}}Sensitive({{- end -}}
			{{- if isNullable . -}}
			Nullable(
			{{- end -}}
			{{- if isRequired . -}}
			Required
			{{- else if isOptional . -}}
			Optional
			{{- end -}}
			{{- if isComputed . -}}
			Computed
			{{- end -}}
			{{ template "schema" .Type }}
			{{- if isNullable . -}}
			)
			{{- end -}}
			{{- if $suppressDiff -}}){{- end -}}
			{{- if $forceNew -}}){{- end -}}
			{{- if $sensitive -}}){{- end -}}
			,
			{{- end -}}
			{{- end }}
		},
	}
	{{ if hasVersion . -}}
	res.SchemaVersion = {{ schemaVersion . }}
	res.StateUpgraders = []schema.StateUpgrader{
		{{ range $version := until (schemaVersion .) -}}
		{
				Type:    res.CoreConfigSchema().ImpliedType(),
				Upgrade: func (ctx context.Context, rawState map[string]interface{}, meta interface{}) (map[string]interface{}, error) {
					ret := Flatten{{ scope }}{{ $.Name }}(Expand{{ scope }}{{ $.Name }}(rawState))
					ret["id"] = rawState["id"]
					return ret, nil
				},
				Version: {{ $version}},
		},
		{{- end }}
	}
	{{- end }}
	return res
}
{{ end }}

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
	return ({{ mapping . }}Expand{{ scope }}{{ .Name }}(in.(map[string]interface{})))
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
	if in == nil {
		return nil
	}
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
	if in, ok := in.(map[string]interface{}); ok {
		if len(in) > 0 {
			out := {{ .String }}{}
			for key, in := range in {
				out[key] = {{ template "expand" .Elem }}
			}
			return out
		}
	}
	return nil
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
	return ({{ mapping . }}Expand{{ scope }}{{ .Name }}(in.([]interface{})[0].(map[string]interface{})))
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
	return {{ mapping . }}Flatten{{ scope }}{{ .Name }}(in)
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
func (in {{ .String }}) []interface{} {
	return []interface{}{ {{ mapping . }}Flatten{{ scope }}{{ .Name }}(in) }
}(in)
{{- else -}}
Flatten{{ schemaType . }}({{ schemaType . | lower }}(in))
{{- end -}}
{{- end }}

func Expand{{ scope }}{{ .Name }}(in map[string]interface{}) {{ .String }} {
	if in == nil {
		panic("expand {{ .Name }} failure, in is nil")
	}
	return {{ .String }}{
	{{- range (fields . false) }}
	{{- if not (isExcluded .) }}
	{{ .Name }}: func (in interface{}) {{ .Type.String }} {
		{{- if not .Anonymous -}}
		{{- if and (isPtr .Type) (isValueType .Type) (not (isRequired .)) (not (isNullable .)) }}
		if in == nil {
			return nil
		}
		if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
			return nil
		}
		{{- end }}
		{{ if isNullable . -}}
		if in == nil {
			return nil
		}
		if in, ok := in.([]interface{}); ok && len(in) == 1 {
			return func(in interface{}) {{ .Type.String }} {
				return {{ template "expand" .Type }}
			}(in[0].(map[string]interface{})["value"])
		}
		return nil
		{{- else -}}
		return {{ template "expand" .Type }}
		{{- end -}}
		{{- else -}}
		return {{ mapping .Type }}Expand{{ scope }}{{ .Type.Name }}(in.(map[string]interface{}))
		{{- end }}
	}({{ if .Anonymous }}in{{ else }}in[{{ fieldName . | snakecase | quote }}]{{ end }}),
	{{- end }}
	{{- end }}
	}
}

func Flatten{{ scope }}{{ .Name }}Into(in {{ .String }}, out map[string]interface{}) {
	{{- range (fields . false) }}
	{{- if not (isExcluded .) }}
	{{ if .Anonymous -}}
	{{ mapping .Type }}Flatten{{ scope }}{{ .Type.Name }}Into(in.{{ .Name }}, out)
	{{- else -}}
	out[{{ fieldName . | snakecase | quote }}] = func (in {{ .Type.String }}) interface{} {
		{{ if isNullable . -}}
		if in == nil {
			return nil
		}
		return []interface{}{map[string]interface{}{"value": {{ template "flatten" .Type }}}}
		{{- else -}}
		return {{ template "flatten" .Type }}
		{{- end }}
	}(in.{{ .Name }})
	{{- end }}
	{{- end }}
	{{- end }}
}

func Flatten{{ scope }}{{ .Name }}(in {{ .String }}) map[string]interface{} {
	out := map[string]interface{}{}
	Flatten{{ scope }}{{ .Name }}Into(in, out)
	return out
}
`
	tmpl := template.New("doc")
	for _, funcMap := range funcMaps {
		tmpl = tmpl.Funcs(funcMap)
	}
	if _, err := tmpl.Parse(tmplString); err != nil {
		panic(err)
	}
	folder := p
	if err := os.MkdirAll(folder, 0755); err != nil {
		panic(fmt.Sprintf("Failed to create directories for %s", folder))
	}
	fileName := fmt.Sprintf("%s_%s.generated.go", scope, t.Name())
	f, err := os.Create(path.Join(folder, fileName))
	if err != nil {
		panic(fmt.Sprintf("Failed to create file %s", path.Join(p, fileName)))
	}
	defer f.Close()
	if err := tmpl.Execute(f, t); err != nil {
		panic(err)
	}
}

func buildTests(t reflect.Type, p string, scope string, funcMaps ...template.FuncMap) {
	tmplString := `
package schemas

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	{{ range $k, $v := imports -}}
	{{ $v }} {{ $k | quote }}
	{{ end -}}
)

{{- define "defaultValue" -}}
{{- if isPtr . -}}
nil
{{- else if isList . -}}
nil
{{- else if isMap . -}}
nil
{{- else if isDuration . -}}
""
{{- else if isQuantity . -}}
""
{{- else if isIntOrString . -}}
""
{{- else if isStruct . -}}
{{ .String }}{}
{{- else if isInt . -}}
0
{{- else if isBool . -}}
false
{{- else if isFloat . -}}
0
{{- else if isString . -}}
""
{{- end -}}
{{- end }}

{{- define "expandDefault" -}}
{{- if isPtr . -}}
nil
{{- else if isList . -}}
func() []interface{} { return nil }()
{{- else if isMap . -}}
func() map[string]interface{} { return nil }()
{{- else if isDuration . -}}
""
{{- else if isQuantity . -}}
""
{{- else if isIntOrString . -}}
""
{{- else if isStruct . -}}
func() []interface{}{ return []interface{}{ {{ mapping . }}Flatten{{ scope }}{{ .Name }}({{ .String }}{}) } }()
{{- else if isInt . -}}
0
{{- else if isBool . -}}
false
{{- else if isFloat . -}}
0
{{- else if isString . -}}
""
{{- end -}}
{{- end -}}

{{- define "flattenDefault" -}}
{{- if isPtr . -}}
nil
{{- else if isList . -}}
func() []interface{} { return nil }()
{{- else if isMap . -}}
func() map[string]interface{} { return nil }()
{{- else if isDuration . -}}
""
{{- else if isQuantity . -}}
""
{{- else if isIntOrString . -}}
""
{{- else if isStruct . -}}
func() []interface{}{ return []interface{}{ {{ mapping . }}Flatten{{ scope }}{{ .Name }}({{ .String }}{}) } }()
{{- else if isInt . -}}
0
{{- else if isBool . -}}
false
{{- else if isFloat . -}}
0
{{- else if isString . -}}
""
{{- end -}}
{{- end -}}

{{- define "expandCases" -}}
{{- $fields := (fields . true) -}}
_default := {{ $.String }}{}
type args struct {
	in map[string]interface{}
}
tests := []struct {
	name string
	args args
	want {{ .String }}
}{
	{
		name: "default",
		args: args{
			in: map[string]interface{}{
				{{- range $fields }}
				{{- if not (isExcluded .) }}
				{{ fieldName . | snakecase | quote }}: {{ template "expandDefault" .Type }},
				{{- end }}
				{{- end }}
			},
		},
		want: _default,
	},
}
{{- end -}}

{{- define "flattenCases" -}}
{{- $fields := (fields . true) -}}
_default := map[string]interface{}{
	{{- range $fields }}
	{{- if not (isExcluded .) }}
	{{ fieldName . | snakecase | quote }}: {{ template "flattenDefault" .Type }},
	{{- end }}
	{{- end }}
}
type args struct {
	in {{ .String }}
}
tests := []struct {
	name string
	args args
	want map[string]interface{}
}{
	{
		name: "default",
		args: args{
			in: {{ .String }}{},
		},
		want: _default,
	},
	{{- range $fields }}
	{{- $field := . }}
	{{- if not (isExcluded .) }}
	{
		name: "{{ fieldName . }} - default",
		args: args{
			in: func() {{ $.String }}{
				subject := {{ $.String }}{}
				subject.{{ .Name }} = {{ template "defaultValue" .Type }}
				return subject
			}(),
		},
		want: _default,
	},
	{{- end }}
	{{- end }}
}
{{- end }}

func TestExpand{{ scope }}{{ .Name }}(t *testing.T) {
	{{ template "expandCases" . }}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Expand{{ scope }}{{ .Name }}(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("Expand{{ scope }}{{ .Name }}() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlatten{{ scope }}{{ .Name }}Into(t *testing.T) {
	{{ template "flattenCases" . }}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			Flatten{{ scope }}{{ .Name }}Into(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("Flatten{{ scope }}{{ .Name }}() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlatten{{ scope }}{{ .Name }}(t *testing.T) {
	{{ template "flattenCases" . }}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Flatten{{ scope }}{{ .Name }}(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("Flatten{{ scope }}{{ .Name }}() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
`
	tmpl := template.New("doc")
	for _, funcMap := range funcMaps {
		tmpl = tmpl.Funcs(funcMap)
	}
	if _, err := tmpl.Parse(tmplString); err != nil {
		panic(err)
	}
	folder := p
	if err := os.MkdirAll(folder, 0755); err != nil {
		panic(fmt.Sprintf("Failed to create directories for %s", folder))
	}
	fileName := fmt.Sprintf("%s_%s.generated_test.go", scope, t.Name())
	f, err := os.Create(path.Join(folder, fileName))
	if err != nil {
		panic(fmt.Sprintf("Failed to create file %s", path.Join(p, fileName)))
	}
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
	t := reflect.TypeOf(i)
	o := newOptions()
	for _, opt := range opts {
		opt(o)
	}
	if err := o.verify(t); err != nil {
		panic(err)
	}
	return generated{
		t: t,
		o: o,
	}
}

func build(scope string, parser *parser, g ...generated) map[reflect.Type]*options {
	o := map[reflect.Type]*options{}
	for _, gen := range g {
		o[gen.t] = gen.o
	}
	for _, gen := range g {
		funcMaps := []template.FuncMap{
			reflectFuncs(),
			optionFuncs(scope == "DataSource", o),
			funcMap(gen.t, o, scope, parser),
			sprig.TxtFuncMap(),
		}
		buildSchema(gen.t, path.Join("pkg/schemas", mappings[gen.t.PkgPath()]), scope, funcMaps...)
		buildTests(gen.t, path.Join("pkg/schemas", mappings[gen.t.PkgPath()]), scope, funcMaps...)
	}
	return o
}

func main() {
	log.Println("loading packages...")
	parser, err := initParser(
		"github.com/eddycharly/terraform-provider-kops/pkg/api/config",
		"github.com/eddycharly/terraform-provider-kops/pkg/api/datasources",
		"github.com/eddycharly/terraform-provider-kops/pkg/api/kube",
		"github.com/eddycharly/terraform-provider-kops/pkg/api/resources",
	)
	if err != nil {
		panic(err)
	}
	log.Println("generating schemas, expanders and flatteners...")
	resourcesMap := build(
		"Resource",
		parser,
		generate(resources.ClusterUpdater{},
			required("ClusterName"),
			computedOnly("Revision"),
		),
		generate(utils.RollingUpdateOptions{},
			noSchema(),
		),
		generate(resources.ClusterSecrets{},
			sensitive("DockerConfig"),
		),
		generate(resources.ValidateOptions{}),
		generate(utils.ValidateOptions{},
			noSchema(),
		),
		generate(resources.ApplyOptions{}),
		generate(resources.Cluster{},
			version(1),
			required("Name", "AdminSshKey"),
			computedOnly("Revision"),
			sensitive("AdminSshKey"),
			forceNew("Name"),
		),
		generate(kops.ClusterSpec{},
			noSchema(),
			exclude("GossipConfig", "DNSControllerGossipConfig", "Target"),
			rename("Subnets", "Subnet"),
			rename("EtcdClusters", "EtcdCluster"),
			required("CloudProvider", "Subnets", "NetworkID", "Topology", "EtcdClusters", "Networking"),
			computed("MasterPublicName", "MasterInternalName", "ConfigBase", "NetworkCIDR", "NonMasqueradeCIDR", "IAM"),
		),
		generate(kops.InstanceMetadataOptions{}),
		generate(kops.NodeTerminationHandlerConfig{}),
		generate(kops.MetricsServerConfig{}),
		generate(kops.ClusterAutoscalerConfig{}),
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
		generate(kops.PackagesConfig{}),
		generate(kops.DockerConfig{}),
		generate(kops.KubeDNSConfig{}),
		generate(kops.KubeAPIServerConfig{},
			nullable("AnonymousAuth"),
		),
		generate(kops.KubeControllerManagerConfig{}),
		generate(kops.CloudControllerManagerConfig{}),
		generate(kops.KubeSchedulerConfig{}),
		generate(kops.KubeProxyConfig{}),
		generate(kops.KubeletConfigSpec{},
			nullable("AnonymousAuth"),
		),
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
		generate(resources.InstanceGroup{},
			version(1),
			required("ClusterName", "Name"),
			forceNew("ClusterName", "Name"),
			computedOnly("Revision"),
		),
		generate(kops.InstanceGroupSpec{},
			noSchema(),
			required("Role", "MinSize", "MaxSize", "MachineType", "Subnets"),
			computed("Image"),
		),
		generate(kops.AccessSpec{}),
		generate(kops.DNSAccessSpec{}),
		generate(kops.LoadBalancerAccessSpec{},
			required("Type"),
		),
		generate(kops.EtcdClusterSpec{},
			required("Name", "Members"),
			rename("Members", "Member"),
		),
		generate(kops.EtcdBackupSpec{},
			required("BackupStore", "Image"),
		),
		generate(kops.EtcdManagerSpec{}),
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
		generate(kops.HubbleSpec{}),
		generate(kops.LyftVPCNetworkingSpec{}),
		generate(kops.GCENetworkingSpec{}),
		generate(kops.VolumeSpec{},
			required("Device"),
		),
		generate(kops.VolumeMountSpec{},
			required("Device", "Filesystem", "Path"),
		),
		generate(kops.MixedInstancesPolicySpec{},
			nullable("OnDemandBase", "OnDemandAboveBase"),
		),
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
		// 1.20
		generate(resources.RollingUpdateOptions{}),
		generate(kops.AzureConfiguration{}),
		generate(kops.AWSEBSCSIDriver{}),
		generate(kops.NTPConfig{}),
		generate(kops.CertManagerConfig{}),
		generate(kops.AWSLoadBalancerControllerConfig{}),
		generate(kops.GossipConfigSecondary{}),
		generate(kops.LoadBalancerSubnetSpec{}),
		generate(kops.DNSControllerGossipConfigSecondary{}),
		generate(kops.OpenstackNetwork{}),
		// 1.21
		generate(kops.WarmPoolSpec{}),
		generate(kops.ServiceAccountIssuerDiscoveryConfig{}),
		generate(kops.SnapshotControllerConfig{}),
		generate(kops.ServiceAccountExternalPermission{}),
		generate(kops.AWSPermission{}),
	)
	configMap := build(
		"Config",
		parser,
		generate(config.Provider{},
			required("StateStore"),
		),
		generate(config.Aws{}),
		generate(config.AwsAssumeRole{}),
		generate(config.Openstack{}),
		generate(config.Klog{},
			nullable("Verbosity"),
		),
	)
	dataSourcesMap := build(
		"DataSource",
		parser,
		generate(datasources.KubeConfig{},
			required("ClusterName"),
			computed("Admin", "Internal"),
		),
		generate(datasources.ClusterStatus{},
			required("ClusterName"),
		),
		generate(resources.ClusterSecrets{},
			sensitive("DockerConfig"),
		),
		generate(kube.Config{},
			noSchema(),
			sensitive("KubeBearerToken", "KubePassword", "CaCert", "ClientCert", "ClientKey"),
		),
		generate(resources.Cluster{},
			version(1),
			required("Name"),
			exclude("Revision"),
		),
		generate(kops.ClusterSpec{},
			exclude("GossipConfig", "DNSControllerGossipConfig", "Target"),
			rename("Subnets", "Subnet"),
			rename("EtcdClusters", "EtcdCluster"),
		),
		generate(kops.InstanceMetadataOptions{}),
		generate(kops.NodeTerminationHandlerConfig{}),
		generate(kops.MetricsServerConfig{}),
		generate(kops.ClusterAutoscalerConfig{}),
		generate(kops.AddonSpec{}),
		generate(kops.GossipConfig{}),
		generate(kops.ClusterSubnetSpec{}),
		generate(kops.TopologySpec{}),
		generate(kops.DNSControllerGossipConfig{}),
		generate(kops.EgressProxySpec{}),
		generate(kops.EtcdClusterSpec{},
			rename("Members", "Member"),
		),
		generate(kops.ContainerdConfig{}),
		generate(kops.PackagesConfig{}),
		generate(kops.DockerConfig{}),
		generate(kops.KubeDNSConfig{}),
		generate(kops.KubeAPIServerConfig{},
			nullable("AnonymousAuth"),
		),
		generate(kops.KubeControllerManagerConfig{}),
		generate(kops.CloudControllerManagerConfig{}),
		generate(kops.KubeSchedulerConfig{}),
		generate(kops.KubeProxyConfig{}),
		generate(kops.CloudConfiguration{}),
		generate(kops.ExternalDNSConfig{}),
		generate(kops.NetworkingSpec{}),
		generate(kops.AccessSpec{}),
		generate(kops.AuthenticationSpec{}),
		generate(kops.DNSAccessSpec{}),
		generate(kops.LoadBalancerAccessSpec{}),
		generate(kops.KopeioAuthenticationSpec{}),
		generate(kops.AwsAuthenticationSpec{}),
		generate(kops.OpenstackConfiguration{}),
		generate(kops.LeaderElectionConfiguration{}),
		generate(kops.AuthorizationSpec{}),
		generate(kops.NodeAuthorizationSpec{}),
		generate(kops.Assets{}),
		generate(kops.IAMSpec{}),
		generate(kops.AlwaysAllowAuthorizationSpec{}),
		generate(kops.RBACAuthorizationSpec{}),
		generate(kops.HTTPProxy{}),
		generate(kops.EtcdMemberSpec{}),
		generate(kops.EtcdBackupSpec{}),
		generate(kops.EtcdManagerSpec{}),
		generate(kops.NodeLocalDNSConfig{}),
		generate(kops.ClassicNetworkingSpec{}),
		generate(kops.KubenetNetworkingSpec{}),
		generate(kops.ExternalNetworkingSpec{}),
		generate(kops.CNINetworkingSpec{}),
		generate(kops.EnvVar{}),
		generate(kops.KopeioNetworkingSpec{}),
		generate(kops.WeaveNetworkingSpec{}),
		generate(kops.FlannelNetworkingSpec{}),
		generate(kops.CalicoNetworkingSpec{}),
		generate(kops.CanalNetworkingSpec{}),
		generate(kops.KuberouterNetworkingSpec{}),
		generate(kops.RomanaNetworkingSpec{}),
		generate(kops.AmazonVPCNetworkingSpec{}),
		generate(kops.CiliumNetworkingSpec{}),
		generate(kops.HubbleSpec{}),
		generate(kops.LyftVPCNetworkingSpec{}),
		generate(kops.GCENetworkingSpec{}),
		generate(kops.NodeAuthorizerSpec{}),
		generate(kops.OpenstackLoadbalancerConfig{}),
		generate(kops.OpenstackMonitor{}),
		generate(kops.OpenstackRouter{}),
		generate(kops.OpenstackBlockStorageConfig{}),
		generate(kops.BastionSpec{}),
		generate(kops.DNSSpec{}),
		generate(kops.BastionLoadBalancerSpec{}),
		generate(resources.InstanceGroup{},
			version(1),
			required("ClusterName", "Name"),
			exclude("Revision"),
		),
		generate(kops.InstanceGroupSpec{},
			noSchema(),
		),
		generate(kops.VolumeSpec{}),
		generate(kops.VolumeMountSpec{}),
		generate(kops.HookSpec{}),
		generate(kops.FileAssetSpec{}),
		generate(kops.KubeletConfigSpec{},
			nullable("AnonymousAuth"),
		),
		generate(kops.MixedInstancesPolicySpec{},
			nullable("OnDemandBase", "OnDemandAboveBase"),
		),
		generate(kops.UserData{}),
		generate(kops.LoadBalancer{}),
		generate(kops.IAMProfileSpec{}),
		generate(kops.RollingUpdate{}),
		generate(kops.ExecContainerAction{}),
		// 1.20
		generate(kops.AzureConfiguration{}),
		generate(kops.AWSEBSCSIDriver{}),
		generate(kops.NTPConfig{}),
		generate(kops.CertManagerConfig{}),
		generate(kops.AWSLoadBalancerControllerConfig{}),
		generate(kops.GossipConfigSecondary{}),
		generate(kops.LoadBalancerSubnetSpec{}),
		generate(kops.DNSControllerGossipConfigSecondary{}),
		generate(kops.OpenstackNetwork{}),
		// 1.21
		generate(kops.WarmPoolSpec{}),
		generate(kops.ServiceAccountIssuerDiscoveryConfig{}),
		generate(kops.SnapshotControllerConfig{}),
		generate(kops.ServiceAccountExternalPermission{}),
		generate(kops.AWSPermission{}),
	)
	// resources
	buildDoc(resources.Cluster{}, "docs/resources/", resourcesMap, "Resource", parser, resourceClusterHeader, resourceClusterFooter)
	buildDoc(resources.ClusterUpdater{}, "docs/resources/", resourcesMap, "Resource", parser, resourceClusterUpdaterHeader, "")
	buildDoc(resources.InstanceGroup{}, "docs/resources/", resourcesMap, "Resource", parser, resourceInstanceGroupHeader, resourceInstanceGroupFooter)
	// data sources
	buildDoc(datasources.ClusterStatus{}, "docs/data-sources/", dataSourcesMap, "DataSource", parser, dataClusterStatusHeader, "")
	buildDoc(datasources.KubeConfig{}, "docs/data-sources/", dataSourcesMap, "DataSource", parser, dataKubeConfigHeader, "")
	buildDoc(resources.Cluster{}, "docs/data-sources/", dataSourcesMap, "DataSource", parser, dataClusterHeader, "")
	buildDoc(resources.InstanceGroup{}, "docs/data-sources/", dataSourcesMap, "DataSource", parser, dataInstanceGroupHeader, "")
	// config
	buildDoc(config.Provider{}, "docs/guides/", configMap, "Config", parser, configProviderHeader, "")
}
