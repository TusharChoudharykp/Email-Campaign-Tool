package main

import (
	"github.com/TusharChoudharykp/Email-Campaign-Tool/config"
	routes "github.com/TusharChoudharykp/Email-Campaign-Tool/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()

	config.ConnectDB()

	app := gin.Default()

	//routes.RegisterRoutes(app)
	routes.RegisterRoutes(app)

	port := config.GetEnv("PORT")

	app.Run(":" + port)
}