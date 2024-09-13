package main

import (
	"context"
	"delivery/cmd/command"
	"delivery/internal/app"
	"delivery/internal/config"
	"delivery/internal/version"
	"fmt"
	"os/signal"
	"syscall"

	log "delivery/pkg/logger"
	sentryPkg "delivery/pkg/sentry"

	"github.com/spf13/cobra"
)

func main() {
	fmt.Printf("Version: %v\nRelease Date: %v\nCommit Hash: %v\n\n\n", version.Version, version.ReleaseDate, version.CommitHash)
	const description = "shop service"
	root := &cobra.Command{Short: description}

	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	app.InitLogger()

	sentry := sentryPkg.NewSentry(&sentryPkg.Config{
		Dsn:              cfg.Sentry.Dsn,
		EnableTracing:    cfg.Sentry.EnableTracing,
		TracesSampleRate: cfg.Sentry.TracesSampleRate,
		Active:           cfg.Sentry.Active,
	})
	err = sentry.InitSentry()
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	root.AddCommand(
		command.Version{}.Command(),
		command.Server{}.Command(ctx, cfg),
		command.Consumer{}.Command(ctx, cfg),
	)

	if err := root.Execute(); err != nil {
		log.Fatal(fmt.Sprintf("failed go execute root command: \n %s", err.Error()))
	}
}
