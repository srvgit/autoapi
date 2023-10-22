package graph

import "autoapi/service"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	ServerConfigService *service.ServerConfigService
	VehicleService      *service.VehicleService
	DealerService       *service.DealerService
}
