package frequent_questions

import (
	"github.com/gin-gonic/gin"
	"sentry-example/database"
)

// GetAll retrieves all FAQ from the database
func GetAll(ctx *gin.Context) (results []FAQ, err error) {
	err = database.GetMultipleRecords(
		ctx,
		&results,
		`SELECT question, answer FROM frequent_questions;`,
	)
	return
}
