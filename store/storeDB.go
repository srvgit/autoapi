package store

import "autoapi/graph/model"

type ServerConfigStorer interface {
	StoreConfig(config *model.ServerConfig) (*model.ServerConfig, error)
	GetAllConfigs() ([]*model.ServerConfig, error)
	DeleteConfig(id string) error
	DeleteAllConfigs() error
}
