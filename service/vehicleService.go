package service

import (
	"autoapi/graph/model"
	"autoapi/store"
	"context"
)

type VehicleService struct {
	Store store.ServerConfigStorer
}

func NewVehicleService(store store.ServerConfigStorer) *ServerConfigService {
	return &ServerConfigService{
		Store: store,
	}
}

func (s *VehicleService) DeleteServerConfig(ctx context.Context, id string) (bool, error) {
	err := s.Store.DeleteConfig(id)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (s *VehicleService) DeleteAllServerConfigs(ctx context.Context) (bool, error) {
	err := s.Store.DeleteAllConfigs()
	if err != nil {
		return false, err
	}
	return true, nil
}

func (s *VehicleService) GetAllConfigs(ctx context.Context) ([]*model.ServerConfig, error) {
	return s.Store.GetAllConfigs()
}
