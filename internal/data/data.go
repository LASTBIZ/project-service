package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	slog "log"
	"os"
	"project-service/api/video"
	"project-service/internal/biz"
	"project-service/internal/conf"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData,
	NewDB,
	NewTransaction,
	NewProjectRepo,
	NewInvestorRepo,
	NewRoadmapRepo,
	NewCategoryRepo,
	NewVideoServiceClient)

// Data .
type Data struct {
	db *gorm.DB
}

// NewData .
func NewData(c *conf.Data, logger log.Logger, db *gorm.DB) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{db: db}, cleanup, nil
}

type contextTxKey struct{}

func NewTransaction(d *Data) biz.Transaction {
	return d
}

func (d *Data) DB(ctx context.Context) *gorm.DB {
	tx, ok := ctx.Value(contextTxKey{}).(*gorm.DB)
	if ok {
		return tx
	}
	return d.db
}

func (d *Data) ExecTx(ctx context.Context, fn func(ctx context.Context) error) error {
	return d.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		ctx = context.WithValue(ctx, contextTxKey{}, tx)
		return fn(ctx)
	})
}

func NewDB(c *conf.Data) *gorm.DB {
	newLogger := logger.New(
		slog.New(os.Stdout, "\r\n", slog.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			Colorful:      true,
			LogLevel:      logger.Info,
		},
	)
	log.Info("failed opening connection to ")
	db, err := gorm.Open(postgres.Open(c.Database.Source), &gorm.Config{
		Logger:                                   newLogger,
		DisableForeignKeyConstraintWhenMigrating: true,
		//NamingStrategy:                           schema.NamingStrategy{},
	})

	if err != nil {
		log.Errorf("failed opening connection to postgres: %v", err)
		panic("failed to connect database")
	}
	db.AutoMigrate(&Project{}, &Category{}, &Investor{}, &RoadMap{}, &InvestProject{})
	return db
}

func NewVideoServiceClient(sr *conf.Services) video.VideoClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint(sr.Video.Endpoint),
		grpc.WithMiddleware(
			tracing.Client(),
			recovery.Recovery()),
		grpc.WithTimeout(2*time.Second),
	)
	if err != nil {
		panic(err)
	}
	c := video.NewVideoClient(conn)
	return c
}
