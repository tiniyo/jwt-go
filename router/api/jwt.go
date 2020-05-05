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
		ExpiresAt: time.Now().Add(time.Hour * config.Conf.Jwt.ExpireMinutes).Unix(),
	}

	jwtInfo := models.JwtInfo{
		Name:"Test",
		Admin:true,
	}
	// Set custom claims
	claims := &models.JwtClaims{
		JwtInfo: jwtInfo,
		StandardClaims:jwtStandardClaims,
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

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
	claims := user.Claims.(*models.JwtClaims)
	jwtInfo := claims.JwtInfo
	return c.JSON(http.StatusOK, jwtInfo)
}
