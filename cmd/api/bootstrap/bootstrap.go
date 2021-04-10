package bootstrap

import (
	"fmt"
	"github.com/CarosDrean/amachay/infraestructure/handler/router"
)

func Run() error {
	config := newConfiguration("./configuration.json")
	api := newEcho()
	db := newDB(config)

	router.InitRoutes(api, db)

	port := fmt.Sprintf(":%d", config.Port)
	return api.Start(port)
}
