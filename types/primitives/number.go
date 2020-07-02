package primitives

import "fmt"

type Number struct {
	Nullable              bool
	DisableNegativeValues bool
	EnableDecimal         bool
	MaxValue              int64
}

func (s Number) SQLType() string {
	typeOf := "smallint"
	max := s.MaxValue

	// check if negative
	if max < 0 {
		max *= -1
	}

	if max > 32767 {
		typeOf = "integer"
	}

	if max > 2147483647 {
		typeOf = "bigint"
	}

	if s.EnableDecimal {
		typeOf = "float"
	}

	opts := ""
	if !s.Nullable {
		opts += " NOT NULL"
	}

	return fmt.Sprintf("%s%s", typeOf, opts)
}

func (s Number) GoType() string {
	// todo handle complex int sizes and nullability

	return "int"
}

func (s Number) GoTypeZeroValueLiteral() string {
	return "0"
}

func (s Number) ExecTemplates() (string, error) {
	return "", nil
}

func (s Number) ExecPrecreateFuncTemplates() (string, error) {
	return "", nil
}

func (d Number) SQLIndexType(fieldName string) string {
	return fmt.Sprintf("BTREE (\"%s\")", fieldName)
}
