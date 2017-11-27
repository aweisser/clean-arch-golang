package main

import (
	"fmt"

	"github.com/aweisser/clean-arch-golang/domain/gca"
)

func main() {
	user := gca.User{Name: "Armin"}
	fmt.Printf("Hello, %v.\n", user.Name)
}
