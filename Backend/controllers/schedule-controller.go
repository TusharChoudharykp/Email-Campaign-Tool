package controllers

import (
	"net/http"

	models "github.com/TusharChoudharykp/Email-Campaign-Tool/models"
	"github.com/TusharChoudharykp/Email-Campaign-Tool/services"
	"github.com/gin-gonic/gin"
)

func CreateSchedule(c *gin.Context) {
	var data models.ScheduledCampaign

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body",
			"error":   err.Error(),
		})
		return
	}

	result, err := services.AddSchedule(data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Campaign scheduled successfully",
		"data":    result,
	})
}
