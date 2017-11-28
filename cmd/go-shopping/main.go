package main

import (
	"flag"

	"github.com/aweisser/clean-arch-golang/domain/shop"
	"github.com/aweisser/clean-arch-golang/io/stdout"
	"github.com/aweisser/clean-arch-golang/usecases/welcome"
)

func main() {
	// parse cli flags
	var customerName string
	flag.StringVar(&customerName, "name", "John Doe", "enter your name")
	flag.Parse()

	// assemble app
	welcome.Greeter = stdout.Greeter{}
	welcome.Desk = shop.Desk{}

	// execute use case
	welcome.TheCustomer(customerName)
}
