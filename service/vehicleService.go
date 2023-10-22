package service

import (
	"autoapi/graph/model"
	"autoapi/store"
	"autoapi/util"
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

func (s *VehicleService) StoreConfig(ctx context.Context, config model.ServerConfigInput) (*model.ServerConfig, error) {
	conf := &model.ServerConfig{
		ID:               util.GenerateUUID(),
		GraphPackagePath: config.GraphPackagePath,
		PlaygroundPath:   config.PlaygroundPath,
		QueryPath:        config.QueryPath,
		GinMode:          model.GinMode(config.GinMode),
		Port:             config.Port,
	}

	return s.Store.StoreConfig(conf)
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
