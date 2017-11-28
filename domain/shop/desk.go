package shop

// Desk interacts with the customers
type Desk struct {
	customers []Customer
}

// RegisterCustomer adds the customer to the list of current customers
func (desk Desk) RegisterCustomer(customer Customer) {
	desk.customers = append(desk.customers, customer)
}
