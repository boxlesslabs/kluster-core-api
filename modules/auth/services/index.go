package services

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/klusters-core/api/config/secrets"
	model "github.com/klusters-core/api/modules/account/models"
	"github.com/klusters-core/api/modules/auth/models"
	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"time"
)

// ReturnUser internal/private methods
func (auth *authService) ReturnUser(authID primitive.ObjectID, ctx echo.Context, userAccount *model.AccountsModel) error {
	token, err := auth.SignToken(userAccount, authID)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, result.ReturnErrorResult(err.Error()))
	}
	return ctx.JSON(http.StatusOK, result.ReturnAuthResult(userAccount, token))
}

func (auth *authService) VerifyTokenValidity(token string) error {
	jwtKey := []byte(secrets.GetSecrets().JWTSecrets)
	claims := &models.JwtCustomClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	return err
}

func (auth *authService) SignToken (account *model.AccountsModel, authID primitive.ObjectID) (string, error) {
	claims := &models.JwtCustomClaims{
		authID,
		account.ID,
		account.Phone,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	_token, tokenErr := token.SignedString([]byte(secrets.GetSecrets().JWTSecrets))
	if tokenErr != nil {
		return "", tokenErr
	}
	return _token, nil
}

func (auth *authService) DecodeClaims(token string) *models.JwtCustomClaims {
	jwtKey := []byte(secrets.GetSecrets().JWTSecrets)
	claims := &models.JwtCustomClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return nil
	}
	return claims
}