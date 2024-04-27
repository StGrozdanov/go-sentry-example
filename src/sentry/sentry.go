package sentry

import (
	"fmt"
	"github.com/getsentry/sentry-go"
	log "github.com/sirupsen/logrus"
	"sentry-example/utils"
	"sync"
)

type sentryInstance struct {
	dsn                string
	enableTracing      bool
	tracesSampleRate   float64
	profilesSampleRate float64
	appEnv             string
}

var instance *sentryInstance

// Init initialises sentry connection
func Init(dsn string, enableTracing bool, tracesSampleRate, profilesSampleRate float64, appEnv string) {
	var syncOnce sync.Once
	if instance == nil {
		syncOnce.Do(
			func() {
				instance = &sentryInstance{
					dsn:                dsn,
					enableTracing:      enableTracing,
					tracesSampleRate:   tracesSampleRate,
					appEnv:             appEnv,
					profilesSampleRate: profilesSampleRate,
				}
				connect()
			},
		)
	}
}

func connect() {
	if err := sentry.Init(sentry.ClientOptions{
		Dsn:                instance.dsn,
		EnableTracing:      instance.enableTracing,
		TracesSampleRate:   instance.tracesSampleRate,
		Environment:        instance.appEnv,
		IgnoreTransactions: []string{`^(HEAD|GET) /healths$`, `^GET /posts-socket$`, `^OPTIONS.*$`},
		Debug:              true,
		ProfilesSampleRate: instance.profilesSampleRate,
		Release:            fmt.Sprintf("sentry-go-example-api-%s", instance.appEnv),
	}); err != nil {
		utils.
			GetLogger().
			WithFields(log.Fields{"error": err.Error(), "DSN": instance.dsn}).
			Error("Error on connection attempt to the Sentry instance")
	}
}
