package complexes

import (
	"github.com/cosmotek/_jenn/typesys"
	"github.com/cosmotek/_jenn/typesys/primitives"
)

type Name struct{}

func (n Name) PrimitiveRoot() typesys.Primitive {
	return primitives.String{
		Nullable:  false,
		MaxLength: 64,
	}
}

func (n Name) ExecTemplates() (string, error) {
	return "", nil
}

func (n Name) ExecPrecreateFuncTemplates() (string, error) {
	return "", nil
}
