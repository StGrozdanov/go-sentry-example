package handlers

import (
	sentry "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"sentry-example/internal/courses"
	"sentry-example/utils"
)

func GetCourses(ctx *gin.Context) {
	coursesResults, err := courses.GetAll(ctx)
	if err != nil {
		utils.
			GetLogger().
			WithFields(log.Fields{"error": err.Error()}).
			Error("Error on attempting to get all courses")

		sentry.GetHubFromContext(ctx).CaptureException(err)

		ctx.JSON(http.StatusInternalServerError, map[string]interface{}{})
		return
	}

	if len(coursesResults) > 0 {
		ctx.JSON(http.StatusOK, coursesResults)
		return
	}

	ctx.JSON(http.StatusOK, []courses.Course{})
}
