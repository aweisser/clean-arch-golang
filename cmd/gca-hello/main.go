package main

import (
	"flag"

	"github.com/aweisser/clean-arch-golang/domain/gca"
	"github.com/aweisser/clean-arch-golang/usecases/user"
)

func main() {
	var userName string
	flag.StringVar(&userName, "user", "John Doe", "enter a user")
	flag.Parse()
	u := gca.User{Name: userName}
	user.SayHello(u)
}
