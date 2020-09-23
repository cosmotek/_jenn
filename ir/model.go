package ir

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"

	irTypes "github.com/cosmotek/_jenn/ir/types"
	"github.com/cosmotek/_jenn/types"
	"github.com/cosmotek/_jenn/types/registry"
)

func getKeys(in map[string]string) []string {
	out := []string{}

	for key, _ := range in {
		out = append(out, key)
	}

	return out
}

type Structure struct {
	Name        string
	Description string

	Fields     []Field
	Namespaces []string
}

type Form struct {
	// TypeName -> FieldNames
	Fields map[string][]string
}

type Enum struct {
	Description string
	Name        string
	Values      map[string]string
	// TODO namespace enums
}

type ModelIR struct {
	Name        string
	Description string

	EnableUniversalArchiving bool
	EnableEventStreams       bool

	Types []Structure
	Enums []Enum
	Forms []Structure `yaml:"-"`

	FieldDerviedForms   map[string][]string `yaml:"forms"`
	GeneratedTypeBlocks []string            `yaml:"-"`
}

type Field struct {
	Name        string
	Description string

	TypeOf    irTypes.CanonicalName
	Primitive types.Primitive `yaml:"-"`

	Selector      bool
	SelectorTypes []string

	Optional     bool
	DefaultValue interface{}

	Namespaces []string
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

	reg := registry.Registry{}
	for _, enum := range model.Enums {
		reg.Add(enum.Name, getKeys(enum.Values)...)
	}

	for i, structure := range model.Types {
		for j, field := range structure.Fields {
			prim, typeBlocks, err := irTypes.ResolvePrimitive(reg, field.TypeOf)
			if err != nil {
				return ModelIR{}, err
			}

			model.Types[i].Fields[j].Primitive = prim
			model.GeneratedTypeBlocks = append(model.GeneratedTypeBlocks, typeBlocks...)
		}
	}

	for typeOf, fieldNames := range model.FieldDerviedForms {
		fields := []Field{}

		for _, structure := range model.Types {
			if structure.Name == typeOf {
				for _, field := range structure.Fields {
					for _, fieldName := range fieldNames {
						if field.Name == fieldName {
							fields = append(fields, field)
						}
					}
				}
			}
		}

		newForm := Structure{
			Name:   typeOf,
			Fields: fields,
		}

		model.Forms = append(model.Forms, newForm)
	}

	return model, nil
}
