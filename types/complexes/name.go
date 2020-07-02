package complexes

import (
	"github.com/cosmotek/_jenn/types"
	"github.com/cosmotek/_jenn/types/primitives"
)

type Name struct{}

func (n Name) PrimitiveRoot() types.Primitive {
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
