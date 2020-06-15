package templates

const (
	GoServiceTemplate = `
{{- $root := . }}
{{- range $j, $t := $root.Types }}
type {{ $t.Name | title }} {
	ID uuid.UUID
	{{- if $root.EnableUniversalArchiving }}
	Archived bool
	{{- end }}
	{{- $f := $t.Fields }}
	{{- range $i, $e := $f }}
	{{ $e.Name | title }} {{ $e.Primitive.GoType }}
	{{- end }}
}

func Archive{{ $t.Name | title }}(id string) error {
	return nil
}

func Get{{ $t.Name | title }}(id string) ({{ $t.Name | title }}, error) {
	return {{ $t.Name | title }}{}, nil
}

func Update{{ $t.Name | title }}(id string) ({{ $t.Name | title }}, error) {
	return {{ $t.Name | title }}{}, nil
}
{{ "\n" }}
{{- end }}
`
)
