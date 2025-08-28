package provider

import (
	"fiber-gorm-channel-ecommerce/src/interface/rest"
	"github.com/gofiber/fiber/v2"
)

func RegisterRouteProvider(router *fiber.App) {
	rest.ProductRoutes(router)
}
