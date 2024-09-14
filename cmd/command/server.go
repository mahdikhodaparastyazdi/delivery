package command

import (
	"context"
	"delivery/internal/api/rest"
	delivery_handler "delivery/internal/api/rest/handlers/delivery"
	"delivery/internal/api/rest/middleware"
	"delivery/internal/api/rest/transformer"
	"delivery/internal/config"
	"delivery/internal/repositories"
	clientservice "delivery/internal/services/client"
	delivery_service "delivery/internal/services/delivery"

	tasks "delivery/internal/tasks/send_courior"
	"delivery/pkg/asynq"
	"fmt"

	log "delivery/pkg/logger"
	"delivery/pkg/logger/shoplog"
	"delivery/pkg/mysql"
	response_formatter "delivery/pkg/response_formatter"

	"github.com/spf13/cobra"
)

type Server struct {
	logger log.Logger
}

func (cmd Server) Command(ctx context.Context, cfg *config.Config) *cobra.Command {
	cmd.logger = shoplog.NewStdOutLogger(cfg.LogLevel, "delivery:server:command")

	return &cobra.Command{
		Use:   "server",
		Short: "run setting server",
		Run: func(_ *cobra.Command, _ []string) {
			cmd.main(ctx, cfg)
		},
	}
}

func (cmd Server) main(ctx context.Context, cfg *config.Config) {
	db, err := mysql.NewClient(ctx, &cfg.Database.MySQL)
	if err != nil {
		cmd.logger.Fatal("failed to connect to mysql database", log.J{"error": err.Error()})
		return
	}
	gormDB, err := mysql.NewGormWithInstance(db, cfg.AppDebug)
	if err != nil {
		cmd.logger.Fatal("failed to connect to mysql database", log.J{"error": err.Error()})
		return
	}

	err = mysql.Migrate(db)
	if err != nil {
		cmd.logger.Fatal(fmt.Errorf("mysql migration failed: %w", err))
	}

	deliveryRepo := repositories.NewCouriorRepository(gormDB)
	clientRepo := repositories.NewClientRepository(gormDB)

	clientService := clientservice.New(clientRepo)
	asynqClient := asynq.NewClient(cfg.Database.Redis)
	Queue3PL := tasks.NewQueue3PL(asynqClient, cfg.CouriorConsumer.AsynqLowMaxRetry, cfg.CouriorConsumer.AsynqTimeoutSeconds)
	deliveryService := delivery_service.New(deliveryRepo, Queue3PL)

	deliverTr := transformer.NewDeliveryTransformer()

	logger := shoplog.NewStdOutLogger(cfg.LogLevel, "delivery::api:response-formatter-pkg")
	respFormatter := response_formatter.NewResponseFormatter(logger)

	deliveryHandler := delivery_handler.New(deliveryService, respFormatter, deliverTr)

	internalMiddleware := middleware.NewInternalMiddleware(clientService)

	logger = shoplog.NewStdOutLogger(cfg.LogLevel, "delivery:api:server")
	server := rest.New(logger)
	server.SetupAPIRoutes(
		internalMiddleware,
		deliveryHandler,
	)
	if err := server.Serve(ctx, fmt.Sprintf("%s:%d", cfg.HTTP.Host, cfg.HTTP.Port)); err != nil {
		cmd.logger.Fatal(err)
	}
}
