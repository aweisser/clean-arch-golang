package stdout

import (
	"fmt"

	"github.com/aweisser/clean-arch-golang/domain/shop"
)

// Greeter prints welcome messages to stdout
type Greeter struct {
}

// Welcome the shop.Customer via stdout
func (greeter Greeter) Welcome(customer shop.Customer, message string) {
	fmt.Printf("Welcome %v. %v\n", customer.Name, message)
}
