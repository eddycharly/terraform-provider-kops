# kops_{{ schemaName .Name | snakecase }}

{{ header }}

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

{{ footer }}
