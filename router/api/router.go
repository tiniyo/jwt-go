package api

import (
	"github.com/labstack/echo/v4"
	mw "github.com/labstack/echo/v4/middleware"
	"github.com/tiniyo/jwt-go/config"
)

//-----
// API Routers
//-----
func Routers() *echo.Echo {
	e := echo.New()

	e.Use(NewContext())

	if config.Conf.ReleaseMode {
		e.Debug = false
	}
	e.Logger.SetPrefix("api")
	e.Logger.SetLevel(config.GetLogLvl())

	e.Use(mw.GzipWithConfig(mw.GzipConfig{
		Level: 5,
	}))

	// Configure middleware with the custom claims type
	config := mw.JWTConfig{
		SigningKey: []byte(config.Conf.Jwt.JWTSecretKey),
	}

	// Restricted group
	r := e.Group("/jwt")
	r.Use(mw.JWTWithConfig(config))
	r.GET("", handler(verifyJwt))

	e.Use(mw.Logger())
	e.Use(mw.Recover())

	e.POST("/jwt", handler(createJwt))

	return e
}

type (
	HandlerFunc func(*Context) error
)

func handler(h HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.(*Context)
		return h(ctx)
	}
}
