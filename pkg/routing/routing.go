package routing

import (
	"fiber-gorm-channel-ecommerce/interface/provider"
	"github.com/gofiber/fiber/v2"
)

var app *fiber.App

func Init() {
	app = fiber.New()
}

func GetRouter() *fiber.App {
	return app
}

func RegisterRoutes() {
	provider.RegisterRouteProvider(GetRouter())
}
