package postgres

type Value struct {
	Literal string
}

var (
	EmptyString = Value{
		Literal: "''",
	}

	Null = Value{
		Literal: "NULL",
	}
)
