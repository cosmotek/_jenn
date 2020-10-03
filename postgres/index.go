package postgres

type IndexType struct {
	Literal            string
	SubTypeLiteral     string
	RequiresExtensions []Extension
}

var (
	BtreeIndex = IndexType{
		Literal:            "btree",
		RequiresExtensions: nil,
	}

	GinIndex = IndexType{
		Literal:            "gin",
		RequiresExtensions: nil,
	}

	TrigramIndex = IndexType{
		Literal:            "gin",
		SubTypeLiteral:     "gin_trgm_ops",
		RequiresExtensions: []Extension{TrigramExtension},
	}
)

type Index struct {
	Name   string
	TypeOf IndexType

	TableName string
	FieldName string

	// TODO support multifield indexes
}
