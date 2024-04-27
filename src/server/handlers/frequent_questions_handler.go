package handlers

import (
	sentry "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"sentry-example/internal/frequent_questions"
	"sentry-example/utils"
)

func GetFrequentQuestions(ctx *gin.Context) {
	questionsResults, err := frequent_questions.GetAll(ctx)
	if err != nil {
		utils.
			GetLogger().
			WithFields(log.Fields{"error": err.Error()}).
			Error("Error on attempting to get all FAQ")

		sentry.GetHubFromContext(ctx).CaptureException(err)

		ctx.JSON(http.StatusInternalServerError, map[string]interface{}{})
		return
	}

	if len(questionsResults) > 0 {
		ctx.JSON(http.StatusOK, questionsResults)
		return
	}

	ctx.JSON(http.StatusOK, []frequent_questions.FAQ{})
}
