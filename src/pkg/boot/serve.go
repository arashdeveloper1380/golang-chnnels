package boot

import (
	"fiber-gorm-channel-ecommerce/src/pkg/configCore"
	"fiber-gorm-channel-ecommerce/src/pkg/databaseCore"
	"fiber-gorm-channel-ecommerce/src/pkg/routing"
)

func Serve() {
	configCore.ConfigurationSet()

	databaseCore.ConnectMysqlDB()

	routing.Init()

	routing.RegisterRoutes()

	routing.RunServer()

}
