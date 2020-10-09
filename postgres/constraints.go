package postgres

import "fmt"

func NotNull(in Type) Type {
	return Type{
		Literal: fmt.Sprintf("%s NOT NULL", in.Literal),
	}
}
