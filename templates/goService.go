package templates

const (
	GoServiceTemplate = `
{{- range $j, $t := .Types }}
type {{ $t.Name | title }} {
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
{{- end }}
`
)
