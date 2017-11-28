package main

import (
	"github.com/aweisser/clean-arch-golang"

	"github.com/aweisser/clean-arch-golang/io/cobracli"
	"github.com/aweisser/clean-arch-golang/usecase/welcome"
)

func main() {
	// init a concrete shopping app
	factory.InitCliShop()

	// parse cli arguments
	cobracli.Execute(cobracli.Definition{
		Name: "go-shopping-cobra",
		Callback: func(name string) {
			// execute use case with concrete arguments
			welcome.TheCustomer(name)
		},
	})

}
