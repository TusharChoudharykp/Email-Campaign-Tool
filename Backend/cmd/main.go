package main

import (
	"github.com/TusharChoudharykp/Email-Campaign-Tool/config"
	routes "github.com/TusharChoudharykp/Email-Campaign-Tool/routers"
	"github.com/TusharChoudharykp/Email-Campaign-Tool/services"

	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
)

func main() {
	config.LoadEnv()

	config.ConnectDB()

	app := gin.Default()

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
