package main

import (
	"fmt"
	"go/ast"
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
	noSchema     bool
	exclude      sets.String
	asSets       sets.String
	rename       map[string]string
	required     sets.String
	computed     sets.String
	computedOnly sets.String
	forceNew     sets.String
	sensitive    sets.String
	suppressDiff sets.String
}

func newOptions() *options {
	return &options{
		exclude:      sets.NewString(),
		asSets:       sets.NewString(),
		rename:       make(map[string]string),
		required:     sets.NewString(),
		computed:     sets.NewString(),
		computedOnly: sets.NewString(),
		forceNew:     sets.NewString(),
		sensitive:    sets.NewString(),
		suppressDiff: sets.NewString(),
	}
}

func noSchema() func(o *options) {
	return func(o *options) {
		o.noSchema = true
	}
}

func exclude(excluded ...string) func(o *options) {
	return func(o *options) {
		o.exclude.Insert(excluded...)
	}
}

func asSets(s ...string) func(o *options) {
	return func(o *options) {
		o.asSets.Insert(s...)
	}
}

func rename(old, new string) func(o *options) {
	return func(o *options) {
		o.rename[old] = new
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

func forceNew(forceNew ...string) func(o *options) {
	return func(o *options) {
		o.forceNew.Insert(forceNew...)
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

func verifyFields(t reflect.Type, fields ...string) {
	for _, field := range fields {
		valid := false
		for i := 0; i < t.NumField(); i++ {
			f := t.Field(i)
			if f.Name == field {
				valid = true
			}
		}
		if !valid {
			panic(fmt.Sprintf("field %s is not part of struct %s", field, t.Name()))
		}
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

func getFields(t reflect.Type, flatten bool) []_field {
	var ret []_field
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if f.Anonymous && flatten {
			ret = append(ret, getFields(f.Type, flatten)...)
		} else {
			ret = append(ret, _field{
				StructField: t.Field(i),
				Owner:       t,
			})
		}
	}
	return ret
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

var mappings = map[string]string{
	"github.com/eddycharly/terraform-provider-kops/pkg/api/config":      "config",
	"github.com/eddycharly/terraform-provider-kops/pkg/api/datasources": "datasources",
	"github.com/eddycharly/terraform-provider-kops/pkg/api/kube":        "kube",
	"github.com/eddycharly/terraform-provider-kops/pkg/api/resources":   "resources",
	"github.com/eddycharly/terraform-provider-kops/pkg/api/utils":       "utils",
	"k8s.io/kops/pkg/apis/kops":                                         "kops",
}

func funcMap(baseType reflect.Type, optionsMap map[reflect.Type]*options, scope string, parser parser) template.FuncMap {
	dataSource := scope == "DataSource"
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
		"fields": getFields,
		"needsSchema": func(t reflect.Type) bool {
			return !optionsMap[t].noSchema
		},
		"schemaType": schemaType,
		"resourceComment": func(t reflect.Type) string {
			return getResourceComment(t.PkgPath(), t.Name(), parser.packs)
		},
		"attributeComment": func(f _field) string {
			return getAttributeComment(f.Owner.PkgPath(), f.Owner.Name(), f.Name, parser.packs)
		},
		"subResources": func(t reflect.Type) []reflect.Type {
			return getSubResources(t, map[reflect.Type]bool{}, func(in _field) bool {
				return optionsMap[in.Owner].exclude.Has(in.Name)
			})[1:]
		},
		"isExcluded": func(in _field) bool {
			return optionsMap[in.Owner].exclude.Has(in.Name)
		},
		"isSet": func(in _field) bool {
			return optionsMap[in.Owner].asSets.Has(in.Name)
		},
		"isRequired": func(in _field) bool {
			return optionsMap[in.Owner].required.Has(in.Name)
		},
		"isOptional": func(in _field) bool {
			if dataSource {
				return optionsMap[in.Owner].computed.Has(in.Name)
			}
			return !optionsMap[in.Owner].required.Has(in.Name) && !optionsMap[in.Owner].computedOnly.Has(in.Name)
		},
		"isComputed": func(in _field) bool {
			if optionsMap[in.Owner].required.Has(in.Name) {
				return false
			}
			return dataSource || (optionsMap[in.Owner].computed.Has(in.Name) || optionsMap[in.Owner].computedOnly.Has(in.Name))
		},
		"isSensitive": func(in _field) bool {
			return optionsMap[in.Owner].sensitive.Has(in.Name)
		},
		"forceNew": func(in _field) bool {
			return optionsMap[in.Owner].forceNew.Has(in.Name)
		},
		"suppressDiff": func(in _field) bool {
			return optionsMap[in.Owner].suppressDiff.Has(in.Name)
		},
		"fieldName": func(in _field) string {
			if optionsMap[in.Owner].rename[in.Name] != "" {
				return fieldName(optionsMap[in.Owner].rename[in.Name])
			}
			return fieldName(in.Name)
		},
		"schemaName": func(in string) string {
			return fieldName(in)
		},
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

func buildDoc(i interface{}, p string, optionsMap map[reflect.Type]*options, scope string, parser parser, header, footer string) {
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
{{- if isComputed . }} - (Computed){{ end }} - {{ template "type" .Type }}{{ if (attributeComment .) }} - {{ attributeComment . }}{{ end }}
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
{{- if isComputed . }} - (Computed){{ end }} - {{ template "type" .Type }}{{ if (attributeComment .) }} - {{ attributeComment . }}{{ end }}
{{- end -}}
{{- end }}
{{- end }}
{{ end }}
{{- end }}

%s
`, header, footer)

	tmpl := template.New("doc")
	funcMaps := []template.FuncMap{
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
	return &schema.Resource{
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
			{{- if isRequired . -}}
			Required
			{{- else if isOptional . -}}
			Optional
			{{- end -}}
			{{- if isComputed . -}}
			Computed
			{{- end -}}
			{{- if isSet . -}}
			Set
			{{- end -}}
			{{ template "schema" .Type }}
			{{- if $suppressDiff -}}){{- end -}}
			{{- if $forceNew -}}){{- end -}}
			{{- if $sensitive -}}){{- end -}}
			,
			{{- end -}}
			{{- end }}
		},
	}
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
func (in {{ .String }}) []map[string]interface{} {
	return []map[string]interface{}{ {{ mapping . }}Flatten{{ scope }}{{ .Name }}(in) }
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
		{{- if and (isPtr .Type) (isValueType .Type) (not (isRequired .)) }}
		if reflect.DeepEqual(in, reflect.Zero(reflect.TypeOf(in)).Interface()) {
			return nil
		}
		{{- end }}
		{{ if isSet . -}}
		return func (in interface{}) {{ .Type.String }} {
			return {{ template "expand" .Type }}
		}(in.(*schema.Set).List())
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
		return {{ template "flatten" .Type }}
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
	verifyFields(t, o.exclude.List()...)
	verifyFields(t, o.asSets.List()...)
	verifyFields(t, o.required.List()...)
	verifyFields(t, o.computed.List()...)
	verifyFields(t, o.computedOnly.List()...)
	verifyFields(t, o.computedOnly.List()...)
	verifyFields(t, o.forceNew.List()...)
	verifyFields(t, o.suppressDiff.List()...)
	for k := range o.rename {
		verifyFields(t, k)
	}
	return generated{
		t: t,
		o: o,
	}
}

func build(scope string, parser parser, g ...generated) map[reflect.Type]*options {
	o := map[reflect.Type]*options{}
	for _, gen := range g {
		o[gen.t] = gen.o
	}
	for _, gen := range g {
		funcMaps := []template.FuncMap{
			funcMap(gen.t, o, scope, parser),
			sprig.TxtFuncMap(),
		}
		buildSchema(gen.t, path.Join("pkg/schemas", mappings[gen.t.PkgPath()]), scope, funcMaps...)
	}
	return o
}

func buildImportsLookup(pack *packages.Package, file *ast.File) map[string]*packages.Package {
	out := make(map[string]*packages.Package)
	for _, i := range file.Imports {
		path := strings.ReplaceAll(i.Path.Value, "\"", "")
		if i.Name != nil {
			out[i.Name.Name] = pack.Imports[path]
		} else {
			out[pack.Imports[path].Name] = pack.Imports[path]
		}
	}
	return out
}

func (p *parser) parseStruct(pack *packages.Package, typeSpec *ast.TypeSpec, doc string, file *ast.File) _struct {
	ret := _struct{
		doc:    doc,
		fields: make(map[string]string),
	}
	structure, ok := typeSpec.Type.(*ast.StructType)
	if ok {
		importsLookup := buildImportsLookup(pack, file)
		for _, field := range structure.Fields.List {
			if len(field.Names) == 0 {
				switch t := field.Type.(type) {
				case *ast.Ident:
					ret.nested = append(ret.nested, _nested{
						pack:       pack.ID,
						structName: t.Name,
					})
				case *ast.SelectorExpr:
					x := t.X.(*ast.Ident)
					sel := t.Sel
					from := importsLookup[x.Name]
					ret.nested = append(ret.nested, _nested{
						pack:       from.ID,
						structName: sel.Name,
					})
				}
			} else {
				for _, name := range field.Names {
					ret.fields[name.Name] = field.Doc.Text()
				}
			}
		}
	}
	return ret
}

func (p *parser) parsePackage(pack *packages.Package) {
	p.packages[pack.PkgPath] = pack
	if _, ok := p.packs[pack.ID]; !ok {
		p.packs[pack.ID] = make(map[string]_struct)
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
							p.packs[pack.ID][typeSpec.Name.Name] = p.parseStruct(pack, typeSpec, genDecl.Doc.Text(), node)
						}
					}
				}
			}
		}
	}
}

type _field struct {
	reflect.StructField
	Owner reflect.Type
}

type _nested struct {
	pack       string
	structName string
}

type _struct struct {
	doc    string
	fields map[string]string
	nested []_nested
}

func (s _struct) lookup(in string, c map[string]map[string]_struct) string {
	if s, ok := s.fields[in]; ok && s != "" {
		return s
	}
	for _, n := range s.nested {
		if p, ok := c[n.pack]; ok {
			if s, ok := p[n.structName]; ok {
				out := s.lookup(in, c)
				if out != "" {
					return out
				}
			}
		}
	}
	return ""
}

type parser struct {
	packages map[string]*packages.Package
	packs    map[string]map[string]_struct
}

func main() {
	log.Println("loading packages...")
	cfg := packages.Config{
		Mode: packages.NeedName | packages.NeedFiles | packages.NeedCompiledGoFiles | packages.NeedImports | packages.NeedDeps | packages.NeedTypes | packages.NeedSyntax | packages.NeedTypesInfo,
	}
	packs, err := packages.Load(
		&cfg,
		"github.com/eddycharly/terraform-provider-kops/pkg/api/config",
		"github.com/eddycharly/terraform-provider-kops/pkg/api/datasources",
		"github.com/eddycharly/terraform-provider-kops/pkg/api/kube",
		"github.com/eddycharly/terraform-provider-kops/pkg/api/resources",
	)
	if err != nil {
		panic(err)
	}
	parser := parser{
		packages: make(map[string]*packages.Package),
		packs:    make(map[string]map[string]_struct),
	}
	for _, pack := range packs {
		parser.parsePackage(pack)
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
		generate(resources.ValidateOptions{}),
		generate(utils.ValidateOptions{},
			noSchema(),
		),
		generate(resources.ApplyOptions{}),
		generate(resources.Cluster{},
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
		generate(resources.InstanceGroup{},
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
	)
	configMap := build(
		"Config",
		parser,
		generate(config.Provider{},
			required("StateStore"),
		),
		generate(config.Aws{}),
		generate(config.AwsAssumeRole{}),
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
		generate(kube.Config{},
			noSchema(),
			sensitive("KubeBearerToken", "KubePassword", "CaCert", "ClientCert", "ClientKey"),
		),
		generate(resources.Cluster{},
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
		generate(kops.KubeAPIServerConfig{}),
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
		generate(kops.KubeletConfigSpec{}),
		generate(kops.MixedInstancesPolicySpec{}),
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
	)
	log.Println("generating docs...")
	// resources
	buildDoc(resources.Cluster{}, "docs/resources/", resourcesMap, "Resource", parser, resourceClusterHeader, resourceClusterFooter)
	buildDoc(resources.InstanceGroup{}, "docs/resources/", resourcesMap, "Resource", parser, resourceInstanceGroupHeader, resourceInstanceGroupFooter)
	buildDoc(resources.ClusterUpdater{}, "docs/resources/", resourcesMap, "Resource", parser, resourceClusterUpdaterHeader, "")
	// data sources
	buildDoc(datasources.ClusterStatus{}, "docs/data-sources/", dataSourcesMap, "DataSource", parser, dataClusterStatusHeader, "")
	buildDoc(datasources.KubeConfig{}, "docs/data-sources/", dataSourcesMap, "DataSource", parser, dataKubeConfigHeader, "")
	buildDoc(resources.Cluster{}, "docs/data-sources/", dataSourcesMap, "DataSource", parser, dataClusterHeader, "")
	buildDoc(resources.InstanceGroup{}, "docs/data-sources/", dataSourcesMap, "DataSource", parser, dataInstanceGroupHeader, "")
	// config
	buildDoc(config.Provider{}, "docs/guides/", configMap, "Config", parser, configProviderHeader, "")
	log.Println(resourcesMap)
	log.Println(configMap)
	log.Println(dataSourcesMap)
}
