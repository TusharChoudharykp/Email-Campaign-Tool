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
	// Contacts routes
	app.POST("/contacts", controllers.CreateContact)
	app.GET("/contacts", controllers.GetContacts)
	app.GET("/contacts/:id", controllers.GetContactByID)
	app.PUT("/contacts/:id", controllers.UpdateContact)
	app.DELETE("/contacts/:id", controllers.DeleteContact)

	//Campaign routes
	app.POST("/campaigns", controllers.CreateCampaign)
	app.GET("/campaigns", controllers.GetCampaigns)
	app.GET("/campaigns/:id", controllers.GetCampaignByID)
	app.PUT("/campaigns/:id", controllers.UpdateCampaign)
	app.DELETE("/campaigns/:id", controllers.DeleteCampaign)
}
