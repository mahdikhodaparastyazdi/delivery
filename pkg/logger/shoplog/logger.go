package shoplog

import (
	"os"

	log "delivery/pkg/logger"

	"github.com/rs/zerolog"
)

const skipStackFrameCount = 2

type ShoplogLogger struct {
	shoplog zerolog.Logger
	prefix  string
}

func NewStdOutLogger(logLevel log.LogLevelStr, prefix string) ShoplogLogger {
	return ShoplogLogger{
		shoplog: zerolog.New(os.Stdout).With().Timestamp().Logger().Level(convertToZeroLogLevel(log.StrLogLevelToInt(logLevel))),
		prefix:  prefix,
	}
}

func NewFromShoplog(z zerolog.Logger) ShoplogLogger {
	return ShoplogLogger{
		shoplog: z,
	}
}

func (z ShoplogLogger) Fatal(message any, data ...map[string]any) {
	z.shoplog.Fatal().Caller(skipStackFrameCount).Interface("context", data).Msgf("%v:%v", z.prefix, message)
}

func (z ShoplogLogger) Error(message any, data ...map[string]any) {
	z.shoplog.Error().Caller(skipStackFrameCount).Interface("context", data).Msgf("%v:%v", z.prefix, message)
}

func (z ShoplogLogger) Warn(message any, data ...map[string]any) {
	z.shoplog.Warn().Caller(skipStackFrameCount).Interface("context", data).Msgf("%v:%v", z.prefix, message)
}

func (z ShoplogLogger) Info(message any, data ...map[string]any) {
	z.shoplog.Info().Caller(skipStackFrameCount).Interface("context", data).Msgf("%v:%v", z.prefix, message)
}

func (z ShoplogLogger) Debug(message any, data ...map[string]any) {
	z.shoplog.Debug().Interface("context", data).Msgf("%v:%v", z.prefix, message)
}

func (z ShoplogLogger) CloneWithPrefix(prefix string) log.ClonableLogger {
	z.prefix += ":" + prefix
	return z
}
