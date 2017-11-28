package cobracli

import (
	"fmt"
	"os"

	"github.com/aweisser/clean-arch-golang/io/cobracli/cmd"
)

// Definition of the cli tool
type Definition struct {
	Name     string
	Callback func(name string)
}

// Execute uses cobra to setup and execute the cli
func Execute(definition Definition) {
	cmd.RootCallback = definition.Callback
	cmd.RootCmd.Use = definition.Name
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
