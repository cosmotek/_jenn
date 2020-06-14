package complexes

import (
	"github.com/cosmotek/_jenn/types"
	"github.com/cosmotek/_jenn/types/primitives"
)

type PhoneNumber struct{}

func (p PhoneNumber) PrimitiveRoot() types.Primitive {
	return primitives.String{
		Nullable:  false,
		MaxLength: 10,
	}
}
