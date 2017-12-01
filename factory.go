package factory

import (
	"github.com/aweisser/clean-arch-golang/domain/shop"
	"github.com/aweisser/clean-arch-golang/io/stdout"
	"github.com/aweisser/clean-arch-golang/usecase/welcome"
)

// InitCliShop assembles a shop for CLI tools
func InitCliShop() {
	welcome.Greeter = &stdout.Greeter{}
	welcome.Registry = &shop.SingleDesk{}
}
