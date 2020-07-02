package primitives

import "fmt"

type DateTime struct{}

func (d DateTime) SQLType() string {
	return "timestamptz NOT NULL"
}

func (d DateTime) SQLIndexType(fieldName string) string {
	return fmt.Sprintf("BTREE (\"%s\")", fieldName)
}

func (d DateTime) GoType() string {
	return "time.Time"
}

func (d DateTime) GoTypeZeroValueLiteral() string {
	return "time.Time{}"
}

func (d DateTime) ExecTemplates() (string, error) {
	return "", nil
}

func (d DateTime) ExecPrecreateFuncTemplates() (string, error) {
	return "", nil
}
