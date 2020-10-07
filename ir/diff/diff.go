package diff

import "github.com/cosmotek/_jenn/ir"

type ChangeMode int

const (
	Addition ChangeMode = iota
	Subtraction
)

type TypeChange struct {
	TypeName string
}

type EnumChange struct {
	TypeName string
	ChangeMode
}

type Result struct {
	TypeChanges []TypeChange
	EnumChanges []EnumChange
}

func Compare(a ir.ModelIR, b ir.ModelIR) (Result, error) {
	return Result{}, nil
}
