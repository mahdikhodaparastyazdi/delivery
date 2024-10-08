package config

import (
	"time"

	log "delivery/pkg/logger"
	"delivery/pkg/mysql"

	"delivery/pkg/redis"

	"delivery/pkg/sentry"
)

type LogLevel string

type AppEnv string

var ValidTimeSlots = []int{9, 11, 13, 15, 17, 19, 21, 23}

const (
	ProductionEnv AppEnv = "production"
	StageEnv      AppEnv = "stage"
	DevelopEnv    AppEnv = "develop"
	LocalEnv      AppEnv = "locale"
)

type (
	Config struct {
		AppEnv          AppEnv
		AppDebug        bool
		LogLevel        log.LogLevelStr
		HealthToken     string
		HTTP            HTTP
		Database        Database
		Sentry          *sentry.Config
		CoreBaseUrl     string
		CouriorBaseUrl  string
		Tz              string
		CouriorConsumer CouriorConsumer
		APIKeyCore      string
		APIKey3PL       string
	}

	HTTP struct {
		Host string
		Port int
	}

	Database struct {
		MySQL mysql.Config
		Redis redis.Config
	}

	Sentry struct {
		Active           bool
		Dsn              string
		EnableTracing    bool
		TracesSampleRate float64
	}

	CouriorConsumer struct {
		AsynqHighWorkerCount int
		AsynqLowWorkerCount  int
		AsynqLowMaxRetry     int
		AsynqHighMaxRetry    int
		AsynqTimeoutSeconds  time.Duration
	}
)
