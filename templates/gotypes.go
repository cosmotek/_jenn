package templates

type GoTypeDefinitionTemplate string

const (
	TypeTemplate GoTypeDefinitionTemplate = `
	type {{ .Name | title }} struct{
		{{- range $i, $e := .Fields }}
		{{$e.Name | title}} {{ sprintf "db:%v" ($e.Name | quotes) | backticks }}
		{{- end}}
	}
	`
)
