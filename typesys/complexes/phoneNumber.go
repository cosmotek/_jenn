package complexes

import (
	"github.com/cosmotek/_jenn/typesys"
	"github.com/cosmotek/_jenn/typesys/primitives"
)

type PhoneNumber struct{}

func (p PhoneNumber) PrimitiveRoot() typesys.Primitive {
	return primitives.String{
		Nullable:  false,
		MaxLength: 10,
	}
}

func (p PhoneNumber) ExecTemplates() (string, error) {
	return "", nil
}

func (p PhoneNumber) ExecPrecreateFuncTemplates() (string, error) {
	return "", nil
}
