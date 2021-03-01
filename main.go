package main

import (
	"fmt"
	"os"

	"github.com/cosmotek/_jenn/cmd"
	"github.com/spf13/cobra"
)

var root = &cobra.Command{
	Use: "jenn",
}

func main() {
	root.AddCommand(cmd.Build, cmd.Serialize)

	err := root.Execute()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
