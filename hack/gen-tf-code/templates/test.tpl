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