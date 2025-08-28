package routing

import (
	"fiber-gorm-channel-ecommerce/src/pkg/configCore"
	"fmt"
	"log"
)

func RunServer() {
	r := GetRouter()

	configs := configCore.ConfigurationGet()

	err := r.Listen(fmt.Sprintf("%s:%s", configs.Server.Host, configs.Server.Port))

	if err != nil {
		log.Fatal("Error in routing")
		return
	}
}
