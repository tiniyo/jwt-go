package api

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/tiniyo/jwt-go/config"
	"github.com/tiniyo/jwt-go/models"
	"net/http"
	"time"
)

func createJwt(c *Context) error {

	response := make(map[string]interface{})

	jwtStandardClaims := jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
	}
	// Set custom claims
	claims := &models.JwtClaims{
		Name:"Test",
		Admin:true,
		StandardClaims:jwtStandardClaims,
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(config.Conf.JWTSecretKey))
	if err != nil {
		return err
	}

	response["token"] = t

	return c.JSON(http.StatusOK, response)

	return nil
}

func verifyJwt(c *Context) error {

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*models.JwtClaims)
	name := claims.Name
	return c.String(http.StatusOK, "Welcome "+name+"!")

	return nil
}
