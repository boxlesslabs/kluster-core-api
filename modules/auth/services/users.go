package services

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	err_res "github.com/klusters-core/api/config/error_response"
	"github.com/klusters-core/api/config/secrets"
	"github.com/klusters-core/api/middlewares"
	model "github.com/klusters-core/api/modules/account/models"
	accountRepo "github.com/klusters-core/api/modules/account/repo"
	"github.com/klusters-core/api/modules/auth/models"
	"github.com/klusters-core/api/modules/auth/repo"
	"github.com/klusters-core/api/utils"
	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
	"time"
)

type (
	authService struct {
		IAuthRepo repo.AuthRepo
		*utils.Result
		*utils.ValidateUtil
		*utils.GeneralUtil
		Model *models.AuthModel
	}

	UserService interface {
		Authenticate(ctx echo.Context) error
		RefreshToken(ctx echo.Context) error
		ForgotPassword(ctx echo.Context) error
		ChangePassword(ctx echo.Context) error
	}
)

func NewAuthService(repo repo.AuthRepo) *authService {
	return &authService{IAuthRepo:repo}
}


// PUBLIC
func (auth *authService) Authenticate(ctx echo.Context) (err error) {
	request := &models.AuthModel{}
	err = json.NewDecoder(ctx.Request().Body).Decode(request)
	if err = request.ValidateAuth(); err != nil {
		return ctx.JSON(http.StatusBadRequest, auth.ReturnValidateError(err))
	}

	if request, err = auth.IAuthRepo.GetByCredentials(request); err != nil {
		return ctx.JSON(http.StatusBadRequest, auth.ReturnErrorResult(err_res.InvalidLoginCredentials{}.Error()))
	}

	accRepo := accountRepo.NewAccountRepo(auth.IAuthRepo.ReturnClient())
	account, err := accRepo.GetByPhone(request.Phone)
	if err != nil {
		log.Println(err)
		return ctx.JSON(http.StatusBadRequest, auth.ReturnErrorResult(err_res.NotFound{Resource:"account"}.Error()))
	}

	return auth.SignToken(ctx, account, request.ID)
}

func (auth *authService) RefreshToken(ctx echo.Context) (err error) {
	claims := ctx.(*middlewares.CustomContext)
	auth.Model, err = auth.IAuthRepo.GetByPhone(&claims.AccountClaims.Phone)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, auth.ReturnErrorResult(err_res.ErrorGetting{Resource:"user auth"}.Error()))
	}

	accRepo := accountRepo.NewAccountRepo(auth.IAuthRepo.ReturnClient())
	account, err := accRepo.GetByPhone(auth.Model.Phone)
	if err != nil {
		log.Println(err)
	}

	return auth.SignToken(ctx, account, claims.AccountClaims.AuthID)
}

func (auth *authService) ForgotPassword(ctx echo.Context) (err error) {
	request := &models.AuthModel{}
	err = json.NewDecoder(ctx.Request().Body).Decode(request)
	if err = request.ValidateForgotPassword(); err != nil {
		log.Println(err)
		return ctx.JSON(http.StatusBadRequest, auth.ReturnValidateError(err))
	}

	if request, err = auth.IAuthRepo.GetByPhone(&request.Phone); err != nil {
		return ctx.JSON(http.StatusBadRequest, err_res.ErrorGetting{Resource:"user account"}.Error())
	}

	// todo: handle email implementation here
	// ...

	accRepo := accountRepo.NewAccountRepo(auth.IAuthRepo.ReturnClient())
	account, err := accRepo.GetByPhone(request.Phone)
	if err != nil {
		log.Println(err)
		return ctx.JSON(http.StatusBadRequest, err_res.ErrorGetting{Resource:"user account"}.Error())
	}

	return auth.SignToken(ctx, account, request.ID)
}

func (auth *authService) ChangePassword(ctx echo.Context) error {
	claims := ctx.(*middlewares.CustomContext)
	var request = &models.ChangePasswordRequest{}
	err := json.NewDecoder(ctx.Request().Body).Decode(request)
	if err := request.ValidateChangePassword(); err != nil {
		return ctx.JSON(http.StatusBadRequest, auth.ReturnValidateError(err))
	}

	// check if old password is valid against user
	 if auth.Model, err = auth.IAuthRepo.ComparePasswords(&claims.AccountClaims.AuthID, request.OldPassword, request.NewPassword); err != nil {
		return ctx.JSON(http.StatusBadRequest, auth.ReturnErrorResult(err.Error()))
	 }

	if auth.Model, err = auth.IAuthRepo.UpdatePassword(&claims.AccountClaims.AuthID, request.NewPassword); err != nil {
		log.Println(err)
		return ctx.JSON(http.StatusBadRequest, auth.ReturnErrorResult(err_res.ErrorUpdating{Resource: "user"}.Error()))
	}

	return ctx.JSON(http.StatusOK, auth.ReturnSuccessResult(auth.Model, "password updated successfully"))
}


// PRIVATE
func (auth *authService) SignToken (ctx echo.Context, account *model.AccountsModel, authID primitive.ObjectID) error {
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
		return ctx.JSON(http.StatusBadRequest, auth.ReturnErrorResult(tokenErr.Error()))
	}
	return ctx.JSON(http.StatusOK, auth.ReturnAuthResult(account, _token))
}