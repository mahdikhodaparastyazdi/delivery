package command

import (
	"context"
	"delivery/internal/config"
	"delivery/internal/constants"
	courior_consumer "delivery/internal/consumers/courior_consumer"
	courior_resolver "delivery/internal/factories/courior_resolver"
	"delivery/internal/repositories"
	"delivery/internal/tasks"
	"delivery/pkg/asynq"
	"fmt"

	log "delivery/pkg/logger"
	"delivery/pkg/logger/shoplog"
	"delivery/pkg/mysql"

	"github.com/spf13/cobra"
)

type Consumer struct {
	courior bool
	logger  log.Logger
}

func (cmd Consumer) Command(ctx context.Context, cfg *config.Config) *cobra.Command {
	cmd.logger = shoplog.NewStdOutLogger(cfg.LogLevel, "courior:consumer:command")

	consumerCmd := &cobra.Command{
		Use:   "consumer",
		Short: "run auth consumer",
		Run: func(_ *cobra.Command, _ []string) {
			cmd.main(ctx, cfg)
		},
	}

	consumerCmd.Flags().BoolVarP(&cmd.courior, "courior", "", false, "Run high consumer")

	return consumerCmd
}
func (cmd Consumer) main(ctx context.Context, cfg *config.Config) {
	if cmd.courior {
		cmd.couriorConsumer(ctx, cfg)
	}
}

func (cmd Consumer) couriorConsumer(ctx context.Context, cfg *config.Config) {
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

	_ = repositories.NewCouriorRepository(gormDB)
	logger := shoplog.NewStdOutLogger(cfg.LogLevel, "delivery:courior:provider-resolver")
	_ = courior_resolver.NewResolver(cfg.AppEnv, logger)

	logger = shoplog.NewStdOutLogger(cfg.LogLevel, "delivery:consumer:courior:courior")
	couriorConsumer := courior_consumer.New(logger)

	logger = shoplog.NewStdOutLogger(cfg.LogLevel, "delivery:courior:asynq-courior-server")
	server := asynq.NewServer(logger, cfg.Database.Redis, constants.QUEUE_COURIOR, cfg.CouriorConsumer.AsynqHighWorkerCount)

	logger = shoplog.NewStdOutLogger(cfg.LogLevel, "delivery:courior:courior-worker")
	worker := tasks.NewWorker(server, couriorConsumer, logger)
	if err := worker.StartWorker(constants.QUEUE_COURIOR); err != nil {
		cmd.logger.Error(err)
		return
	}
}
