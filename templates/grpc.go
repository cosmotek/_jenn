package templates

type GRPCTypeDefinitionTemplate string

const (
	MessageTemplate GRPCTypeDefinitionTemplate = `
message {{ .Name | title }} {
	{{- range $i, $e := .Fields }}
	{{$e.TypeOf.GRPCTypeName}} {{$e.Name}} = {{$i | incr}};
	{{- end}}
}
`
)
