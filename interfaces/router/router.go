package router

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var apikey = os.Getenv("API_KEY")

type router struct {
	router *echo.Echo
}

type Router interface {
	GetDevice(path string, f func(c echo.Context) error)
	AddNewDevice(path string, f func(c echo.Context) error)
	Start(address string)
}

func NewRouter() Router {
	return &router{router: echo.New()}
}

func (r router) GetDevice(path string, f func(c echo.Context) error) {
	r.router.GET(path, f)
}

func (r router) AddNewDevice(path string, f func(c echo.Context) error) {
	r.router.POST(path, f)
}

func (r router) Start(address string) {
	r.router.Use(middleware.Logger())
	r.router.Use(middleware.Recover())

	if err := r.router.Start(address); err != http.ErrServerClosed {
		r.router.Logger.Fatal(err)
	}
}
