package migrations

import (
	"fiber-gorm-channel-ecommerce/src/domain/model/entity"
	"fiber-gorm-channel-ecommerce/src/pkg/databaseCore"
	"fmt"
	"log"
)

func Migrate() {
	db := databaseCore.Connection()

	err := db.AutoMigrate(
		&entity.User{},
		&entity.Category{},
		&entity.Product{},
		&entity.OrderItem{},
		&entity.Order{},
		&entity.Payment{},
	)

	if err != nil {
		log.Fatal("Cant migrate", err)
		return
	}

	fmt.Println("Migration done ..")
}
