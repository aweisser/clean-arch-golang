package welcome_test

import (
	"testing"

	"github.com/aweisser/clean-arch-golang/domain/shop"
	"github.com/aweisser/clean-arch-golang/usecase/welcome"

	"github.com/aweisser/clean-arch-golang"

	"github.com/stretchr/testify/assert"
)

type greeter struct {
	welcomeClbck func(customer shop.Customer, message string)
}

func (g *greeter) Welcome(customer shop.Customer, message string) {
	g.welcomeClbck(customer, message)
}

func TestOneCustomer(t *testing.T) {
	factory.InitCliShop()
	welcome.TheCustomer("John")
	assert.Equal(t, 1, welcome.Counter.CountCustomers(), "There should only be one customers")
}

func TestTwoCustomers(t *testing.T) {
	factory.InitCliShop()
	welcome.TheCustomer("John")
	welcome.TheCustomer("Lisa")
	assert.Equal(t, 2, welcome.Counter.CountCustomers(), "There should be two customers")
}

func TestThereCanOnlyBeOneFirstCustomer(t *testing.T) {
	factory.InitCliShop()
	welcome.Greeter = &greeter{welcomeClbck: func(customer shop.Customer, message string) {
		if customer.Name == "John" {
			assert.Contains(t, message, "You're our first customer.")
		} else {
			assert.NotContains(t, message, "You're our first customer.")
		}
	}}
	welcome.TheCustomer("John")
	welcome.TheCustomer("Lisa")
}

func TestEachCustomerIsUnique(t *testing.T) {
	factory.InitCliShop()
	welcome.TheCustomer("John")
	welcome.TheCustomer("John")
	assert.Equal(t, 1, welcome.Counter.CountCustomers(), "There should be one customers")
}
