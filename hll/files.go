package hll

import (
	"fmt"
	"io/ioutil"

	"github.com/cosmotek/_jenn/hll/lexer"
	"github.com/cosmotek/_jenn/hll/parser"
	"github.com/cosmotek/_jenn/ir"
)

func FromFile(filename string) (ir.ModelIR, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return ir.ModelIR{}, err
	}

	tokens := lexer.New(string(data)).Tokens()
	output, err := parser.New(tokens...).Parse()
	if err != nil {
		return ir.ModelIR{}, fmt.Errorf("%s:%s\n", filename, err.Error())
	}

	return output, nil
}
