package command

import (
	"context"
	"delivery/internal/config"
	"delivery/internal/constants"
	receiver_consumer "delivery/internal/consumers/received_courior_consumer"
	courior_consumer "delivery/internal/consumers/send_courior_consumer"
	courior_resolver "delivery/internal/factories/courior_resolver"
	"delivery/internal/repositories"
	receive_tasks "delivery/internal/tasks/received_courior_status"
	received_tasks "delivery/internal/tasks/received_courior_status"
	send_tasks "delivery/internal/tasks/send_courior"

	"delivery/pkg/asynq"

	log "delivery/pkg/logger"
	"delivery/pkg/logger/shoplog"
	"delivery/pkg/mysql"

	"github.com/spf13/cobra"
)

type Consumer struct {
	courior  bool
	receiver bool
	logger   log.Logger
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

	consumerCmd.Flags().BoolVarP(&cmd.courior, "courior", "", false, "run courior consumer")
	consumerCmd.Flags().BoolVarP(&cmd.receiver, "receiver", "", false, "run receiver consumer")
	return consumerCmd
}
func (cmd Consumer) main(ctx context.Context, cfg *config.Config) {
	if cmd.courior {
		cmd.couriorConsumer(ctx, cfg)
	}
	if cmd.receiver {
		cmd.receiverConsumer(ctx, cfg)
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
	asynqClient := asynq.NewClient(cfg.Database.Redis)
	Queue3PL := send_tasks.NewQueue3PL(asynqClient, cfg.CouriorConsumer.AsynqLowMaxRetry, cfg.CouriorConsumer.AsynqTimeoutSeconds)

	_ = repositories.NewCouriorRepository(gormDB)
	logger := shoplog.NewStdOutLogger(cfg.LogLevel, "delivery:courior:provider-resolver")
	resolver3PL := courior_resolver.NewResolver(cfg.AppEnv, logger)

	logger = shoplog.NewStdOutLogger(cfg.LogLevel, "delivery:consumer:courior")
	couriorConsumer := courior_consumer.New(logger, Queue3PL, resolver3PL)

	logger = shoplog.NewStdOutLogger(cfg.LogLevel, "delivery:courior:asynq-courior-server")
	server := asynq.NewServer(logger, cfg.Database.Redis, constants.SEND_COURIOR, cfg.CouriorConsumer.AsynqHighWorkerCount)

	logger = shoplog.NewStdOutLogger(cfg.LogLevel, "delivery:courior:worker")
	worker := send_tasks.NewWorker(server, couriorConsumer, logger)
	if err := worker.StartWorker(constants.SEND_COURIOR); err != nil {
		cmd.logger.Error(err)
		return
	}
}
func (cmd Consumer) receiverConsumer(ctx context.Context, cfg *config.Config) {
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

	_ = repositories.NewCouriorRepository(gormDB)
	logger := shoplog.NewStdOutLogger(cfg.LogLevel, "delivery:receiver:provider-resolver")
	_ = courior_resolver.NewResolver(cfg.AppEnv, logger)

	asynqClient := asynq.NewClient(cfg.Database.Redis)
	QueueCore := received_tasks.NewQueue3PL(asynqClient, cfg.CouriorConsumer.AsynqLowMaxRetry, cfg.CouriorConsumer.AsynqTimeoutSeconds)

	logger = shoplog.NewStdOutLogger(cfg.LogLevel, "delivery:consumer:receiver")
	receiverConsumer := receiver_consumer.New(logger, QueueCore)

	logger = shoplog.NewStdOutLogger(cfg.LogLevel, "delivery:receiver:asynq-receiver-server")
	server := asynq.NewServer(logger, cfg.Database.Redis, constants.SEND_COURIOR, cfg.CouriorConsumer.AsynqHighWorkerCount)

	logger = shoplog.NewStdOutLogger(cfg.LogLevel, "delivery:receiver:worker")
	worker := receive_tasks.NewWorker(server, receiverConsumer, logger)
	if err := worker.StartWorker(constants.SEND_COURIOR); err != nil {
		cmd.logger.Error(err)
		return
	}
}
