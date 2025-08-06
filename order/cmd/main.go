package main

import (
	"log"

	"github.com/joseguilherme-fs/microservices/order/config"
	"github.com/joseguilherme-fs/microservices/order/internal/adapters/db"
	// "github.com/joseguilherme-fs/microservices/order/internal/adapters/rest"
	"github.com/joseguilherme-fs/microservices/order/internal/adapters/grpc"
	"github.com/joseguilherme-fs/microservices/order/internal/application/core/api"
)

func main() {
	dbAdapter, err := db.NewAdapter(config.GetDataSourceURL())
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	application := api.NewApplication(dbAdapter)
	grpcAdapter := grpc.NewAdapter(application, config.GetApplicationPort())
	grpcAdapter.Run()
}
