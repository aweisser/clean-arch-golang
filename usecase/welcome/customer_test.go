package welcome_test

import (
	"testing"

	"github.com/aweisser/clean-arch-golang/usecase/welcome"

	"github.com/aweisser/clean-arch-golang"

	"github.com/stretchr/testify/assert"
)

func TestTwoCustomers(t *testing.T) {
	factory.InitCliShop()
	welcome.TheCustomer("John")
	welcome.TheCustomer("Lisa")
	assert.Equal(t, 2, welcome.Counter.CountCustomers(), "There should be two customers")
}

func TestUnique(t *testing.T) {
	factory.InitCliShop()
	welcome.TheCustomer("John")
	welcome.TheCustomer("John")
	assert.Equal(t, 1, welcome.Counter.CountCustomers(), "There should be one customers")
}
