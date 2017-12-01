package shop

// SingleDesk provides interactions with the customers
type SingleDesk struct {
	customers []Customer
}

// RegisterCustomer adds the customer to the list of current customers
func (desk *SingleDesk) RegisterCustomer(customer Customer) {
	desk.customers = append(desk.customers, customer)
}
