package types

import (
	"fmt"

	"github.com/cosmotek/_jenn/types"
	"github.com/cosmotek/_jenn/types/complexes"
	"github.com/cosmotek/_jenn/types/primitives"
)

type CanonicalName string

const (
	// Primitives
	String   CanonicalName = "string"
	DateTime CanonicalName = "datetime"

	// Complexes
	Name        CanonicalName = "name"
	PhoneNumber CanonicalName = "phoneNumber"
)

func ResolvePrimitive(typeOf CanonicalName) (types.Primitive, error) {
	switch typeOf {
	case String:
		return primitives.String{}, nil
	case DateTime:
		return primitives.DateTime{}, nil
	case Name:
		return complexes.Name{}.PrimitiveRoot(), nil
	case PhoneNumber:
		return complexes.PhoneNumber{}.PrimitiveRoot(), nil
	default:
		return nil, fmt.Errorf("unable to resolve type: '%s'", string(typeOf))
	}
}
