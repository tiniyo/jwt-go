package api

import (
	"github.com/labstack/echo/v4"
	mw "github.com/labstack/echo/v4/middleware"
	"github.com/tiniyo/jwt-go/config"
	"github.com/tiniyo/jwt-go/models"
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
		Claims:     &models.JwtClaims{},
		SigningKey: []byte(config.Conf.Jwt.JWTSecretKey),
	}

	e.Use(mw.JWTWithConfig(config))

	e.Use(mw.Logger())
	e.Use(mw.Recover())

	e.POST("/jwt", handler(createJwt))
	e.POST("/verifyJwt", handler(verifyJwt))

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
