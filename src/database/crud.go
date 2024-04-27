package database

import (
	"context"
	"fmt"
	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
	"time"
)

// GetMultipleRecords selects multiple records from the database
func GetMultipleRecords(ginCtx *gin.Context, destination interface{}, query string) (err error) {
	var ctx, cancel = context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	// Results in unlabeled transaction on the performance tab
	span := generateSentryDBOperationSpan(ginCtx, query, "db.query")
	defer span.Finish()

	// Results in panic .. nil pointer because of the transaction from ginCtx proves to be nil
	//span := generateSentryDBOperationSpanFromTransaction(ginCtx, query, "db.query")
	//defer span.Finish()

	// Results in panic ... for whatever reason
	//span := generateSentryDBOperationSpanFromNewTransaction(ginCtx, query, "db.query")
	//defer span.Finish()

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

func generateSentryDBOperationSpanFromTransaction(ctx *gin.Context, sqlQuery, operation string) (span *sentry.Span) {
	transaction := sentry.TransactionFromContext(ctx)

	span = sentry.StartSpan(
		transaction.Context(),
		operation,
		sentry.WithDescription(sqlQuery),
	)

	span.SetData("db.system", "postgresql")
	return
}

func generateSentryDBOperationSpanFromNewTransaction(ctx *gin.Context, sqlQuery, operation string) (span *sentry.Span) {
	options := []sentry.SpanOption{
		sentry.WithOpName(operation),
		sentry.WithDescription(sqlQuery),
	}

	transaction := sentry.StartTransaction(
		ctx,
		fmt.Sprintf("%s %s", ctx.Request.Method, ctx.Request.URL),
		options...,
	)

	transaction.SetData("db.system", "postgresql")
	return
}
