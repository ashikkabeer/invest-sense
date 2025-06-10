package server

import (
	"net/http"

	"github.com/ashikkabeer/invest-sense/internal/handler"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RegisterRoutes(e *echo.Echo, h *handler.Handler) {
	// Global middlewares
	e.Use(middleware.Logger())

	// Public routes
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"status": "ok"})
	})

	// Auth group
	// auth := e.Group("/api")
	// auth.Use(middleware.JWTMiddleware())
	// auth.GET("/profile", h.GetProfile)
}
