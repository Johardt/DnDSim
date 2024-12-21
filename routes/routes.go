package routes

import (
	"DnDSim/handlers"
	"DnDSim/views"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RegisterRoutes(e *echo.Echo) {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return handlers.RenderTempl(c, http.StatusOK, views.IndexPage())
	})

	e.GET("/index", func(c echo.Context) error {
		return handlers.RenderTempl(c, http.StatusOK, views.IndexPage())
	})

	e.GET("/login", func(c echo.Context) error {
		return handlers.RenderTempl(c, http.StatusOK, views.LoginForm())
	})

	e.GET("/register", func(c echo.Context) error {
		return handlers.RenderTempl(c, http.StatusOK, views.RegisterPage())
	})

	handlers.RegisterUserRoutes(e.Group("/users"))
	handlers.RegisterSessionRoutes(e.Group("/sessions"))

	// http.Handle("/play", middleware.Auth(templ.Handler(views.GameSelector())))
}