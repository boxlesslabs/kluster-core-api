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
	"encoding/json"
	"github.com/klusters-core/api/middlewares"
	model "github.com/klusters-core/api/modules/account/models"
	"github.com/klusters-core/api/modules/account/repo"
	"github.com/klusters-core/api/modules/auth/models"
	"github.com/klusters-core/api/utils"
	"github.com/labstack/echo"
	"net/http"
)

type (
	userService struct {
		IAccountRepo repo.AccountRepo
		Model        *model.AccountsModel
		*utils.Result
	}

	UserInterface interface {
		CreateUser(ctx echo.Context) error
		GetUser(ctx echo.Context) error
		UpdateProfile(ctx echo.Context) error
	}
)


func NewAccountService(service repo.AccountRepo) *userService {
	return &userService{IAccountRepo:service}
}

func (account *userService) CreateUser(ctx echo.Context) (err error) {
	var request *models.AuthModel
	err = json.NewDecoder(ctx.Request().Body).Decode(&request)
	if err = request.ValidateAuth(); err != nil {
		return ctx.JSON(http.StatusBadRequest, account.ReturnValidateError(err))
	}

	if account.Model, err = account.IAccountRepo.CreateAccount(request); err != nil {
		return ctx.JSON(http.StatusInternalServerError, account.ReturnErrorResult(err.Error()))
	}

	// todo: send email or sms confirmation after successful registration

	return ctx.JSON(http.StatusCreated, account.ReturnBasicResult(account.Model))
}

func (account *userService) GetUser(ctx echo.Context) error {
	userAccount, _ := ctx.(*middlewares.CustomContext)
	return ctx.JSON(http.StatusOK, account.ReturnBasicResult(userAccount.Account))
}

func (account *userService) UpdateProfile(ctx echo.Context) (err error) {
	userAccount, _ := ctx.(*middlewares.CustomContext)
	var request *model.AccountsModel
	err = json.NewDecoder(ctx.Request().Body).Decode(&request)
	if err = request.ValidateProfileReq(); err != nil {
		return ctx.JSON(http.StatusBadRequest, account.ReturnValidateError(err))
	}

	request.ID = userAccount.AccountClaims.UserID
	if request, err = account.IAccountRepo.UpdateAccountByModel(request); err != nil {
		return ctx.JSON(http.StatusInternalServerError, account.ReturnErrorResult(err.Error()))
	}

	return ctx.JSON(http.StatusOK, account.ReturnSuccessResult(request, "account updated successfully"))
}