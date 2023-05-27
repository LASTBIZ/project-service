package server

import (
	"project-service/api/category"
	"project-service/api/investor"
	"project-service/api/project"
	"project-service/api/roadmap"
	"project-service/internal/conf"
	"project-service/internal/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(c *conf.Server,
	cat *service.CategoryService,
	invest *service.InvestorService,
	pro *service.ProjectService,
	road *service.RoadmapService,
	logger log.Logger) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
		),
	}
	if c.Grpc.Network != "" {
		opts = append(opts, grpc.Network(c.Grpc.Network))
	}
	if c.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Grpc.Addr))
	}
	if c.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Grpc.Timeout.AsDuration()))
	}
	srv := grpc.NewServer(opts...)
	category.RegisterCategoryServer(srv, cat)
	investor.RegisterInvestorServer(srv, invest)
	project.RegisterProjectServer(srv, pro)
	roadmap.RegisterRoadmapServer(srv, road)
	return srv
}
