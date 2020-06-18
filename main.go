package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/cosmotek/_jenn/ir"
	"github.com/cosmotek/_jenn/templates"

	"github.com/vmihailenco/msgpack/v5"
)

func main() {
	model, err := ir.FromFile("ir/examples/shakenNotStirred.yaml")
	if err != nil {
		panic(err)
	}

	provisionDatabaseTmpl, err := templates.LoadFile(templates.ProvisionDatabaseTemplate)
	if err != nil {
		panic(err)
	}

	protoStr, err := templates.ToStr(model, provisionDatabaseTmpl)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("out/provision.sql", []byte(protoStr), os.ModePerm)
	if err != nil {
		panic(err)
	}

	goServiceTmpl, err := templates.LoadFile(templates.GoServiceTemplate)
	if err != nil {
		panic(err)
	}

	protoStr, err = templates.ToStr(model, goServiceTmpl)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("out/service.go", []byte(protoStr), os.ModePerm)
	if err != nil {
		panic(err)
	}

	jsClientTmpl, err := templates.LoadFile(templates.JSClientTemplate)
	if err != nil {
		panic(err)
	}

	protoStr, err = templates.ToStr(model, jsClientTmpl)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("out/client.js", []byte(protoStr), os.ModePerm)
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/msg", func(res http.ResponseWriter, req *http.Request) {
		body := map[string]interface{}{}
		err := msgpack.NewDecoder(req.Body).Decode(&body)
		if err != nil {
			http.Error(res, err.Error(), 500)
			fmt.Println(err.Error())
			return
		}

		fmt.Println("got", body)

		err = msgpack.NewEncoder(res).Encode(map[string]interface{}{
			"name": "seth",
			"age":  20,
			"male": true,
			"dob":  "07/25/99",
		})
		if err != nil {
			http.Error(res, err.Error(), 500)
			return
		}
	})

	http.ListenAndServe(":5000", nil)

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
