package service

import (
	"context"
	"errors"
	"os"

	"net/http"

	"github.com/GabrielMoody/mikronet-dashboard-service/internal/dto"
	"github.com/GabrielMoody/mikronet-dashboard-service/internal/helper"
	"github.com/GabrielMoody/mikronet-dashboard-service/internal/models"
	"github.com/GabrielMoody/mikronet-dashboard-service/internal/repository"
)

type DashboardService interface {
	GetAllDrivers(c context.Context, q dto.GetDriverQuery) (res []models.Drivers, err *helper.ErrorStruct)
	GetAllPassengers(c context.Context) (res []models.Passengers, err *helper.ErrorStruct)
	GetDriverById(c context.Context, id string) (res models.Drivers, err *helper.ErrorStruct)
	GetPassengerById(c context.Context, id string) (res models.Passengers, err *helper.ErrorStruct)
	GetAllReviews(c context.Context) (res []models.Reviews, err *helper.ErrorStruct)
	GetAllBlockAccount(c context.Context) (res []models.BlockDriver, err *helper.ErrorStruct)
	GetReviewById(c context.Context, id string) (res models.Reviews, err *helper.ErrorStruct)
	GetAllHistories(c context.Context) (res []models.Histories, err *helper.ErrorStruct)
	EditAmountRoute(c context.Context, data dto.EditAmount, id string) (res models.Route, err *helper.ErrorStruct)
	BlockAccount(c context.Context, accountId string) (res models.BlockedAccount, err *helper.ErrorStruct)
	UnblockAccount(c context.Context, accountId string) (res string, err *helper.ErrorStruct)
	SetDriverStatusVerified(c context.Context, id string) (res string, err *helper.ErrorStruct)
	DeleteDriver(c context.Context, id string) (res string, err *helper.ErrorStruct)
	AddRoute(c context.Context, data dto.AddRoute) (res models.Route, err *helper.ErrorStruct)
	MonthlyReport(c context.Context, query dto.MonthReport) (res dto.Report, err *helper.ErrorStruct)
	GetImage(c context.Context, id string) (res string, err *helper.ErrorStruct)
}

type DashboardServiceImpl struct {
	DashboardRepo repository.DashboardRepo
}

func (a *DashboardServiceImpl) GetImage(c context.Context, id string) (res string, err *helper.ErrorStruct) {
	resRepo, errRepo := a.DashboardRepo.GetDriverByID(c, id)

	if errRepo != nil {
		var code int

		switch {
		case errors.Is(errRepo, helper.ErrNotFound):
			code = http.StatusNotFound
		default:
			code = http.StatusInternalServerError
		}

		return res, &helper.ErrorStruct{
			Err:  errRepo,
			Code: code,
		}
	}

	if resRepo.KTP == "" {
		return res, &helper.ErrorStruct{
			Err:  errors.New("KTP not found"),
			Code: http.StatusNotFound,
		}
	}

	return resRepo.KTP, nil
}

func (a *DashboardServiceImpl) GetAllHistories(c context.Context) (res []models.Histories, err *helper.ErrorStruct) {
	resRepo, errRepo := a.DashboardRepo.GetAllTripHistories(c)

	if errRepo != nil {
		var code int
		switch {
		case errors.Is(errRepo, helper.ErrNotFound):
			code = http.StatusNotFound
		default:
			code = http.StatusInternalServerError
		}

		return res, &helper.ErrorStruct{
			Code: code,
			Err:  errRepo,
		}
	}

	return resRepo, nil
}

func (a *DashboardServiceImpl) EditAmountRoute(c context.Context, data dto.EditAmount, id string) (res models.Route, err *helper.ErrorStruct) {
	route := models.Route{
		Amount: data.Amount,
	}

	resRepo, errRepo := a.DashboardRepo.EditAmountRoute(c, route, id)

	if errRepo != nil {
		var code int
		switch {
		case errors.Is(errRepo, helper.ErrNotFound):
			code = http.StatusNotFound
		default:
			code = http.StatusInternalServerError
		}

		return res, &helper.ErrorStruct{
			Code: code,
			Err:  errRepo,
		}
	}

	return resRepo, nil
}

func (a *DashboardServiceImpl) MonthlyReport(c context.Context, query dto.MonthReport) (res dto.Report, err *helper.ErrorStruct) {
	resRepo, errRepo := a.DashboardRepo.MonthlyReport(c, query.Month)

	if errRepo != nil {
		var code int
		switch {
		case errors.Is(errRepo, helper.ErrNotFound):
			code = http.StatusNotFound
		default:
			code = http.StatusInternalServerError
		}

		return res, &helper.ErrorStruct{
			Code: code,
			Err:  errRepo,
		}
	}

	return resRepo, nil
}

func (a *DashboardServiceImpl) AddRoute(c context.Context, data dto.AddRoute) (res models.Route, err *helper.ErrorStruct) {
	resRepo, errRepo := a.DashboardRepo.AddRoute(c, models.Route{
		RouteName: data.RouteName,
	})

	if errRepo != nil {
		var code int
		switch {
		case errors.Is(errRepo, helper.ErrNotFound):
			code = http.StatusNotFound
		default:
			code = http.StatusInternalServerError
		}

		return res, &helper.ErrorStruct{
			Code: code,
			Err:  errRepo,
		}
	}

	return resRepo, nil
}

func (a *DashboardServiceImpl) DeleteDriver(c context.Context, id string) (res string, err *helper.ErrorStruct) {
	resRepo, errRepo := a.DashboardRepo.DeleteDriver(c, id)

	if errRepo != nil {
		var code int
		switch {
		case errors.Is(errRepo, helper.ErrNotFound):
			code = http.StatusNotFound
		default:
			code = http.StatusInternalServerError
		}

		return res, &helper.ErrorStruct{
			Code: code,
			Err:  errRepo,
		}
	}

	return resRepo, nil
}

func (a *DashboardServiceImpl) SetDriverStatusVerified(c context.Context, id string) (res string, err *helper.ErrorStruct) {
	resRepo, errRepo := a.DashboardRepo.SetDriverStatusVerified(c, id)

	if errRepo != nil {
		var code int
		switch {
		case errors.Is(errRepo, helper.ErrNotFound):
			code = http.StatusNotFound
		default:
			code = http.StatusInternalServerError
		}

		return res, &helper.ErrorStruct{
			Code: code,
			Err:  errRepo,
		}
	}

	return resRepo, nil
}

func (a *DashboardServiceImpl) GetAllBlockAccount(c context.Context) (res []models.BlockDriver, err *helper.ErrorStruct) {
	resRepo, errRepo := a.DashboardRepo.GetAllBlcokAccount(c)

	if errRepo != nil {
		var code int
		switch {
		case errors.Is(errRepo, helper.ErrNotFound):
			code = http.StatusNotFound
		default:
			code = http.StatusInternalServerError
		}

		return res, &helper.ErrorStruct{
			Code: code,
			Err:  errRepo,
		}
	}

	return resRepo, nil
}

func (a *DashboardServiceImpl) GetAllReviews(c context.Context) (res []models.Reviews, err *helper.ErrorStruct) {
	resRepo, errRepo := a.DashboardRepo.GetAllReview(c)

	if errRepo != nil {
		var code int
		switch {
		case errors.Is(errRepo, helper.ErrNotFound):
			code = http.StatusNotFound
		default:
			code = http.StatusInternalServerError
		}

		return res, &helper.ErrorStruct{
			Code: code,
			Err:  errRepo,
		}
	}

	return resRepo, nil
}

func (a *DashboardServiceImpl) GetReviewById(c context.Context, id string) (res models.Reviews, err *helper.ErrorStruct) {
	resRepo, errRepo := a.DashboardRepo.GetReviewById(c, id)

	if errRepo != nil {
		var code int
		switch {
		case errors.Is(errRepo, helper.ErrNotFound):
			code = http.StatusNotFound
		default:
			code = http.StatusInternalServerError
		}

		return res, &helper.ErrorStruct{
			Code: code,
			Err:  errRepo,
		}
	}

	return resRepo, nil
}

func (a *DashboardServiceImpl) GetAllDrivers(c context.Context, q dto.GetDriverQuery) (res []models.Drivers, err *helper.ErrorStruct) {
	resRepo, errRepo := a.DashboardRepo.GetAllDrivers(c, q.Verified)

	if errRepo != nil {
		var code int
		switch {
		case errors.Is(errRepo, helper.ErrNotFound):
			code = http.StatusNotFound
		default:
			code = http.StatusInternalServerError
		}

		return res, &helper.ErrorStruct{
			Code: code,
			Err:  errRepo,
		}
	}

	for i := range resRepo {
		if resRepo[i].ProfilePicture != "" {
			resRepo[i].ProfilePicture = os.Getenv("BASE_URL") + "/api/driver/images/" + resRepo[i].ID
		}

		if resRepo[i].KTP != "" {
			resRepo[i].KTP = os.Getenv("BASE_URL") + "/api/dashboard/ktp/" + resRepo[i].ID
		}
	}

	return resRepo, nil
}

func (a *DashboardServiceImpl) GetAllPassengers(c context.Context) (res []models.Passengers, err *helper.ErrorStruct) {
	resRepo, errRepo := a.DashboardRepo.GetAllPassengers(c)

	if errRepo != nil {
		var code int
		switch {
		case errors.Is(errRepo, helper.ErrNotFound):
			code = http.StatusNotFound
		default:
			code = http.StatusInternalServerError
		}

		return res, &helper.ErrorStruct{
			Code: code,
			Err:  errRepo,
		}
	}

	return resRepo, nil
}

func (a *DashboardServiceImpl) GetDriverById(c context.Context, id string) (res models.Drivers, err *helper.ErrorStruct) {
	resRepo, errRepo := a.DashboardRepo.GetDriverByID(c, id)

	if errRepo != nil {
		var code int
		switch {
		case errors.Is(errRepo, helper.ErrNotFound):
			code = http.StatusNotFound
		default:
			code = http.StatusInternalServerError
		}

		return res, &helper.ErrorStruct{
			Code: code,
			Err:  errRepo,
		}
	}

	resRepo.ProfilePicture = os.Getenv("BASE_URL") + "/api/driver/images/" + resRepo.ID

	return resRepo, nil
}

func (a *DashboardServiceImpl) GetPassengerById(c context.Context, id string) (res models.Passengers, err *helper.ErrorStruct) {
	resRepo, errRepo := a.DashboardRepo.GetPassengerByID(c, id)

	if errRepo != nil {
		var code int
		switch {
		case errors.Is(errRepo, helper.ErrNotFound):
			code = http.StatusNotFound
		default:
			code = http.StatusInternalServerError
		}

		return res, &helper.ErrorStruct{
			Code: code,
			Err:  errRepo,
		}
	}

	return resRepo, nil
}

func (a *DashboardServiceImpl) UnblockAccount(c context.Context, accountId string) (res string, err *helper.ErrorStruct) {
	resRepo, errRepo := a.DashboardRepo.UnblockAccount(c, accountId)

	if errRepo != nil {
		var code int
		switch {
		case errors.Is(errRepo, helper.ErrNotFound):
			code = http.StatusNotFound
		default:
			code = http.StatusInternalServerError
		}

		return res, &helper.ErrorStruct{
			Code: code,
			Err:  errRepo,
		}
	}

	return resRepo, nil
}

func (a *DashboardServiceImpl) BlockAccount(c context.Context, accountId string) (res models.BlockedAccount, err *helper.ErrorStruct) {
	data := models.BlockedAccount{
		UserID: accountId,
	}

	resRepo, errRepo := a.DashboardRepo.BlockAccount(c, data)

	if errRepo != nil {
		var code int
		switch {
		case errors.Is(errRepo, helper.ErrDuplicateEntry):
			code = http.StatusConflict
		case errors.Is(errRepo, helper.ErrNotFound):
			code = http.StatusNotFound
		default:
			code = http.StatusInternalServerError
		}

		return res, &helper.ErrorStruct{
			Code: code,
			Err:  errRepo,
		}
	}

	return resRepo, nil
}

func NewDashboardService(DashboardRepo repository.DashboardRepo) DashboardService {
	return &DashboardServiceImpl{
		DashboardRepo: DashboardRepo,
	}
}
