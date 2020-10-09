package typesys

// Complex types are custom types that eventually are reduced
// to primitives that can be added to generated SQL and Go.
type Complex interface {
	PrimitiveRoot() Primitive

	// generates templates for type definitions etc
	ExecTemplates() (string, error)

	// generates templates for precreate hooks
	ExecPrecreateFuncTemplates() (string, error)
}
