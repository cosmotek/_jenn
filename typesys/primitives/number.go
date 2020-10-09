package primitives

import (
	"github.com/cosmotek/_jenn/golang"
	"github.com/cosmotek/_jenn/postgres"
	"github.com/cosmotek/_jenn/typesys"
)

type Number struct {
	Nullable              bool
	DisableNegativeValues bool
	EnableDecimal         bool
	MaxValue              int64
}

func (d Number) SQLType() postgres.Type {
	typeOf := postgres.TypeSmallInt
	max := d.MaxValue

	// check if negative
	if max < 0 {
		max *= -1
	}

	if max > 32767 {
		typeOf = postgres.TypeInteger
	}

	if max > 2147483647 {
		typeOf = postgres.TypeBigInt
	}

	if d.EnableDecimal {
		typeOf = postgres.TypeFloat
	}

	if !d.Nullable {
		return postgres.NotNull(typeOf)
	}

	return typeOf
}

func (d Number) SQLIndexType() postgres.IndexType {
	return postgres.BtreeIndex
}

func (d Number) SQLDefault() *postgres.Value {
	return nil
}

func (d Number) GoType() golang.Type {
	typeOf := golang.Int8
	max := d.MaxValue

	// TODO add unsigned int support
	// check if negative
	if max < 0 {
		max *= -1
	}

	if max > 32767 {
		typeOf = golang.Int16
	}

	if max > 2147483647 {
		typeOf = golang.Int64
	}

	if d.EnableDecimal {
		typeOf = golang.Float64
	}

	if d.Nullable {
		return golang.PointerTo(typeOf)
	}

	return typeOf
}

func (d Number) GoTypeZeroValue() golang.Value {
	return golang.Zero
}

func (d Number) ResolveValue(name string) (golang.Value, error) {
	return golang.Value{}, typesys.ErrValueUndefined
}

func (d Number) ResolveFunc(name string) (golang.Expression, error) {
	return golang.Expression(""), typesys.ErrFuncUndefined
}

func (d Number) Validators() []typesys.Validator {
	return nil
}

func (d Number) Plugins() []typesys.Plugin {
	return nil
}
