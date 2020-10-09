package golang

import "fmt"

type Type struct {
	Imports []Import
	Literal string
}

var (
	Time = Type{
		Imports: []Import{TimePackage},
		Literal: "time.Time",
	}

	Int8 = Type{
		Literal: "int8",
	}

	Int16 = Type{
		Literal: "int16",
	}

	Int32 = Type{
		Literal: "int32",
	}

	Int64 = Type{
		Literal: "int64",
	}

	Float64 = Type{
		Literal: "float64",
	}

	String = Type{
		Literal: "string",
	}
)

func PointerTo(in Type) Type {
	return Type{
		Imports: in.Imports,
		Literal: fmt.Sprintf("*%s", in.Literal),
	}
}
