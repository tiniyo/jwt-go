package api

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/tiniyo/jwt-go/config"
	"net/http"
	"time"
)

func createJwt(c *Context) error {

	var claimDetails map[string]interface{}

	err := c.Bind(&claimDetails)

	if err!= nil{
		return err
	}

	response := make(map[string]interface{})

	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	for key, value := range claimDetails {
		claims[key] = value
	}
	claims["exp"] = time.Now().Add(time.Hour * config.Conf.Jwt.ExpireMinutes).Unix()

	// Generate encoded token and send it as response.
	tok, err := token.SignedString([]byte(config.Conf.Jwt.JWTSecretKey))
	if err != nil {
		return err
	}

	response["token"] = tok

	return c.JSON(http.StatusOK, response)
}

func verifyJwt(c *Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	return c.JSON(http.StatusOK, claims)
}
