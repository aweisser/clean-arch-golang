package shop

// Greeter greets Customer
type Greeter interface {
	Welcome(customer Customer, message string)
}
