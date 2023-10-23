package store

import "autoapi/graph/model"

type ServerConfigStorer interface {
	CreateService(config *model.ServerConfig) (*model.ServerConfig, error)
	GetAllConfigs() ([]*model.ServerConfig, error)
	DeleteConfig(id string) error
	DeleteAllConfigs() error
}
