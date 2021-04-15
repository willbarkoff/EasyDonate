package logging

import (
	"github.com/stripe/stripe-go"
	"strings"
)

type StripeLogger struct {
	Level stripe.Level
}

func (s StripeLogger) Debugf(format string, v ...interface{}) {
	if s.Level >= stripe.LevelDebug {
		Logger.Debug().Msgf(strings.TrimSpace(format), v...)
	}
}

func (s StripeLogger) Errorf(format string, v ...interface{}) {
	if s.Level >= stripe.LevelError {
		Logger.Error().Msgf(strings.TrimSpace(format), v...)
	}
}

func (s StripeLogger) Infof(format string, v ...interface{}) {
	if s.Level >= stripe.LevelInfo {
		Logger.Info().Msgf(strings.TrimSpace(format), v...)
	}
}

func (s StripeLogger) Warnf(format string, v ...interface{}) {
	if s.Level >= stripe.LevelWarn {
		Logger.Warn().Msgf(strings.TrimSpace(format), v...)
	}
}
