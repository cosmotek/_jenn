package ir

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
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
	Name   string
	TypeOf FieldType
}

type FieldType struct {
	SQLTypeName  string
	GRPCTypeName string
}

var (
	UUID FieldType = FieldType{
		SQLTypeName:  "UUID PRIMARY KEY",
		GRPCTypeName: "string",
	}

	// name only allows alpha and a few special chars
	Name FieldType = FieldType{
		SQLTypeName:  "VARCHAR(64) NOT NULL",
		GRPCTypeName: "string",
	}

	// allows for alphanumeric and a small set of specials up to a specific size
	Username FieldType = FieldType{
		SQLTypeName:  "VARCHAR(128) NOT NULL",
		GRPCTypeName: "string",
	}
	// phone number, only allows US 10 digit (not including country code 1)
	PhoneNumber FieldType = FieldType{
		SQLTypeName:  "VARCHAR(10) NOT NULL",
		GRPCTypeName: "string",
	}
)

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

	return model, nil
}
