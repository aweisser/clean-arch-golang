package factory

import (
	"github.com/aweisser/clean-arch-golang/domain/shop"
	"github.com/aweisser/clean-arch-golang/io/stdout"
	"github.com/aweisser/clean-arch-golang/usecase/welcome"
)

// InitCliShop assembles a shop for CLI tools
func InitCliShop() {
	greeter := &stdout.Greeter{}
	desk := &shop.SingleDesk{}
	welcome.Greeter = greeter
	welcome.Registry = desk
	welcome.Counter = desk
}
