package types

// Complex types are custom types that eventually are reduced
// to primitives that can be added to generated SQL and Go.
type Complex interface {
	PrimitiveRoot() Primitive
}
