package welcome

import (
	"github.com/aweisser/clean-arch-golang/domain/shop"
)

// To welcome a new customer we need a greeter
type greeter interface {
	Welcome(customer shop.Customer, message string)
}

// To register a new customer we need a customer registry
type registry interface {
	RegisterCustomer(customer shop.Customer)
}

// Greeter to greet the Customer
var Greeter greeter

// Registry to register new customers
var Registry registry

// TheCustomer welcomes the new Customer
func TheCustomer(customerName string) {
	customer := shop.Customer{Name: customerName}
	Registry.RegisterCustomer(customer)
	Greeter.Welcome(customer, "You're now registered. We wish you a pleasent stay in our shop.")
}
