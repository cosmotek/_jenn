package primitives

type DateTime struct{}

func (d DateTime) SQLType() string {
	return "timestamptz NOT NULL"
}

func (d DateTime) GoType() string {
	return "time.Time"
}

func (d DateTime) GRPCType() string {
	return "int64"
}
