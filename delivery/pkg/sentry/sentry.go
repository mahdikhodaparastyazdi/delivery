package sentry

import (
	"fmt"

	"github.com/getsentry/sentry-go"
)

type Sentry struct {
	Dsn                string
	EnableTracing      bool
	TracesSampleRate   float64
	Active             bool
	Debug              bool
	Environment        string
	SampleRate         float64
	ProfilesSampleRate float64
}

func NewSentry(cfg *Config) *Sentry {
	return &Sentry{
		Dsn:                cfg.Dsn,
		EnableTracing:      cfg.EnableTracing,
		TracesSampleRate:   cfg.TracesSampleRate,
		Active:             cfg.Active,
		Environment:        cfg.Environment,
		SampleRate:         cfg.SampleRate,
		ProfilesSampleRate: cfg.ProfilesSampleRate,
		Debug:              cfg.Debug,
	}
}

func (s *Sentry) InitSentry() error {
	if !s.Active {
		return nil
	}
	if err := sentry.Init(sentry.ClientOptions{
		Dsn:           s.Dsn,
		EnableTracing: s.EnableTracing,
		TracesSampleRate:   s.TracesSampleRate,
		Debug:              s.Debug,
		AttachStacktrace:   true,
		SampleRate:         s.SampleRate,
		ProfilesSampleRate: s.ProfilesSampleRate,
		Environment:        s.Environment,
	}); err != nil {
		return fmt.Errorf("Sentry initialization failed: %v", err)
	}

	return nil
}
