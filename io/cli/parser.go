package cli

import "flag"

// ParseNameFlag parses the --name flag
func ParseNameFlag() string {
	var name string
	flag.StringVar(&name, "name", "John Doe", "enter your name")
	flag.Parse()
	return name
}
