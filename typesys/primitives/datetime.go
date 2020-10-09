package primitives

import (
	"github.com/cosmotek/_jenn/golang"
	"github.com/cosmotek/_jenn/postgres"
	"github.com/cosmotek/_jenn/typesys"
)

type DateTime struct{}

func (d DateTime) SQLType() postgres.Type {
	return postgres.NotNull(postgres.TypeTimestampWithTimezone)
}

func (d DateTime) SQLIndexType() postgres.IndexType {
	return postgres.BtreeIndex
}

func (d DateTime) SQLDefault() *postgres.Value {
	return nil
}

func (d DateTime) GoType() golang.Type {
	return golang.Time
}

func (d DateTime) GoTypeZeroValue() golang.Value {
	return golang.EmptyTime
}

func (d DateTime) ResolveValue(name string) (golang.Value, error) {
	switch name {
	case "CURRENT_TIMESTAMP":
		return golang.TimeNow, nil
	default:
		return golang.Value{}, typesys.ErrValueUndefined
	}
}

func (d DateTime) ResolveFunc(name string) (golang.Expression, error) {
	return golang.Expression(""), typesys.ErrFuncUndefined
}

func (d DateTime) Validators() []typesys.Validator {
	return nil
}

func (d DateTime) Plugins() []typesys.Plugin {
	return nil
}
