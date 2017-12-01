package shop

// SingleDesk provides interactions with the customers
type SingleDesk struct {
	customers map[Customer]bool
}

// RegisterCustomer adds the customer to the list of current customers
func (desk *SingleDesk) RegisterCustomer(customer Customer) {
	if desk.customers == nil {
		desk.customers = make(map[Customer]bool)
	}
	if desk.customers[customer] == false {
		desk.customers[customer] = true
	}
}

// CountCustomers counts all registered customers
func (desk *SingleDesk) CountCustomers() int {
	return len(desk.customers)
}
