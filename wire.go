package main

import (
	handler "github.com/acxe21/crud-app.git/handlers"
	"github.com/acxe21/crud-app.git/repository"
	service "github.com/acxe21/crud-app.git/services"
	"github.com/google/wire"
	"gorm.io/gorm"
)

func ProductHandler(db *gorm.DB) handler.ProductHandler {
	wire.Build(repository.ProvideProductRepository, service.ProvideProductService, handler.ProvideProductHandle)
	return handler.ProductHandler{}
}
