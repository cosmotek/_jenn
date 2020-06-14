package main

import (
	"io/ioutil"
	"os"

	"github.com/cosmotek/_jenn/ir"
	"github.com/cosmotek/_jenn/templates"
)

func main() {
	model, err := ir.FromFile("ir/examples/shakenNotStirred.yaml")
	if err != nil {
		panic(err)
	}

	protoStr, err := templates.ToStr(model, string(templates.MessageTemplate))
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("out/shakenNotStirred.proto", []byte(protoStr), os.ModePerm)
	if err != nil {
		panic(err)
	}
}
