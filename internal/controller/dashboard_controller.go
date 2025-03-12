package controller

import (
	"net/http"

	"github.com/GabrielMoody/mikronet-dashboard-service/internal/dto"
	"github.com/GabrielMoody/mikronet-dashboard-service/internal/service"
	"github.com/gofiber/fiber/v2"
)

type DashboardController interface {
	SetDriverStatusVerified(c *fiber.Ctx) error
	GetUsers(c *fiber.Ctx) error
	GetUserDetails(c *fiber.Ctx) error
	GetDrivers(c *fiber.Ctx) error
	GetDriverDetails(c *fiber.Ctx) error
	DeleteDriver(c *fiber.Ctx) error
	BlockAccount(c *fiber.Ctx) error
	UnblockAccount(c *fiber.Ctx) error
	GetReviews(c *fiber.Ctx) error
	GetReviewByID(c *fiber.Ctx) error
	GetAllBlockAccount(c *fiber.Ctx) error
}

type DashboardControllerImpl struct {
	DashboardService service.DashboardService
}

func (a *DashboardControllerImpl) GetAllBlockAccount(c *fiber.Ctx) error {
	ctx := c.Context()

	res, err := a.DashboardService.GetAllBlockAccount(ctx)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status": "error",
			"errors": err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "Success",
		"data": fiber.Map{
			"block_accounts": res,
			"count":          len(res),
		},
	})
}

func (a *DashboardControllerImpl) DeleteDriver(c *fiber.Ctx) error {
	ctx := c.Context()
	id := c.Params("id")

	_, err := a.DashboardService.DeleteDriver(ctx, id)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status": "error",
			"errors": err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "Success",
		"message": "Berhasil menghapus akun!",
	})
}

func (a *DashboardControllerImpl) SetDriverStatusVerified(c *fiber.Ctx) error {
	ctx := c.Context()
	id := c.Params("id")

	res, err := a.DashboardService.SetDriverStatusVerified(ctx, id)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status": "error",
			"errors": err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "Success",
		"data":   res,
	})
}

func (a *DashboardControllerImpl) UnblockAccount(c *fiber.Ctx) error {
	ctx := c.Context()
	accountId := c.Params("id")

	_, err := a.DashboardService.UnblockAccount(ctx, accountId)

	if err != nil {
		return c.Status(err.Code).JSON(fiber.Map{
			"status": "error",
			"errors": err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "Success",
		"data":   "Berhasil membuka blokir akun!",
	})
}

func (a *DashboardControllerImpl) GetReviews(c *fiber.Ctx) error {
	ctx := c.Context()

	res, err := a.DashboardService.GetAllReviews(ctx)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "error",
			"errors": err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "Success",
		"data": fiber.Map{
			"reviews": res,
			"count":   len(res),
		},
	})
}

func (a *DashboardControllerImpl) GetReviewByID(c *fiber.Ctx) error {
	ctx := c.Context()
	id := c.Params("id")

	res, err := a.DashboardService.GetReviewById(ctx, id)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "error",
			"errors": err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "Success",
		"data":   res,
	})
}

func (a *DashboardControllerImpl) BlockAccount(c *fiber.Ctx) error {
	ctx := c.Context()
	accountId := c.Params("id")

	_, err := a.DashboardService.BlockAccount(ctx, accountId)

	if err != nil {
		return c.Status(err.Code).JSON(fiber.Map{
			"status": "error",
			"errors": err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "Success",
		"message": "Berhasil memblokir akun!",
	})
}

func (a *DashboardControllerImpl) GetDriverDetails(c *fiber.Ctx) error {
	id := c.Params("id")
	ctx := c.Context()

	res, err := a.DashboardService.GetDriverById(ctx, id)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "error",
			"errors": err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "Success",
		"data":   res,
	})
}

func (a *DashboardControllerImpl) GetDrivers(c *fiber.Ctx) error {
	var q dto.GetDriverQuery
	ctx := c.Context()

	if err := c.QueryParser(&q); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"errors": err.Error(),
		})
	}

	res, err := a.DashboardService.GetAllDrivers(ctx)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "error",
			"errors": err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "Success",
		"data": fiber.Map{
			"drivers": res,
			"count":   len(res),
		},
	})
}

func (a *DashboardControllerImpl) GetUsers(c *fiber.Ctx) error {
	ctx := c.Context()
	res, err := a.DashboardService.GetAllPassengers(ctx)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "error",
			"errors": err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "Success",
		"data": fiber.Map{
			"users": res,
			"count": len(res),
		},
	})
}

func (a *DashboardControllerImpl) GetUserDetails(c *fiber.Ctx) error {
	id := c.Params("id")
	ctx := c.Context()

	res, err := a.DashboardService.GetPassengerById(ctx, id)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status": "error",
			"errors": err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "Success",
		"data":   res,
	})
}

func NewDashboardController(service service.DashboardService) DashboardController {
	return &DashboardControllerImpl{DashboardService: service}
}
