package config

import (
	"errors"
	"fmt"

	log "delivery/pkg/logger"
	"delivery/pkg/redis"
	"delivery/pkg/sentry"

	"delivery/pkg/mysql"

	"github.com/spf13/viper"
)

func Load() (*Config, error) {
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.AllowEmptyEnv(true)
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		if !errors.As(err, &viper.ConfigFileNotFoundError{}) {
			return nil, fmt.Errorf("reading config: %w", err)
		}
	}
	appEnv := loadString("APP_ENV")
	debug := loadBool("APP_DEBUG")
	cfg := Config{
		AppEnv:   AppEnv(appEnv),
		Locale:   loadString("LOCALE"),
		AppDebug: debug,
		Tz:       loadString("TZ"),
		LogLevel: log.LogLevelStr(loadString("LOG_LEVEL")),
		HTTP: HTTP{
			Host: loadString("API_HTTP_HOST"),
			Port: loadInt("API_HTTP_PORT"),
		},
		Database: Database{
			MySQL: mysql.Config{
				Host:         loadString("MYSQL_HOST"),
				Port:         loadInt("MYSQL_PORT"),
				Username:     loadString("MYSQL_USER"),
				Password:     loadString("MYSQL_PASSWORD"),
				DatabaseName: loadString("MYSQL_DATABASE"),
				Timezone:     loadString("TZ"),
			},
			Redis: redis.Config{
				Host:     loadString("REDIS_HOST"),
				Port:     loadString("REDIS_PORT"),
				Password: loadString("REDIS_PASSWORD"),
				Database: loadInt("REDIS_DATABASE"),
			},
		},
		Sentry: &sentry.Config{
			Dsn:                loadString("SENTRY_DSN"),
			EnableTracing:      loadBool("SENTRY_ENABLE_TRACING"),
			TracesSampleRate:   loadFloat64("SENTRY_TRACES_SAMPLE_RATE"),
			Active:             loadBool("SENTRY_ACTIVE"),
			Debug:              debug,
			Environment:        appEnv,
			SampleRate:         loadFloat64("SENTRY_SAMPLE_RATE"),
			ProfilesSampleRate: loadFloat64("SENTRY_PROFILES_SAMPLE_RATE"),
		},
		CouriorConsumer: CouriorConsumer{
			AsynqHighWorkerCount: loadInt("ASYNQ_HIGH_WORKER_COUNT"),
			AsynqLowWorkerCount:  loadInt("ASYNQ_LOW_WORKER_COUNT"),
			AsynqLowMaxRetry:     loadInt("ASYNQ_JOB_LOW_MAX_RETRY"),
			AsynqHighMaxRetry:    loadInt("ASYNQ_JOB_HIGH_MAX_RETRY"),
			AsynqTimeoutSeconds:  loadDuration("ASYNQ_JOB_TIMEOUT_IN_SECONDS"),
		},
	}

	return &cfg, nil
}
