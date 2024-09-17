package app

import (
	"os"

	log "delivery/pkg/logger"
	zerologLogger "delivery/pkg/logger/zerologger"

	"github.com/rs/zerolog"
)

func InitLogger() {
	z := zerolog.New(os.Stdout).With().Timestamp().Logger()
	log.SetDefaultLogger(zerologLogger.NewFromZerolog(z))
}
