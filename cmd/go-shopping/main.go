package main

import (
	"github.com/aweisser/clean-arch-golang/io/cli"

	"github.com/aweisser/clean-arch-golang"

	"github.com/aweisser/clean-arch-golang/usecases/welcome"
)

func main() {
	// init a concrete shopping app
	factory.InitCliShop()

	// parse command line
	customerName := cli.ParseNameFlag()

	// execute use case
	welcome.TheCustomer(customerName)
}
