package repo

import (
	"app/domain"
	"context"

	"gopkg.in/mgo.v2"
)

// NewPostStore initializes the Posts store with given mongo db handle.
func NewMonitorStore(db *mgo.Database) *MonitorStore {
	return &MonitorStore{
		db: db,
	}
}

// PostStore manages persistence and retrieval of posts.
type MonitorStore struct {
	db *mgo.Database
}

func (m *MonitorStore) Save(ctx context.Context, monitor *domain.Monitor) (string, error) {
	//TODO
	return "", nil
}

func (m *MonitorStore) FindById(ctx context.Context, id string) (*domain.Monitor, error) {
	//TODO
	return nil, nil
}

func (m *MonitorStore) FindAll(ctx context.Context) ([]domain.Monitor, error) {
	//TODO
	return nil, nil
}

func (m *MonitorStore) FindAllFree(ctx context.Context) ([]domain.Monitor, error) {
	//TODO
	return nil, nil
}

func (m *MonitorStore) FindAllCompleted(ctx context.Context) ([]domain.Monitor, error) {
	//TODO
	return nil, nil
}
