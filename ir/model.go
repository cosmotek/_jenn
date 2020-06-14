package ir

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"

	irTypes "github.com/cosmotek/_jenn/ir/types"
	"github.com/cosmotek/_jenn/types"
)

type Structure struct {
	Name   string
	Fields []Field
}

type ModelIR struct {
	Name  string
	Types []Structure
}

type Field struct {
	Name      string
	TypeOf    irTypes.CanonicalName
	Primitive types.Primitive `yaml:"-"`

	Selector bool
	Optional bool
}

func FromFile(filename string) (ModelIR, error) {
	fileBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return ModelIR{}, err
	}

	model := ModelIR{}
	err = yaml.Unmarshal(fileBytes, &model)
	if err != nil {
		return ModelIR{}, err
	}

	for i, structure := range model.Types {
		for j, field := range structure.Fields {
			prim, err := irTypes.ResolvePrimitive(field.TypeOf)
			if err != nil {
				return ModelIR{}, err
			}

			model.Types[i].Fields[j].Primitive = prim
		}
	}

	return model, nil
}
