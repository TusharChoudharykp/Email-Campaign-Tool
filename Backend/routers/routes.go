package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(app *gin.Engine) {
	app.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Backend running successfully",
		})
	})
}