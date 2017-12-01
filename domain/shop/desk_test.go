package shop_test

import (
	"testing"

	"github.com/aweisser/clean-arch-golang/domain/shop"
	"github.com/stretchr/testify/assert"
)

func TestNewSingleDeskShouldHaveNoCustomers(t *testing.T) {
	desk := shop.SingleDesk{}
	assert.Equal(t, 0, desk.CountCustomers(), "There should be no customers")
}

func TestAddingCustomersToTheDesk(t *testing.T) {
	desk := shop.SingleDesk{}
	desk.RegisterCustomer(shop.Customer{Name: "John"})
	desk.RegisterCustomer(shop.Customer{Name: "Lisa"})
	assert.Equal(t, 2, desk.CountCustomers(), "There should be two customers")
}
func TestCustomerCanNotBeAddedMoreThanOnce(t *testing.T) {
	desk := shop.SingleDesk{}
	desk.RegisterCustomer(shop.Customer{Name: "John"})
	desk.RegisterCustomer(shop.Customer{Name: "John"})
	assert.Equal(t, 1, desk.CountCustomers(), "There should be only one customer")
}
