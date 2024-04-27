package main

import (
	log "github.com/sirupsen/logrus"
	"sentry-example/config"
	"sentry-example/database"
	"sentry-example/sentry"
	"sentry-example/server"
	"sentry-example/utils"
)

func init() {
	app, err := config.Init()
	if err != nil {
		utils.GetLogger().WithFields(log.Fields{"error": err.Error()}).Error("Error on config initialization")
		return
	}

	if app.AppEnv == "LOC" {
		utils.PrettyPrint(app)
	}

	database.Init(
		app.DBHosts,
		app.DBUsername,
		app.DBPassword,
		app.DBPort,
		app.DBName,
	)

	var (
		sentryTracesSampleRate   float64
		sentryProfilesSampleRate float64
	)

	if app.AppEnv != "LOC" {
		sentryTracesSampleRate = 0.5
		sentryProfilesSampleRate = 0.5
	} else {
		sentryTracesSampleRate = 1.0
		sentryProfilesSampleRate = 1.0
	}

	sentry.Init(
		app.SentryDSN,
		true,
		sentryTracesSampleRate,
		sentryProfilesSampleRate,
		app.AppEnv,
	)
}

func main() {
	server.Run()
}
