package main

import (
	"fmt"

	"github.com/aweisser/clean-arch-golang/domain"
)

func main() {
	user := domain.User{Name: "Armin"}
	fmt.Printf("Hello, %v.\n", user.Name)
}
