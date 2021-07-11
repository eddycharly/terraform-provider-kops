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
