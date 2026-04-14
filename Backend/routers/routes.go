package routes

import (
	"github.com/TusharChoudharykp/Email-Campaign-Tool/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(app *gin.Engine) {
	app.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Backend running successfully",
		})
	})

	app.POST("/contacts", controllers.CreateContact)
	app.GET("/contacts", controllers.GetContacts)
	app.GET("/contacts/:id", controllers.GetContactByID)
	app.PUT("/contacts/:id", controllers.UpdateContact)
	app.DELETE("/contacts/:id", controllers.DeleteContact)
}