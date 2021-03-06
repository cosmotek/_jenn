package cmd

import (
	"flag"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/cosmotek/_jenn/ir"
	"github.com/cosmotek/_jenn/templates"
	"github.com/spf13/cobra"
)

var outputs = map[string]string{
	"server/migrations/1.sql": templates.ProvisionDatabaseTemplate,
	"server/service.go":       templates.GoServiceTemplate,
	"server/controller.go":    templates.GoControllerTemplate,
}

var Build = &cobra.Command{
	Use:   "build",
	Short: "generate the application as a distributable artifact",
	Run: func(cmd *cobra.Command, args []string) {
		outputDir := *flag.String("o", "out/demo", "")

		model, err := ir.FromFile("ir/examples/shakenNotStirred.yaml")
		if err != nil {
			panic(err)
		}

		err = os.MkdirAll(outputDir, os.ModePerm)
		if err != nil {
			panic(err)
		}

		for filename, templateString := range outputs {
			templateData, err := templates.LoadFile(templateString)
			if err != nil {
				panic(err)
			}

			builtTemplate, err := templates.ToStr(model, model.Name, templateData)
			if err != nil {
				panic(err)
			}

			fullFilepath := filepath.Join(outputDir, filename)
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
		// 	Enums: []ir.Enum{
		// 		{
		// 			Description: "Type of drink",
		// 			Name:        "DrinkType",
		// 			Values: map[string]string{
		// 				"liquor": "a high-proof for mixing",
		// 				"wine":   "",
		// 				"beer":   "for partying",
		// 			},
		// 		},
		// 	},
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
		// 	// Forms: []ir.Form{
		// 	// 	{
		// 	// 		Fields: map[ir.FormAssignment]ir.Field{
		// 	// 			ir.FormAssignment{
		// 	// 				TypeName:  "user",
		// 	// 				FieldName: "firstName",
		// 	// 			}: {
		// 	// 				Name:   "firstName",
		// 	// 				TypeOf: types.Name,
		// 	// 			},
		// 	// 			ir.FormAssignment{
		// 	// 				TypeName:  "user",
		// 	// 				FieldName: "lastName",
		// 	// 			}: {
		// 	// 				Name:   "lastName",
		// 	// 				TypeOf: types.Name,
		// 	// 			},
		// 	// 			ir.Form{
		// 	// 				TypeName:  "user",
		// 	// 				FieldName: "phoneNumber",
		// 	// 			}: {
		// 	// 				Name:   "phoneNumber",
		// 	// 				TypeOf: types.Name,
		// 	// 			},
		// 	// 		},
		// 	// 	},
		// 	// },
		// }

		// file, err := os.Create("ir/examples/shakenNotStirredv2.yaml")
		// if err != nil {
		// 	panic(err)
		// }
		// defer file.Close()

		// err = yaml.NewEncoder(file).Encode(model)
		// if err != nil {
		// 	panic(err)
		// }

		// // err := hll.Compile(os.Args[1])
		// // if err != nil {
		// // 	panic(err)
		// // }
	},
}
