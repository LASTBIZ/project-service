package main

import (
	"context"
	"google.golang.org/grpc"
	"lastbiz/project-service/internal/config"
	"lastbiz/project-service/internal/project"
	"lastbiz/project-service/internal/project/services"
	"lastbiz/project-service/pkg/logging"
	"lastbiz/project-service/pkg/postgres"
	project1 "lastbiz/project-service/pkg/project"
	"net"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	logging.Info(ctx, "config initializing")
	cfg := config.GetConfig()

	pgconfig := postgres.NewPGConfig(cfg.Postgres.User, cfg.Postgres.Password,
		cfg.Postgres.Host, cfg.Postgres.Port, cfg.Postgres.DB)

	pgClient := postgres.NewClient(ctx, 5, time.Second*5, pgconfig)
	projectStorage := project.NewProjectStorage(pgClient)
	projectService := services.NewProjectService(projectStorage)

	logging.Info(ctx, "run application")
	lis, err := net.Listen("tcp", "0.0.0.0:"+cfg.GRPCPort)

	if err != nil {
		logging.GetLogger().Fatal(err)
	}

	grpcServer := grpc.NewServer()
	project1.RegisterProjectServiceServer(grpcServer, projectService)

	if err := grpcServer.Serve(lis); err != nil {
		logging.GetLogger().Fatal(err)
	}
}
