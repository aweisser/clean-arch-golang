package welcome

import (
	"github.com/aweisser/clean-arch-golang/domain/shop"
)

type greeter interface {
	Welcome(customer shop.Customer, message string)
}

// Greeter to greet the Customer
var Greeter greeter

// Desk interacts with the customers
var Desk shop.Desk

// TheCustomer welcomes the new Customer
func TheCustomer(customerName string) {
	customer := shop.Customer{Name: customerName}
	Desk.RegisterCustomer(customer)
	Greeter.Welcome(customer, "You're now registered. We wish you a pleasent stay in our shop.")
}
