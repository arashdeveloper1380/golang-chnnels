package boot

import (
	"fiber-gorm-channel-ecommerce/src/infrastructure/database/migrations"
	"fiber-gorm-channel-ecommerce/src/pkg/configCore"
	"fiber-gorm-channel-ecommerce/src/pkg/databaseCore"
)

func Migrate() {
	configCore.ConfigurationSet()

	databaseCore.ConnectMysqlDB()

	migrations.Migrate()
}
