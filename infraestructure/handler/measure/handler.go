package measure

import (
	"fmt"
	"github.com/CarosDrean/amachay/domain/measure"
	"github.com/CarosDrean/amachay/model"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type Handler struct {
	useCase measure.UseCase
}

func NewHandler(useCase measure.UseCase) Handler {
	return Handler{
		useCase: useCase,
	}
}

func (h Handler) create(c echo.Context) error {
	requestData := model.Measure{}
	if err := c.Bind(&requestData); err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Errorf("c.Bind()"))
	}

	err := h.useCase.Create(requestData)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Errorf("unexpedted error"))
	}

	return c.JSON(http.StatusCreated, fmt.Sprintf("created successful"))
}

func (h Handler) update(c echo.Context) error {
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Errorf("unexpedted error"))
	}

	requestData := model.Measure{}
	if err := c.Bind(&requestData); err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Errorf("c.Bind()"))
	}

	err = h.useCase.Update(ID, requestData)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Errorf("unexpedted error"))
	}

	return c.JSON(http.StatusOK, fmt.Sprintf("updated successful"))
}

func (h Handler) delete(c echo.Context) error {
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Errorf("unexpedted error"))
	}

	err = h.useCase.Delete(ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Errorf("unexpedted error"))
	}

	return c.JSON(http.StatusOK, fmt.Sprintf("deleted successful"))
}

func (h Handler) get(c echo.Context) error {
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Errorf("unexpedted error"))
	}

	data, err := h.useCase.Get(ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Errorf("unexpedted error"))
	}

	return c.JSON(http.StatusOK, data)
}

func (h Handler) getAll(c echo.Context) error {
	data, err := h.useCase.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Errorf("unexpedted error"))
	}

	return c.JSON(http.StatusOK, data)
}
