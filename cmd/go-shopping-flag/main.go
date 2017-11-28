package main

import (
	"github.com/aweisser/clean-arch-golang"

	"github.com/aweisser/clean-arch-golang/io/flagcli"
	"github.com/aweisser/clean-arch-golang/usecases/welcome"
)

func main() {
	// init a concrete shopping app
	factory.InitCliShop()

	// parse cli arguments
	flagcli.Execute(flagcli.Definition{
		Callback: func(name string) {
			// execute use case with concrete arguments
			welcome.TheCustomer(name)
		},
	})
}
