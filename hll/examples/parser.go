package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/cosmotek/_jenn/hll/lexer"
	"github.com/cosmotek/_jenn/hll/parser"

	pjson "github.com/hokaccha/go-prettyjson"
)

func main() {
	filename := os.Args[1]
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	tokens := lexer.New(string(data)).Tokens()
	psr := parser.New(tokens...)

	output, err := psr.Parse()
	if err != nil {
		fmt.Printf("%s:%s\n", filename, err.Error())
		os.Exit(1)
	}

	jsonData, err := pjson.Marshal(output)
	if err != nil {
		fmt.Printf("%s:%s\n", filename, err.Error())
		os.Exit(1)
	}

	fmt.Println(string(jsonData))
}
