package boot

import (
	"fiber-gorm-channel-ecommerce/pkg/configCore"
	"fiber-gorm-channel-ecommerce/pkg/routing"
)

func Serve() {
	configCore.ConfigurationSet()

	routing.Init()

	routing.RegisterRoutes()

	routing.RunServer()

}
