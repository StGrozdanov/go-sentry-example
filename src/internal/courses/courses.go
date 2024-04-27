package courses

import (
	"github.com/gin-gonic/gin"
	"sentry-example/database"
)

// GetAll returns all courses from the database
func GetAll(ctx *gin.Context) (results []Course, err error) {
	err = database.GetMultipleRecords(
		ctx,
		&results,
		`SELECT title, sub_title, image_url FROM courses;`,
	)
	return
}
