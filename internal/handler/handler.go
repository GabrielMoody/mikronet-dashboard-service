package handler

import (
	"github.com/GabrielMoody/mikronet-dashboard-service/internal/controller"
	"github.com/GabrielMoody/mikronet-dashboard-service/internal/gRPC"
	"github.com/GabrielMoody/mikronet-dashboard-service/internal/middleware"
	"github.com/GabrielMoody/mikronet-dashboard-service/internal/pb"
	"github.com/GabrielMoody/mikronet-dashboard-service/internal/repository"
	"github.com/GabrielMoody/mikronet-dashboard-service/internal/service"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func DashboardHandler(r fiber.Router, db *gorm.DB, driver pb.DriverServiceClient, user pb.UserServiceClient) {
	repo := repository.NewDashboardRepo(db)
	serviceDashboard := service.NewDashboardService(repo)
	controllerDashboard := controller.NewDashboardController(serviceDashboard, driver, user)

	api := r.Group("/")

	api.Get("/users", controllerDashboard.GetUsers)
	api.Get("/users/:id", controllerDashboard.GetUserDetails)
	api.Delete("/users/:id", middleware.ValidateDashboardRole, controllerDashboard.DeleteUser)

	api.Get("/drivers", controllerDashboard.GetDrivers)
	api.Get("/drivers/:id", controllerDashboard.GetDriverDetails)
	api.Post("/drivers/verified/:id", middleware.ValidateDashboardRole, controllerDashboard.SetDriverStatusVerified)
	api.Delete("/drivers/:id", middleware.ValidateDashboardRole, controllerDashboard.DeleteDriver)

	api.Get("/owners", middleware.ValidateDashboardRole, controllerDashboard.GetBusinessOwners)
	api.Get("/owners/blocked", middleware.ValidateDashboardRole, controllerDashboard.GetBlockedBusinessOwners)
	api.Get("/owners/unverified", middleware.ValidateDashboardRole, controllerDashboard.GetUnverifiedBusinessOwners)
	api.Get("/owners/:id", middleware.ValidateDashboardRole, controllerDashboard.GetBusinessOwnerDetails)
	api.Put("/owners/verified/:id", middleware.ValidateDashboardRole, controllerDashboard.SetOwnerStatusVerified)

	api.Post("/block/:id", middleware.ValidateDashboardRole, controllerDashboard.BlockAccount)
	api.Delete("/block/:id", middleware.ValidateDashboardRole, controllerDashboard.UnblockAccount)

	api.Get("/reviews", controllerDashboard.GetReviews)
	api.Get("/reviews/:id", controllerDashboard.GetReviewByID)
}

func GRPCHandler(db *gorm.DB) *gRPC.GRPC {
	repo := repository.NewDashboardRepo(db)
	grpc := gRPC.NewGRPC(repo)

	return grpc
}
