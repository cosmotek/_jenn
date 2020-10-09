package postgres

import "fmt"

type Type struct {
	Literal string
}

var (
	TypeTimestamp = Type{
		Literal: "timestamp",
	}

	TypeTimestampWithTimezone = Type{
		Literal: "timestamptz",
	}

	TypeText = Type{
		Literal: "TEXT",
	}

	TypeBoolean = Type{
		Literal: "BOOLEAN",
	}

	TypeSmallInt = Type{
		Literal: "smallint",
	}

	TypeInteger = Type{
		Literal: "integer",
	}

	TypeBigInt = Type{
		Literal: "bigint",
	}

	TypeFloat = Type{
		Literal: "float",
	}
)

func VarChar(len uint64) Type {
	return Type{
		Literal: fmt.Sprintf("VARCHAR(%d)", len),
	}
}
