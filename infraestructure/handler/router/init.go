package router

import (
	"github.com/CarosDrean/amachay/infraestructure/handler/measure"
	"github.com/labstack/echo/v4"

	"github.com/CarosDrean/amachay/kit/db"
)

func InitRoutes(app *echo.Echo, db db.DB) {
	measure.NewRoutes(app, db)
}
