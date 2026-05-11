package routes

import (
	"github.com/TusharChoudharykp/Email-Campaign-Tool/controllers"
	"github.com/TusharChoudharykp/Email-Campaign-Tool/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(app *gin.Engine) {

	// Health Route
	app.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Backend running successfully",
		})
	})

	// Public Routes
	app.POST("/auth/register", controllers.Register)
	app.POST("/auth/login", controllers.Login)

	// Protected Routes
	protected := app.Group("/")
	protected.Use(middlewares.AuthMiddleware())

	// Contacts Routes
	protected.POST("/contacts", controllers.CreateContact)
	protected.GET("/contacts", controllers.GetContacts)
	protected.GET("/contacts/:id", controllers.GetContactByID)
	protected.GET("/dashboard/stats", controllers.GetDashboardStats)
	protected.GET("/schedules", controllers.GetSchedules)
	protected.PUT("/contacts/:id", controllers.UpdateContact)
	protected.DELETE("/contacts/:id", controllers.DeleteContact)

	// Campaign Routes
	protected.POST("/campaigns", controllers.CreateCampaign)
	protected.POST("/campaigns/:id/send", controllers.SendCampaign)
	protected.GET("/campaigns", controllers.GetCampaigns)
	protected.GET("/campaigns/:id", controllers.GetCampaignByID)
	protected.PUT("/campaigns/:id", controllers.UpdateCampaign)
	protected.DELETE("/campaigns/:id", controllers.DeleteCampaign)

	// Email Logs Routes
	protected.GET("/email-logs", controllers.GetAllEmailLogs)
	protected.GET("/campaigns/:id/logs", controllers.GetEmailLogsByCampaignID)
	protected.GET("/email-logs/advanced", controllers.GetAdvancedEmailLogs)

	// Schedule Routes
	protected.POST("/schedules", controllers.CreateSchedule)
}
