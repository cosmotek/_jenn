package templates

const PostgresTableCreationTemplate = `
{{- range $j, $t := .Types }}
CREATE TABLE {{ $t.Name | quotes }} (
	{{- $f := $t.Fields }}
	{{- range $i, $e := $f }}
	{{ $e.Name | quotes }} {{ $e.Primitive.SQLType }}{{ if last $i $f }}{{ else }},{{ end }}
	{{- end }}
);
{{ "\n" }}
{{- $f := $t.Fields }}
{{- range $i, $e := $f }}
{{- if $e.Selector }}
CREATE INDEX CONCURRENTLY index_{{ $t.Name }}_on_{{ $e.Name }}_trigram
ON {{ $t.Name | quotes }}
USING gin ({{ $e.Name | quotes }} gin_trgm_ops);
{{ "\n" }}
{{- end }}
{{- end }}
{{- end }}
`
