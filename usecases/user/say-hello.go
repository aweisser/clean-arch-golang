package user

import (
	"fmt"

	"github.com/aweisser/clean-arch-golang/domain/gca"
)

// SayHello to user
func SayHello(user gca.User) {
	fmt.Printf("Hello, %v.\n", user.Name)
}
