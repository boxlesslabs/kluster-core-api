//=============================================================================
// developer: boxlesslabsng@gmail.com
//=============================================================================

/**
 * Define user use case
 **
 * @struct userService
**/

package services

import (
	"github.com/klusters-core/api/middlewares"
	model "github.com/klusters-core/api/modules/account/models"
	"github.com/klusters-core/api/modules/account/repo"
	"github.com/klusters-core/api/modules/auth/models"
	"github.com/klusters-core/api/utils"
	"github.com/labstack/echo"
	"log"
	"net/http"
)

var (
	result utils.Result
	validate utils.ValidateUtil
)

func NewUserService(service repo.AccountRepo) *userService {
	return &userService{service}
}

type (
	userService struct {
		repo.AccountRepo
	}

	UserInterface interface {
		CreateUser(ctx echo.Context) error
		GetUser(ctx echo.Context) error
	}
)

func (account *userService) CreateUser(ctx echo.Context) error {
	var request = new(model.AccountRequest)
	if err := ctx.Bind(request); err != nil {
		return ctx.JSON(http.StatusBadRequest, result.ReturnErrorResult(err.Error()))
	}

	// validate request
	if err := validate.Validate(request); err != nil {
		return ctx.JSON(http.StatusBadRequest, result.ReturnValidateError(err))
	}

	auth := models.AuthRequest{
		Phone:   request.Phone,
		Password: request.Password,
	}

	res, err := account.CreateAccount(request, &auth)
	if err != nil {
		log.Println(err)
		return ctx.JSON(http.StatusInternalServerError, result.ReturnErrorResult(err.Error()))
	}

	return ctx.JSON(http.StatusOK, result.ReturnBasicResult(res))
}

func (account *userService) GetUser(ctx echo.Context) error {
	userAccount, _ := ctx.(*middlewares.AccountContext)
	log.Println(*userAccount, "supposed account")

	return ctx.JSON(http.StatusOK, result.ReturnBasicResult(userAccount.Account))
}