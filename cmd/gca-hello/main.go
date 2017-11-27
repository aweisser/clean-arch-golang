package main

import (
	"github.com/aweisser/clean-arch-golang/domain/gca"
	"github.com/aweisser/clean-arch-golang/usecases/user"
)

func main() {
	u := gca.User{Name: "Armin"}
	user.SayHello(u)
}
