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
	app.Group(privateRoutePrefix)

	app.POST("/", handler.create)
	app.PUT("/:id", handler.update)
	app.DELETE("/:id", handler.delete)
	app.GET("/:id", handler.get)
	app.GET("/", handler.getAll)
}
