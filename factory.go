package factory

import (
	"github.com/aweisser/clean-arch-golang/domain/shop"
	"github.com/aweisser/clean-arch-golang/io/stdout"
	"github.com/aweisser/clean-arch-golang/usecases/welcome"
)

// InitCliShop assembles a shop for CLI tools
func InitCliShop() {
	welcome.Greeter = stdout.Greeter{}
	welcome.Desk = shop.Desk{}
}
