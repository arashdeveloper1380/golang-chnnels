package rest

import (
	"fiber-gorm-channel-ecommerce/src/domain/model/entity"
	"fiber-gorm-channel-ecommerce/src/pkg/databaseCore"
	"fiber-gorm-channel-ecommerce/src/pkg/qrybldr"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func ProductRoutes(router *fiber.App) {
	router.Get("/", func(ctx *fiber.Ctx) error {

		db := databaseCore.Connection()
		qb := qrybldr.Instance(db)

		var categories []entity.Category

		_ = qb.Table("categories").Get(&categories)

		ch := make(chan []entity.Category)

		go func() {
			ch <- categories
			close(ch)
		}()

		result := <-ch

		return ctx.Status(http.StatusOK).JSON(&fiber.Map{
			"data": result,
		})
	})
}
