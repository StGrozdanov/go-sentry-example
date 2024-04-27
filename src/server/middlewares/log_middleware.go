package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"math"
	"net/http"
	"os"
	"time"
)

var (
	timeFormat = "02/Jan/2006:15:04:05 -0700"
	skip       map[string]struct{}
)

// Logger logging middleware that tracks the request and response data and execution times. Accepts fields
// that can be ignored.
func Logger(logger logrus.FieldLogger, notLogged ...string) gin.HandlerFunc {
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknown"
	}
	populateSkippedFields(notLogged...)

	return func(ctx *gin.Context) {
		var (
			path  = ctx.Request.URL.Path
			start = time.Now()
		)

		ctx.Next()

		var (
			stop       = time.Since(start)
			latency    = int(math.Ceil(float64(stop.Nanoseconds()) / 1000000.0))
			statusCode = ctx.Writer.Status()
			clientIP   = ctx.ClientIP()
			referer    = ctx.Request.Referer()
			dataLength = ctx.Writer.Size()
		)

		if dataLength < 0 {
			dataLength = 0
		}

		if _, ok := skip[path]; ok {
			return
		}

		entry := initialiseLogEntry(
			ctx, logger, hostname, statusCode, latency, clientIP, path, referer, dataLength,
		)

		if len(ctx.Errors) > 0 {
			entry.Error(ctx.Errors.ByType(gin.ErrorTypePrivate).String())
		} else {
			msg := generateMessage(ctx, clientIP, hostname, path, statusCode, dataLength, referer, latency)
			if statusCode >= http.StatusInternalServerError {
				entry.Error(msg)
			} else if statusCode >= http.StatusBadRequest {
				entry.Warn(msg)
			} else {
				entry.Info(msg)
			}
		}
	}
}

func initialiseLogEntry(
	ctx *gin.Context,
	logger logrus.FieldLogger,
	hostname string,
	statusCode int,
	latency int,
	clientIP string,
	path string,
	referer string,
	dataLength int,
) *logrus.Entry {
	return logger.WithFields(logrus.Fields{
		"hostname":       hostname,
		"statusCode":     statusCode,
		"responseTimeMs": latency,
		"clientIP":       clientIP,
		"method":         ctx.Request.Method,
		"path":           path,
		"referer":        referer,
		"dataLength":     dataLength,
	})
}

func generateMessage(
	ctx *gin.Context,
	clientIP string,
	hostname string,
	path string,
	statusCode int,
	dataLength int,
	referer string,
	latency int,
) string {
	return fmt.Sprintf(
		"%s - %s [%s] \"%s %s\" %d %d \"%s\" (%dms)",
		clientIP,
		hostname,
		time.Now().Format(timeFormat),
		ctx.Request.Method,
		path,
		statusCode,
		dataLength,
		referer,
		latency,
	)
}

func populateSkippedFields(notLogged ...string) {
	if length := len(notLogged); length > 0 {
		skip = make(map[string]struct{}, length)
		for _, p := range notLogged {
			skip[p] = struct{}{}
		}
	}
}
