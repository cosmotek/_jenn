package types

import (
	"fmt"

	types "github.com/cosmotek/_jenn/typesys"
	"github.com/cosmotek/_jenn/typesys/complexes"
	"github.com/cosmotek/_jenn/typesys/primitives"
)

type CanonicalName string

const (
	// Primitives
	String   CanonicalName = "string"
	DateTime CanonicalName = "datetime"
	Number   CanonicalName = "number"

	// Complexes
	Name        CanonicalName = "name"
	PhoneNumber CanonicalName = "phoneNumber"
)

var prims = map[CanonicalName]types.Primitive{
	String:   primitives.String{},
	DateTime: primitives.DateTime{},
	Number:   primitives.Number{},
}

var comps = map[CanonicalName]types.Complex{
	Name:        complexes.Name{},
	PhoneNumber: complexes.PhoneNumber{},
}

func ResolvePrimitive(typeOf CanonicalName) (types.Primitive, []string, error) {
	// primitive, ok := prims[typeOf]
	// if ok {
	// 	typeTemplates, err := primitive.ExecTemplates()
	// 	if err != nil {
	// 		return nil, nil, err
	// 	}

	// 	blocks := []string{}
	// 	if typeTemplates != "" {
	// 		blocks = append(blocks, typeTemplates)
	// 	}

	// 	return primitive, blocks, nil

	// }

	complex, ok := comps[typeOf]
	if ok {
		typeTemplates, err := complex.ExecTemplates()
		if err != nil {
			return nil, nil, err
		}

		blocks := []string{}
		if typeTemplates != "" {
			blocks = append(blocks, typeTemplates)
		}

		return complex.PrimitiveRoot(), blocks, nil
	}

	// enum, exists := reg.FindEnum(string(typeOf))
	// if !exists {
	// 	return nil, nil, fmt.Errorf("unable to resolve type or enum: '%s'", string(typeOf))
	// }

	// typeTemplates, err := enum.ExecTemplates()
	// if err != nil {
	// 	return nil, nil, err
	// }

	// return enum, []string{typeTemplates}, nil
	return nil, nil, fmt.Errorf("unable to resolve type or enum: '%s'", string(typeOf))
}
