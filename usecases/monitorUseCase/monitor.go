package monitorUseCase

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"

	"app/domain"
	"app/interfaces/repo"
)

type MonitorService interface {
	Add(ctx context.Context, monitor domain.Monitor) (string, error)
	FindById(ctx context.Context, id string) (*domain.Monitor, error)
	FindAll(ctx context.Context) ([]domain.Monitor, error)
	FindAllFree(ctx context.Context) ([]domain.Monitor, error)
	FindAllCompleted(ctx context.Context) ([]domain.Monitor, error)
}

type MonitorUseCase struct {
	repo repo.MonitorRepo
}

func NewMonitorUserCase(r repo.MonitorRepo) MonitorService {
	return MonitorUseCase{repo: r}
}

// Register creates a new user in the system using the given user object.
func (m MonitorUseCase) Add(ctx context.Context, monitor domain.Monitor) (string, error) {

	monitor.ID = uuid.NewString()
	monitor.CreatedAt = time.Now()

	m.repo.Save(ctx, &monitor)

	var err error

	periodStr := ctx.Value("period").(string)
	monitor.Period, err = time.ParseDuration(periodStr)
	if err != nil {
		return monitor.ID, err
	}

	frequencyStr := ctx.Value("freq").(string)
	monitor.Frequency, err = time.ParseDuration(frequencyStr)
	if err != nil {
		return monitor.ID, err
	}

	if err := monitor.Validate(); err != nil {
		return monitor.ID, err
	}

	return monitor.ID, nil
}

func (m MonitorUseCase) FindAll(ctx context.Context) ([]domain.Monitor, error) {
	data, err := m.repo.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (m MonitorUseCase) FindById(ctx context.Context, id string) (*domain.Monitor, error) {
	data, err := m.repo.FindById(ctx, id)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (s MonitorUseCase) FindAllFree(ctx context.Context) ([]domain.Monitor, error) {
	//TODO
	return nil, fmt.Errorf("need implementation")
}

func (s MonitorUseCase) FindAllCompleted(ctx context.Context) ([]domain.Monitor, error) {
	//TODO
	return nil, fmt.Errorf("need implementation")
}
