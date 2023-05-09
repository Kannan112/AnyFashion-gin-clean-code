// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/kannan112/go-gin-clean-arch/pkg/api"
	"github.com/kannan112/go-gin-clean-arch/pkg/api/handler"
	"github.com/kannan112/go-gin-clean-arch/pkg/config"
	"github.com/kannan112/go-gin-clean-arch/pkg/db"
	"github.com/kannan112/go-gin-clean-arch/pkg/repository"
	"github.com/kannan112/go-gin-clean-arch/pkg/usecase"
)

// Injectors from wire.go:

func InitializeAPI(cfg config.Config) (*http.ServerHTTP, error) {
	gormDB, err := db.ConnectDatabase(cfg)
	if err != nil {
		return nil, err
	}
	userRepository := repository.NewUserRepository(gormDB)
	userUseCase := usecase.NewUserUseCase(userRepository)
	userHandler := handler.NewUserHandler(userUseCase)
	adminRepository := repository.NewAdminRepository(gormDB)
	adminUsecase := usecase.NewAdminUseCase(adminRepository)
	adminHandler := handler.NewAdminSHandler(adminUsecase)
	cartRepository := repository.NewCartRepository(gormDB)
	cartUseCases := usecase.NewCartUseCase(cartRepository)
	cartHandler := handler.NewCartHandler(cartUseCases)
	productRepository := repository.NewProductRepository(gormDB)
	productUseCase := usecase.NewProductUsecase(productRepository)
	productHandler := handler.NewProductHandler(productUseCase)
	orderRepository := repository.NewOrderRepository(gormDB)
	orderUseCase := usecase.NewOrderUseCase(orderRepository)
	orderHandler := handler.NewOrderHandler(orderUseCase)
	serverHTTP := http.NewServerHTTP(userHandler, adminHandler, cartHandler, productHandler, orderHandler)
	return serverHTTP, nil
}
