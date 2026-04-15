package controllers

import (
	"net/http"

	"github.com/TusharChoudharykp/Email-Campaign-Tool/services"
	"github.com/gin-gonic/gin"
)

func GetAllEmailLogs(c *gin.Context) {
	data, err := services.FetchAllEmailLogs()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to fetch email logs",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Email logs fetched successfully",
		"data":    data,
	})
}

func GetEmailLogsByCampaignID(c *gin.Context) {
	id := c.Param("id")

	data, err := services.FetchEmailLogsByCampaignID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to fetch campaign email logs",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Campaign email logs fetched successfully",
		"data":    data,
	})
}
