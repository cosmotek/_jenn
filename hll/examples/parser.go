package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/cosmotek/_jenn/hll/lexer"
	"github.com/cosmotek/_jenn/hll/parser"
)

func main() {
	data, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}

	tokens := lexer.New(string(data)).Tokens()
	psr := parser.New(tokens...)

	output, err := psr.Parse()
	if err != nil {
		fmt.Println(err)
	}

	// jsonStr, err := psr.JSON()
	// if err != nil {
	// 	panic(err)
	// }

	json.NewEncoder(os.Stdout).Encode(output)
}
