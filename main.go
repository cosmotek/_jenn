package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/cosmotek/_jenn/ir"
	"github.com/cosmotek/_jenn/templates"
)

var outputs = map[string]string{
	"server/migrations/1.sql": templates.ProvisionDatabaseTemplate,
	"server/service.go":       templates.GoServiceTemplate,
	"server/controller.go":    templates.GoControllerTemplate,
	"clients/client.js":       templates.JSClientTemplate,
	"clients/client.dart":     templates.DartClientTemplate,
}

func main() {
	model, err := ir.FromFile("ir/examples/shakenNotStirred.yaml")
	if err != nil {
		panic(err)
	}

	err = os.MkdirAll(fmt.Sprintf("out/%s", model.Name), os.ModePerm)
	if err != nil {
		panic(err)
	}

	for filename, templateString := range outputs {
		templateData, err := templates.LoadFile(templateString)
		if err != nil {
			panic(err)
		}

		builtTemplate, err := templates.ToStr(model, templateData)
		if err != nil {
			panic(err)
		}

		fullFilepath := filepath.Join(fmt.Sprintf("out/%s", model.Name), filename)
		err = os.MkdirAll(filepath.Dir(fullFilepath), os.ModePerm)
		if err != nil {
			panic(err)
		}

		err = ioutil.WriteFile(fullFilepath, []byte(builtTemplate), os.ModePerm)
		if err != nil {
			panic(err)
		}
	}

	// model := ir.ModelIR{
	// 	Name: "shakenNotStirred",
	// 	Types: []ir.Structure{
	// 		{
	// 			Name: "user",
	// 			Fields: []ir.Field{
	// 				{
	// 					Name:   "lastName",
	// 					TypeOf: types.Name,
	// 				},
	// 				{
	// 					Name:   "lastName",
	// 					TypeOf: types.Name,
	// 				},
	// 				{
	// 					Name:   "joinedAt",
	// 					TypeOf: types.DateTime,
	// 				},
	// 				{
	// 					Name:   "phoneNumber",
	// 					TypeOf: types.PhoneNumber,
	// 				},
	// 			},
	// 		},
	// 	},
	// 	Forms: []ir.Form{
	// 		{
	// 			Fields: map[ir.FormAssignment]ir.Field{
	// 				ir.FormAssignment{
	// 					TypeName:  "user",
	// 					FieldName: "firstName",
	// 				}: {
	// 					Name:   "firstName",
	// 					TypeOf: types.Name,
	// 				},
	// 				ir.FormAssignment{
	// 					TypeName:  "user",
	// 					FieldName: "lastName",
	// 				}: {
	// 					Name:   "lastName",
	// 					TypeOf: types.Name,
	// 				},
	// 				ir.FormAssignment{
	// 					TypeName:  "user",
	// 					FieldName: "phoneNumber",
	// 				}: {
	// 					Name:   "phoneNumber",
	// 					TypeOf: types.Name,
	// 				},
	// 			},
	// 		},
	// 	},
	// }

	// file, err := os.Create("ir/examples/shakenNotStirred.yaml")
	// if err != nil {
	// 	panic(err)
	// }
	// defer file.Close()

	// err = yaml.NewEncoder(file).Encode(model)
	// if err != nil {
	// 	panic(err)
	// }
}
