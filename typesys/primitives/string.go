package primitives

import (
	"github.com/cosmotek/_jenn/golang"
	"github.com/cosmotek/_jenn/postgres"
	"github.com/cosmotek/_jenn/typesys"
)

type String struct {
	Nullable  bool
	MaxLength uint64
}

func (s String) SQLType() postgres.Type {
	if s.MaxLength > 0 {
		if !s.Nullable {
			return postgres.NotNull(postgres.VarChar(s.MaxLength))
		}
	}

	if !s.Nullable {
		return postgres.NotNull(postgres.TypeText)
	}

	return postgres.TypeText
}

func (String) SQLIndexType() postgres.IndexType {
	return postgres.GinIndex
}

func (s String) GoType() golang.Type {
	if s.Nullable {
		return golang.PointerTo(golang.String)
	}

	return golang.String
}

func (s String) GoTypeZeroValue() golang.Value {
	if s.Nullable {
		return golang.Nil
	}

	return golang.EmptyString
}

func (s String) SQLDefault() *postgres.Value {
	if s.Nullable {
		return &postgres.Null
	}

	return &postgres.EmptyString
}

func (String) ResolveValue(name string) (golang.Value, error) {
	return golang.Value{}, typesys.ErrValueUndefined
}

func (String) ResolveFunc(name string) (golang.Expression, error) {
	return golang.Expression(""), typesys.ErrFuncUndefined
}

func (String) Validators() []typesys.Validator {
	return nil
}

func (String) Plugins() []typesys.Plugin {
	return nil
}
