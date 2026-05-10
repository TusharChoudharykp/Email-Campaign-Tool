package main

import (
	"time"

	"github.com/TusharChoudharykp/Email-Campaign-Tool/config"
	routes "github.com/TusharChoudharykp/Email-Campaign-Tool/routers"
	"github.com/TusharChoudharykp/Email-Campaign-Tool/services"
	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
)

func main() {
	config.LoadEnv()

	config.ConnectDB()

	app := gin.Default()

	// CORS Configuration
	app.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:5173",
		},
		AllowMethods: []string{
			"GET",
			"POST",
			"PUT",
			"DELETE",
			"OPTIONS",
		},
		AllowHeaders: []string{
			"Origin",
			"Content-Type",
			"Authorization",
		},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	//routes.RegisterRoutes(app)
	routes.RegisterRoutes(app)

	// Start Scheduler
	scheduler := cron.New()

	_, err := scheduler.AddFunc("@every 1m", func() {
		services.ProcessScheduledCampaigns()
	})

	if err != nil {
		panic(err)
	}

	scheduler.Start()

	port := config.GetEnv("PORT")

	app.Run(":" + port)
}
