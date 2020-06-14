package types

// Primitive types define the root data structure that is encoded
// into generated code (SQL, Go etc).
type Primitive interface {
	SQLType() string
	GoType() string
	GRPCType() string
}
