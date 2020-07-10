package hll

import (
	"fmt"
	"os"

	"github.com/cosmotek/_jenn/hll/parser"
	"gopkg.in/yaml.v2"
)

func Compile(filename string) error {
	model, err := parser.Parse(filename)

	file, err := os.Create(fmt.Sprintf("%s.yaml", model.Name))
	if err != nil {
		return err
	}
	defer file.Close()

	err = yaml.NewEncoder(file).Encode(model)
	if err != nil {
		return err
	}

	return nil
}
