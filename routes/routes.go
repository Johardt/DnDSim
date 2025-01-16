package routes

import (
	"net/http"

	"DnDSim/handlers"
	"DnDSim/views"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RegisterRoutes(e *echo.Echo) {
	e.Use(middleware.Recover())

	e.Static("/static", "./static")

	e.GET("/", func(c echo.Context) error {
		return handlers.RenderTempl(c, http.StatusOK, views.IndexPage())
	})

	e.GET("/login", func(c echo.Context) error {
		return handlers.RenderTempl(c, http.StatusOK, views.LoginPage())
	})

	e.GET("/register", func(c echo.Context) error {
		return handlers.RenderTempl(c, http.StatusOK, views.RegisterPage())
	})

	handlers.RegisterUserRoutes(e.Group("/users"))
	handlers.RegisterSessionRoutes(e.Group("/sessions"))

	// http.Handle("/play", middleware.Auth(templ.Handler(views.GameSelector())))

	// Middleware to redirect HTTP requests to HTTPS
	e.Pre(middleware.HTTPSRedirect())
}
