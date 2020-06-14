package templates

type PostgresQueryTemplate string

const (
	InsertTemplate PostgresQueryTemplate = `
INSERT INTO {{ .Name }} (
	{{- $f := .Fields }}
	{{- range $i, $e := .Fields }}
	{{$e.Name}}{{if last $i $f}}{{else}},{{end}}
	{{- end}}
) VALUES (
	{{- $f := .Fields }}
	{{- range $i, $e := .Fields }}
	{{"$"}}{{$i | incr}}{{if last $i $f}}{{else}},{{end}}
	{{- end}}
)
`

	CreateTableTemplate PostgresQueryTemplate = `
CREATE TABLE {{ .Name }} (
	{{- $f := .Fields }}
	{{- range $i, $e := .Fields }}
	{{$e.Name}} {{$e.TypeOf.SQLTypeName}}{{if last $i $f}}{{else}},{{end}}
	{{- end}}
);
`
)
