package services

import (
	err_res "github.com/klusters-core/api/config/error_response"
	authRepo "github.com/klusters-core/api/modules/auth/repo"
	"github.com/klusters-core/api/modules/auth/services"
	"github.com/klusters-core/api/modules/clusters/models"
	"github.com/klusters-core/api/modules/clusters/repo"
	"github.com/klusters-core/api/utils"
	"github.com/labstack/echo"
	"log"
	"net/http"
)

var (
	result utils.Result
	validate utils.ValidateUtil
)

func NewManagerService(service repo.ClusterRepo, auth authRepo.AuthRepo) *managerService {
	return &managerService{service, services.NewAuthService(auth)}
}

type (
	managerService struct {
		repo.ClusterRepo
		Auth services.UserService
	}

	ManagerService interface {
		AddNewCluster(ctx echo.Context) error
		GetSingleCluster(ctx echo.Context) error
		DeleteMyCluster(ctx echo.Context) error
		EditMyCluster(ctx echo.Context) error
		ToggleVisibility(ctx echo.Context) error
	}
)

func (m *managerService) AddNewCluster(ctx echo.Context) error {
	claims := m.Auth.ReturnSignedInUser(ctx)

	var request = new(models.CreateRequest)
	if err := ctx.Bind(request); err != nil {
		return ctx.JSON(http.StatusBadRequest, result.ReturnErrorResult(err.Error()))
	}

	// validate request
	if err := validate.Validate(request); err != nil {
		return ctx.JSON(http.StatusBadRequest, result.ReturnValidateError(err))
	}

	res, err := m.CreateCluster(request, claims.UserID)
	if err != nil {
		log.Println(err)
		return ctx.JSON(http.StatusInternalServerError, result.ReturnErrorResult(err.Error()))
	}

	return ctx.JSON(http.StatusOK, result.ReturnBasicResult(res))
}

func (m *managerService) GetSingleCluster(ctx echo.Context) error {
	clusterID, err := validate.ValidateParam(ctx, "clusterID", result)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, result.ReturnErrorResult(err.Error()))
	}

	cluster, err := m.GetByID(clusterID)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, result.ReturnErrorResult(err_res.ErrorGetting{Resource: "cluster"}.Error()))
	}

	return ctx.JSON(http.StatusOK, result.ReturnBasicResult(cluster))
}

func (m *managerService) DeleteMyCluster(ctx echo.Context) error {
	clusterID, err := validate.ValidateParam(ctx, "clusterID", result)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, result.ReturnErrorResult(err.Error()))
	}

	if err := m.DeleteByID(clusterID); err != nil {
		return ctx.JSON(http.StatusBadRequest, result.ReturnErrorResult(err_res.ErrorDeleting{Resource: "cluster"}.Error()))
	}

	return ctx.JSON(http.StatusOK, result.ReturnSuccessMessage("cluster has been deleted successfully"))
}

func (m *managerService) EditMyCluster(ctx echo.Context) error {
	clusterID, err := validate.ValidateParam(ctx, "clusterID", result)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, result.ReturnErrorResult(err.Error()))
	}

	var request = new(models.CreateRequest)
	if err := ctx.Bind(request); err != nil {
		return ctx.JSON(http.StatusBadRequest, result.ReturnErrorResult(err.Error()))
	}

	// validate request
	if err := validate.Validate(request); err != nil {
		return ctx.JSON(http.StatusBadRequest, result.ReturnValidateError(err))
	}

	res, err := m.UpdateDetails(request, clusterID)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, result.ReturnErrorResult(err_res.ErrorUpdating{Resource: "cluster"}.Error()))
	}

	return ctx.JSON(http.StatusOK, result.ReturnBasicResult(res))
}

func (m *managerService) ToggleVisibility(ctx echo.Context) error {
	clusterID, err := validate.ValidateParam(ctx, "clusterID", result)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, result.ReturnErrorResult(err.Error()))
	}

	var request = new(models.VisibilityRequest)
	if err := ctx.Bind(request); err != nil {
		return ctx.JSON(http.StatusBadRequest, result.ReturnErrorResult(err.Error()))
	}

	// validate request
	if err := validate.Validate(request); err != nil {
		return ctx.JSON(http.StatusBadRequest, result.ReturnValidateError(err))
	}

	var setVisibility string
	switch request.Visibility {
	case false:
		setVisibility = "public"
	case true:
		setVisibility = "private"
	}

	res, err := m.ChangeVisibility(clusterID, setVisibility)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, result.ReturnErrorResult(err_res.ErrorUpdating{Resource: "cluster"}.Error()))
	}

	return ctx.JSON(http.StatusOK, result.ReturnBasicResult(res))
}