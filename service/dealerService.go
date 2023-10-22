package service

import (
	"autoapi/graph/model"
	"autoapi/store"
	"autoapi/util"
	"context"
)

type DealerService struct {
	Store store.ServerConfigStorer
}

func NewDealerService(store store.ServerConfigStorer) *ServerConfigService {
	return &ServerConfigService{
		Store: store,
	}
}

func (s *DealerService) StoreConfig(ctx context.Context, config model.ServerConfigInput) (*model.ServerConfig, error) {
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

func (s *DealerService) DeleteServerConfig(ctx context.Context, id string) (bool, error) {
	err := s.Store.DeleteConfig(id)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (s *DealerService) DeleteAllServerConfigs(ctx context.Context) (bool, error) {
	err := s.Store.DeleteAllConfigs()
	if err != nil {
		return false, err
	}
	return true, nil
}

func (s *DealerService) GetAllConfigs(ctx context.Context) ([]*model.ServerConfig, error) {
	return s.Store.GetAllConfigs()
}
