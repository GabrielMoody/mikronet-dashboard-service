package repository

import (
	"context"
	"errors"

	"github.com/GabrielMoody/mikronet-dashboard-service/internal/helper"
	"github.com/GabrielMoody/mikronet-dashboard-service/internal/models"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

type DashboardRepo interface {
	GetAllDrivers(c context.Context, verified *bool) ([]models.Drivers, error)
	GetAllPassengers(c context.Context) ([]models.Passengers, error)
	GetDriverByID(c context.Context, id string) (models.Drivers, error)
	GetPassengerByID(c context.Context, id string) (models.Passengers, error)
	GetAllReview(c context.Context) ([]models.Reviews, error)
	GetReviewById(c context.Context, id string) (models.Reviews, error)
	BlockAccount(c context.Context, data models.BlockedAccount) (models.BlockedAccount, error)
	UnblockAccount(c context.Context, id string) (string, error)
	IsBlocked(c context.Context, id string) (bool, error)
	GetAllBlcokAccount(c context.Context) ([]models.BlockDriver, error)
	SetDriverStatusVerified(c context.Context, id string) (string, error)
	DeleteDriver(c context.Context, id string) (string, error)
	AddRoute(c context.Context, data models.Route) (models.Route, error)
}

type DashboardRepoImpl struct {
	db *gorm.DB
}

func (a *DashboardRepoImpl) AddRoute(c context.Context, data models.Route) (res models.Route, err error) {
	if err := a.db.WithContext(c).Create(&data).Error; err != nil {
		return res, helper.ErrDatabase
	}

	return data, nil
}

func (a *DashboardRepoImpl) DeleteDriver(c context.Context, id string) (res string, err error) {
	if err := a.db.WithContext(c).Delete(&models.DriverDetails{}, "id = ?", id).Error; err != nil {
		return res, helper.ErrDatabase
	}

	return "Berhasil menghapus driver", nil
}

func (a *DashboardRepoImpl) SetDriverStatusVerified(c context.Context, id string) (res string, err error) {
	if err := a.db.WithContext(c).Model(&models.DriverDetails{}).Where("id = ?", id).Update("verified", true).Error; err != nil {
		return res, helper.ErrDatabase
	}

	return "Berhasil memverifikasi driver", nil
}

func (a *DashboardRepoImpl) GetAllReview(c context.Context) (res []models.Reviews, err error) {
	if err := a.db.WithContext(c).Table("reviews").
		Select("reviews.id, p.name AS passenger_name, d.name AS driver_name, reviews.comment AS comment, reviews.star AS star").
		Joins("JOIN passenger_details p ON reviews.passenger_id = p.id").
		Joins("JOIN driver_details d ON reviews.driver_id = d.id").
		Scan(&res).Error; err != nil {
		return res, helper.ErrDatabase
	}

	return res, nil
}

func (a *DashboardRepoImpl) GetReviewById(c context.Context, id string) (res models.Reviews, err error) {
	if err := a.db.WithContext(c).Table("reviews").
		Select("reviews.id, p.name AS passenger_name, d.name AS driver_name, reviews.comment AS comment, reviews.star AS star").
		Joins("JOIN passenger_details p ON reviews.passenger_id = p.id").
		Joins("JOIN driver_details d ON reviews.driver_id = d.id").
		Where("reviews.id = ?", id).
		Scan(&res).Error; err != nil {
		return res, helper.ErrDatabase
	}

	return res, nil
}

func (a *DashboardRepoImpl) GetAllDrivers(c context.Context, verified *bool) (res []models.Drivers, err error) {
	if verified == nil {
		if err := a.db.WithContext(c).Table("driver_details").
			Select("driver_details.id as id, users.email, driver_details.name, driver_details.phone_number, driver_details.license_number, driver_details.sim, driver_details.verified, driver_details.profile_picture, driver_details.status as status").
			Joins("JOIN users ON users.id = driver_details.id").
			Scan(&res).Error; err != nil {
			return res, helper.ErrDatabase
		}
	} else {
		if err := a.db.WithContext(c).Table("driver_details").
			Select("driver_details.id as id, users.email, driver_details.name, driver_details.phone_number, driver_details.license_number, driver_details.sim, driver_details.verified, driver_details.profile_picture, driver_details.status as status").
			Joins("JOIN users ON users.id = driver_details.id").
			Where("verified = ?", verified).
			Scan(&res).Error; err != nil {
			return res, helper.ErrDatabase
		}
	}

	return res, nil
}

func (a *DashboardRepoImpl) GetAllPassengers(c context.Context) (res []models.Passengers, err error) {
	if err := a.db.WithContext(c).Table("passenger_details").
		Select("passenger_details.id as id, users.email, passenger_details.name").
		Joins("JOIN users ON users.id = passenger_details.id").
		Scan(&res).Error; err != nil {
		return res, helper.ErrDatabase
	}

	return res, nil
}

func (a *DashboardRepoImpl) GetDriverByID(c context.Context, id string) (res models.Drivers, err error) {
	if err := a.db.WithContext(c).Table("driver_details").
		Select("driver_details.id as id, users.email, driver_details.name, driver_details.phone_number, driver_details.license_number, driver_details.sim, driver_details.verified, driver_details.profile_picture").
		Joins("JOIN users ON users.id = driver_details.id").
		Scan(&res).Error; err != nil {
		return res, helper.ErrDatabase
	}

	return res, nil
}

func (a *DashboardRepoImpl) GetPassengerByID(c context.Context, id string) (res models.Passengers, err error) {
	if err := a.db.WithContext(c).Table("passenger_details").
		Select("passenger_details.id as id, users.email, passenger_details.name").
		Joins("JOIN users ON users.id = passenger_details.id").
		Scan(&res).Error; err != nil {
		return res, helper.ErrDatabase
	}

	return res, nil
}

func (a *DashboardRepoImpl) GetAllBlcokAccount(c context.Context) (res []models.BlockDriver, err error) {
	if err := a.db.WithContext(c).Table("blocked_accounts").
		Select("user_id as id, users.email as email, driver_details.name as name").
		Joins("JOIN users ON users.id = user_id").
		Joins("JOIN driver_details ON driver_details.id = user_id").Scan(&res).Error; err != nil {
		return res, helper.ErrDatabase
	}

	return res, nil
}

func (a *DashboardRepoImpl) IsBlocked(c context.Context, id string) (bool, error) {
	var res models.BlockedAccount
	if err := a.db.WithContext(c).First(&res, "account_id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, helper.ErrDatabase
	}

	return true, nil
}

func (a *DashboardRepoImpl) UnblockAccount(c context.Context, id string) (res string, err error) {
	if err := a.db.WithContext(c).Delete(&models.BlockedAccount{}, "account_id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, helper.ErrNotFound
		}
		return res, helper.ErrDatabase
	}

	return "Berhasil membuka blokir akun", nil
}

func (a *DashboardRepoImpl) BlockAccount(c context.Context, data models.BlockedAccount) (res models.BlockedAccount, err error) {
	if err := a.db.WithContext(c).Create(&data).Error; err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
			return res, helper.ErrDuplicateEntry
		}
		return res, helper.ErrDatabase
	}

	return res, nil
}

func NewDashboardRepo(db *gorm.DB) DashboardRepo {
	return &DashboardRepoImpl{
		db: db,
	}
}
