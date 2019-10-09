package route

import (
	"simple-rest-api-moryku/app/controller/user"
	"simple-rest-api-moryku/app/config"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Init() *echo.Echo {
	configuration := config.GetConfig()
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.CORS())

	eJwt := e.Group("/api/v1")
	eNoJwt := e.Group("/api/v1")
	eJwt.Use(middleware.JWT([]byte(configuration.ENCRYCPTION_KEY)))
	eJwt.POST("/users/logout", user.Logout)
	eNoJwt.POST("/users/register", user.CreateUser)
	eNoJwt.POST("/users/login", user.Login)

	return e
}
