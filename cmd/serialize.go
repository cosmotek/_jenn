package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/cosmotek/_jenn/hll"

	pjson "github.com/hokaccha/go-prettyjson"
	"github.com/spf13/cobra"
)

var (
	outputRawJSON bool
	outputFile    string
)

var Serialize = &cobra.Command{
	Use:        "serialize",
	Short:      "validate schema and generate IR",
	Args:       cobra.ExactArgs(1),
	ArgAliases: []string{"schema file"},
	Run: func(cmd *cobra.Command, args []string) {
		irData, err := hll.FromFile(args[0])
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		var jsonData []byte
		if outputRawJSON {
			jsonData, err = json.MarshalIndent(irData, "", "\t")
		} else {
			jsonData, err = pjson.Marshal(irData)
		}
		if err != nil {
			fmt.Printf("%s:%s\n", args[0], err.Error())
			os.Exit(1)
		}

		if outputFile == "" {
			fmt.Println(string(jsonData))
		} else {
			err = ioutil.WriteFile(outputFile, jsonData, os.ModePerm)
			if err != nil {
				fmt.Printf("%s:%s\n", args[0], err.Error())
				os.Exit(1)
			}
		}
	},
}

func init() {
	Serialize.Flags().BoolVarP(&outputRawJSON, "output-raw-json", "r", false, "output IR as raw JSON")
	Serialize.Flags().StringVarP(&outputFile, "output-file", "o", "", "output to file")
}
