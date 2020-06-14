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
