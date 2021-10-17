package middlewares

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/klusters-core/api/config/db"
	"github.com/klusters-core/api/config/secrets"
	model "github.com/klusters-core/api/modules/account/models"
	"github.com/klusters-core/api/modules/account/repo"
	"github.com/klusters-core/api/modules/auth/models"
	"github.com/klusters-core/api/utils"
	"github.com/labstack/echo"
	"log"
	"net/http"
	"strings"
)

var (
	result utils.Result
)

type CustomContext struct {
	AccountClaims *models.JwtCustomClaims
	Account *model.AccountsModel
	echo.Context
}

// applies to all logged in users
func IsValidUser(con db.StartMongoClient) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			getToken := ctx.Request().Header.Get("Authorization")
			if getToken == "" {
				return echo.NewHTTPError(http.StatusUnauthorized, result.ReturnErrorResult("Could not find authorization token"))
			}

			accountClaims, err := DecodeToken(getToken)
			if err != nil {
				log.Println(err)
				return echo.NewHTTPError(http.StatusUnauthorized, result.ReturnErrorResult("Invalid token, token has expired"))
			}

			accountRepo := repo.NewAccountRepo(con)
			account, err := accountRepo.GetAccount(accountClaims.UserID)
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, result.ReturnErrorResult("account does not exist, kindly recheck your credentials"))
			}

			if account.Status == model.StatusDeactivated {
				log.Println("here")
				return echo.NewHTTPError(http.StatusUnauthorized, result.ReturnErrorResult("account has been deactivated, kindly contact support to activate your account"))
			}

			newContext := &CustomContext{
				AccountClaims: accountClaims,
				Account:account,
				Context: ctx,
			}

			return next(newContext)
		}
	}
}

func DecodeToken(getToken string) (*models.JwtCustomClaims, error) {
	token := strings.Split(getToken, " ")
	jwtKey := []byte(secrets.GetSecrets().JWTSecrets)
	claims := &models.JwtCustomClaims{}
	_, err := jwt.ParseWithClaims(token[1], claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	return claims, err
}