package controllers

import (
	"net/http"

	"github.com/TusharChoudharykp/Email-Campaign-Tool/services"
	"github.com/gin-gonic/gin"
)

func GetDashboardStats(c *gin.Context) {

	data, err := services.FetchDashboardStats()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to fetch dashboard stats",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Dashboard stats fetched successfully",
		"data":    data,
	})
}
