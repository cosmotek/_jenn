package typesys

import (
	"errors"

	"github.com/cosmotek/_jenn/golang"
	"github.com/cosmotek/_jenn/postgres"
)

var (
	ErrValueUndefined = errors.New("value undefined")
	ErrFuncUndefined  = errors.New("function undefined")
)

// Primitive types define the root data structure that is encoded
// into generated code (SQL, Go etc).
type Primitive interface {
	SQLType() postgres.Type
	SQLDefault() *postgres.Value
	SQLIndexType() postgres.IndexType

	GoType() golang.Type
	GoTypeZeroValue() golang.Value

	// used to look up values such as "TIME_NOW"
	ResolveValue(name string) (golang.Value, error)

	// used to look up functions such now "Now()"
	ResolveFunc(name string) (golang.Expression, error)

	// used to validate input from forms
	Validators() []Validator

	// used to load custom functions
	Plugins() []Plugin
}

// TODO break into simpler set of interfaces
