package services

import (
	err_res "github.com/klusters-core/api/config/error_response"
	accountRepo "github.com/klusters-core/api/modules/account/repo"
	"github.com/klusters-core/api/modules/auth/models"
	"github.com/klusters-core/api/modules/auth/repo"
	"github.com/klusters-core/api/utils"
	"github.com/labstack/echo"
	"log"
	"net/http"
	"strings"
)

var (
	result utils.Result
	validate utils.ValidateUtil
)

func NewAuthService(service repo.AuthRepo) *authService {
	return &authService{service}
}

type (
	authService struct {
		repo.AuthRepo
	}

	UserService interface {
		Authenticate(ctx echo.Context) error
		ReturnSignedInUser(ctx echo.Context) *models.JwtCustomClaims
		RefreshToken(ctx echo.Context) error
		ForgotPassword(ctx echo.Context) error
		ChangePassword(ctx echo.Context) error
	}
)

func (auth *authService) Authenticate(ctx echo.Context) error {
	var request = new(models.AuthRequest)
	if err := ctx.Bind(request); err != nil {
		return err
	}

	if err := validate.Validate(request); err != nil {
		return ctx.JSON(http.StatusBadRequest, result.ReturnValidateError(err))
	}

	user, err := auth.GetByCredentials(request)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, result.ReturnErrorResult(err_res.InvalidLoginCredentials{}.Error()))
	}

	accRepo := accountRepo.NewAccountRepo(auth.AuthRepo.ReturnClient())
	account, err := accRepo.GetByPhone(user.Phone)
	if err != nil {
		log.Println(err)
	}

	return auth.ReturnUser(user.ID, ctx, account)
}

func (auth *authService) ReturnSignedInUser(ctx echo.Context) *models.JwtCustomClaims {
	authHeader := ctx.Request().Header.Get("Authorization")
	claims := strings.Split(authHeader, " ")

	return auth.DecodeClaims(claims[1])
}

func (auth *authService) RefreshToken(ctx echo.Context) error {
	claims := auth.ReturnSignedInUser(ctx)
	user, err := auth.GetByPhone(claims.Phone)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, result.ReturnErrorResult(err_res.ErrorGetting{Resource:"user auth"}.Error()))
	}

	accRepo := accountRepo.NewAccountRepo(auth.AuthRepo.ReturnClient())
	account, err := accRepo.GetByPhone(user.Phone)
	if err != nil {
		log.Println(err)
	}

	return auth.ReturnUser(user.ID, ctx, account)
}

func (auth *authService) ForgotPassword(ctx echo.Context) error {
	var request = new(models.ForgotPasswordRequest)
	if err := ctx.Bind(request); err != nil {
		return err
	}

	if err := validate.Validate(request); err != nil {
		return ctx.JSON(http.StatusBadRequest, result.ReturnValidateError(err))
	}

	user, err := auth.GetByPhone(request.Phone)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err_res.ErrorGetting{Resource:"user account"}.Error())
	}

	// handle email implementation here
	// ...

	accRepo := accountRepo.NewAccountRepo(auth.AuthRepo.ReturnClient())
	account, err := accRepo.GetByPhone(user.Phone)
	if err != nil {
		log.Println(err)
	}

	token, err := auth.SignToken(account, user.ID)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err_res.ErrorProcessing{}.Error())
	}

	return ctx.JSON(http.StatusOK, result.ReturnSuccessResult("Email successfully sent to your recovery email.", token))
}

func (auth *authService) ChangePassword(ctx echo.Context) error {
	claims := auth.ReturnSignedInUser(ctx)
	var request = new(models.ChangePasswordRequest)
	if err := ctx.Bind(request); err != nil {
		return err
	}

	if err := validate.Validate(request); err != nil {
		return ctx.JSON(http.StatusBadRequest, result.ReturnValidateError(err))
	}

	// check if old password is valid against user
	authObject, err := auth.ComparePasswords(claims.UserID, request.OldPassword)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, result.ReturnErrorResult("Oops! Your old password is invalid"))
	}

	// check if old password is same as new password
	if authObject.Password == request.NewPassword {
		return ctx.JSON(http.StatusBadRequest, result.ReturnErrorResult("Oops! Your old password and new password is the same"))
	}

	// update user password
	res, err := auth.UpdatePassword(claims.UserID, request.NewPassword)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, result.ReturnErrorResult(err_res.ErrorUpdating{Resource: "user account"}.Error()))
	}

	return ctx.JSON(http.StatusOK, result.ReturnSuccessResult(res, "Password updated successfully"))
}