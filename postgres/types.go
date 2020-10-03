package postgres

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
