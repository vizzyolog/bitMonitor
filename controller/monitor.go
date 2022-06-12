package controller

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"

	"app/domain"
	"app/usecases/monitorUseCase"
	"context"
)

type MonitorController interface {
	FindAll(c echo.Context) error
	FindById(c echo.Context) error
	AddMonitor(c echo.Context) error
}

type Controller struct {
	monitorUseCase monitorUseCase.MonitorService
	ctx            context.Context
}

func NewMonitorController(ctx context.Context, u monitorUseCase.MonitorService) MonitorController {
	return Controller{monitorUseCase: u, ctx: ctx}
}

func (c Controller) FindAll(e echo.Context) error {
	monitors, err := c.monitorUseCase.FindAll(c.ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	if len(monitors) > 0 {
		return e.JSON(http.StatusOK, monitors)
	}

	return e.JSON(http.StatusOK, struct {
		Message string `json:"message"`
	}{
		Message: "No monitors found",
	})
}

func (c Controller) FindById(e echo.Context) error {
	monitor, err := c.monitorUseCase.FindById(c.ctx, e.Param("id"))
	if err != nil {
		if strings.Contains(err.Error(), "mongo: no documents in result") {
			return echo.NewHTTPError(http.StatusNotFound, "Monitor not found")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return e.JSON(http.StatusOK, monitor)
}

func (c Controller) AddMonitor(e echo.Context) error {
	monitor := new(domain.Monitor)
	if err := e.Bind(monitor); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := e.Validate(monitor); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	id, err := c.monitorUseCase.Add(c.ctx, *monitor)
	if err != nil {
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}
	return e.JSON(http.StatusOK, struct {
		Message string `json:"message"`
		Id      string `json:"id"`
	}{
		Message: "New device added successfully",
		Id:      id,
	})
}
