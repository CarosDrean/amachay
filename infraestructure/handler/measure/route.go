package measure

import (
	"github.com/labstack/echo/v4"

	"github.com/CarosDrean/amachay/domain/measure"
	storageMeasure "github.com/CarosDrean/amachay/infraestructure/db/memory/measure"
	"github.com/CarosDrean/amachay/kit/db"
)

const (
	privateRoutePrefix = "api/v1/measure"
)

func NewRoutes(app *echo.Echo, db db.DB) {
	storage := storageMeasure.NewMeasure(db)
	useCase := measure.NewMeasure(storage)
	handler := NewHandler(useCase)

	privateRoutes(app, handler)
}

func privateRoutes(app *echo.Echo, handler Handler) {
	api := app.Group(privateRoutePrefix)

	api.POST("/", handler.create)
	api.PUT("/:id", handler.update)
	api.DELETE("/:id", handler.delete)
	api.GET("/:id", handler.get)
	api.GET("/", handler.getAll)
}
