package flagcli

import "flag"

// Definition of the cli tool
type Definition struct {
	Callback func(name string)
}

// Execute parses the --name flag and executes the defined callback function
func Execute(definition Definition) {
	var name string
	flag.StringVar(&name, "name", "John Doe", "enter your name")
	flag.Parse()
	definition.Callback(name)
}
