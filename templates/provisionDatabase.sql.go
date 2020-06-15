package templates

const PostgresTableCreationTemplate = `
{{- $root := . }}
{{- range $j, $t := $root.Types }}
CREATE TABLE {{ $t.Name | quotes }} (
	"id" UUID PRIMARY KEY,
	{{- if $root.EnableUniversalArchiving }}
	"_archived" BOOLEAN NOT NULL DEFAULT FALSE,
	{{- end }}
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
