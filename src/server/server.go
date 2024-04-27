package server

import (
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"sentry-example/server/handlers"
	"sentry-example/server/middlewares"
	"sentry-example/utils"
)

func setupRouter() (router *gin.Engine) {
	gin.SetMode(gin.ReleaseMode)

	router = gin.New()

	router.Use(sentrygin.New(sentrygin.Options{
		Repanic: true,
	}))

	router.Use(middlewares.Logger(utils.GetLogger()), gin.Recovery(), middlewares.CORS())

	router.GET("/courses", handlers.GetCourses)
	router.GET("/frequent-questions", handlers.GetFrequentQuestions)

	return
}

// Run defines the router endpoints and starts the server
func Run() {
	router := setupRouter()

	err := router.Run()
	if err != nil {
		utils.GetLogger().WithFields(log.Fields{"error": err.Error()}).Error("Unable to start web server")
	}
	utils.GetLogger().Debug("Web server started ...")
}
