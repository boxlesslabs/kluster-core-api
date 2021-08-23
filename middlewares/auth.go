package middlewares

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/klusters-core/api/config/secrets"
	"github.com/klusters-core/api/modules/auth/models"
	"github.com/klusters-core/api/utils"
	"github.com/labstack/echo"
	"log"
	"net/http"
	"strings"
)

var result utils.Result

// applies to all logged in users
func IsLoggedIn(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		getToken := ctx.Request().Header.Get("Authorization")
		if getToken == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, result.ReturnErrorResult("Could not find authorization token"))
		}

		_, err := DecodeToken(getToken)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, result.ReturnErrorResult("Invalid token, token has expired"))
		}

		return next(ctx)
	}
}

func IsOwner(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		getToken := ctx.Request().Header.Get("Authorization")
		if getToken == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, result.ReturnErrorResult("Could not find authorization token"))
		}

		claims, err := DecodeToken(getToken)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, result.ReturnErrorResult("Invalid token, token has expired"))
		}

		log.Println(claims)

		return next(ctx)
	}
}

func DecodeToken(getToken string) (*models.JwtCustomClaims, error) {
	token := strings.Split(getToken, " ")
	jwtKey := []byte(secrets.GetSecrets().JWTSecrets)
	claims := &models.JwtCustomClaims{}
	_, err := jwt.ParseWithClaims(token[1], claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	return nil, err
}