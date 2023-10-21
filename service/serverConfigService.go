package service

import (
	"autoapi/graph/model"
	"autoapi/store"
	"autoapi/util"
	"context"
)

type ServerConfigService struct {
	Store store.ServerConfigStorer
}

func NewServerConfigService(store store.ServerConfigStorer) *ServerConfigService {
	return &ServerConfigService{
		Store: store,
	}
}

func (s *ServerConfigService) StoreConfig(ctx context.Context, config model.ServerConfigInput) (*model.ServerConfig, error) {
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

func (s *ServerConfigService) DeleteServerConfig(ctx context.Context, id string) (bool, error) {
	err := s.Store.DeleteConfig(id)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (s *ServerConfigService) DeleteAllServerConfigs(ctx context.Context) (bool, error) {
	err := s.Store.DeleteAllConfigs()
	if err != nil {
		return false, err
	}
	return true, nil
}

func (s *ServerConfigService) GetAllConfigs(ctx context.Context) ([]*model.ServerConfig, error) {
	return s.Store.GetAllConfigs()
}
