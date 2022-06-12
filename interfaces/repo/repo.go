package repo

import (
	"app/domain"
	"context"
)

type MonitorRepo interface {
	Save(ctx context.Context, monitor *domain.Monitor) (string, error)
	FindById(ctx context.Context, id string) (*domain.Monitor, error)
	FindAll(ctx context.Context) ([]domain.Monitor, error)
	FindAllFree(ctx context.Context) ([]domain.Monitor, error)
	FindAllCompleted(ctx context.Context) ([]domain.Monitor, error)
}
