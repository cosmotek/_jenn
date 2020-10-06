package types

import (
	"github.com/cosmotek/_jenn/golang"
	"github.com/cosmotek/_jenn/postgres"
)

// Primitive types define the root data structure that is encoded
// into generated code (SQL, Go etc).
type Primitive interface {
	SQLType() postgres.Type
	SQLDefault() postgres.Value
	SQLIndex(string) postgres.Index

	GoType() golang.Type
	GoTypeZeroValue() golang.Value

	Validators() []Validator

	// generates templates for type definitions etc
	ExecTemplates() (string, error)

	// generates templates for precreate hooks
	ExecPrecreateFuncTemplates() (string, error)
}
