package server

import (
	"github.com/ashikkabeer/invest-sense/internal/handler"
	"github.com/labstack/echo/v4"
)

func NewServer() *echo.Echo {
	e := echo.New()
	RegisterRoutes(e, handler.NewHandler())
	return e
}
