package welcome

import (
	"github.com/aweisser/clean-arch-golang/domain/shop"
)

type greeter interface {
	Welcome(customer shop.Customer, message string)
}

type registry interface {
	RegisterCustomer(customer shop.Customer)
}

type counter interface {
	CountCustomers() int
}

// Greeter to greet the Customer
var Greeter greeter

// Registry to register new customers
var Registry registry

// Counter counts the registered customers
var Counter counter

// TheCustomer welcomes the new Customer
func TheCustomer(customerName string) {
	customer := shop.Customer{Name: customerName}
	Registry.RegisterCustomer(customer)
	if Counter.CountCustomers() == 1 {
		Greeter.Welcome(customer, "Welcome. You're our first customer. We wish you a pleasent stay in our shop.")
	} else {
		Greeter.Welcome(customer, "We wish you a pleasent stay in our shop.")
	}
}
