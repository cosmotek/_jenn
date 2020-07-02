package types

// Primitive types define the root data structure that is encoded
// into generated code (SQL, Go etc).
type Primitive interface {
	SQLType() string
	SQLIndexType(string) string

	GoType() string
	GoTypeZeroValueLiteral() string

	// generates templates for type definitions etc
	ExecTemplates() (string, error)

	// generates templates for precreate hooks
	ExecPrecreateFuncTemplates() (string, error)
}
