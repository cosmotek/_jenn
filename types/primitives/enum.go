package primitives

import (
	"fmt"
	"strings"

	"github.com/cosmotek/_jenn/templates"
)

const enumTmpl = `
{{- $root := . }}
type {{ $root.Name | title }} string

const (
{{- range $i, $val := $root.Values }}
	{{ ($val | lower) | title }} {{ $root.Name | title }} = "{{ ($val | lower) | title }}"
{{- end }}
)
`

type Enum struct {
	Nullable bool
	Name     string
	Values   []string
}

func (s Enum) SQLType() string {
	length := uint64(0)
	for _, val := range s.Values {
		if uint64(len(val)) > length {
			length = uint64(len(val))
		}
	}

	opts := ""
	if !s.Nullable {
		opts += " NOT NULL"
	}

	return fmt.Sprintf("VARCHAR(%d)%s", length, opts)
}

func (d Enum) SQLIndexType(fieldName string) string {
	return fmt.Sprintf("BTREE (\"%s\")", fieldName)
}

func (s Enum) GoType() string {
	if s.Nullable {
		return fmt.Sprintf("*%s", strings.Title(s.Name))
	}

	return strings.Title(s.Name)
}

func (s Enum) GoTypeZeroValueLiteral() string {
	return "\"\""
}

func (s Enum) ExecTemplates() (string, error) {
	return templates.ToStr(s, s.Name, enumTmpl)
}

func (s Enum) ExecPrecreateFuncTemplates() (string, error) {
	return "", nil
}
