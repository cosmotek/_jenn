package primitives

import "fmt"

type String struct {
	Nullable  bool
	MaxLength uint64
}

func (s String) SQLType() string {
	length := uint64(512)
	if s.MaxLength > 0 {
		length = s.MaxLength
	}

	opts := ""
	if !s.Nullable {
		opts += " NOT NULL"
	}

	return fmt.Sprintf("VARCHAR(%d)%s", length, opts)
}

func (s String) GoType() string {
	if s.Nullable {
		return "*string"
	}

	return "string"
}

func (s String) GRPCType() string {
	return "string"
}

func (s String) GoTypeZeroValueLiteral() string {
	return "\"\""
}
