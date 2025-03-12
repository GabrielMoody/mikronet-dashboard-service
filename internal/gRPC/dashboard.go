package gRPC

import (
	"context"

	"github.com/GabrielMoody/mikronet-dashboard-service/internal/pb"
	"github.com/GabrielMoody/mikronet-dashboard-service/internal/repository"
)

type GRPC struct {
	pb.UnimplementedDashboardServiceServer
	repo repository.DashboardRepo
}

func NewGRPC(repo repository.DashboardRepo) *GRPC {
	return &GRPC{
		repo: repo,
	}
}

func (a *GRPC) IsBlocked(ctx context.Context, data *pb.IsBlockedReq) (*pb.IsBlockedRes, error) {
	res, err := a.repo.IsBlocked(ctx, data.Id)

	if err != nil {
		return nil, err
	}

	return &pb.IsBlockedRes{
		IsBlocked: res,
	}, nil
}
