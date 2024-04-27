package database

import (
	"context"
	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
	"time"
)

// GetMultipleRecords selects multiple records from the database
func GetMultipleRecords(ginCtx *gin.Context, destination interface{}, query string) (err error) {
	var ctx, cancel = context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	span := generateSentryDBOperationSpan(ginCtx, query, "db.query")
	defer span.Finish()

	return instance.DB.Unsafe().SelectContext(ctx, destination, query)
}

func generateSentryDBOperationSpan(ctx *gin.Context, sqlQuery, operation string) (span *sentry.Span) {
	span = sentry.StartSpan(
		ctx,
		operation,
		sentry.WithDescription(sqlQuery),
	)

	span.SetData("db.system", "postgresql")
	return
}
