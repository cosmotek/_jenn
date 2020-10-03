package postgres

type Extension struct {
	Literal string
}

var (
	TrigramExtension = Extension{
		Literal: "pg_trgm",
	}
)
