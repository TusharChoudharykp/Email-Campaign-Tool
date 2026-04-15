package controllers

import (
	"net/http"

	models "github.com/TusharChoudharykp/Email-Campaign-Tool/models"
	"github.com/TusharChoudharykp/Email-Campaign-Tool/services"

	"github.com/gin-gonic/gin"
)

func CreateCampaign(c *gin.Context) {
	var campaign models.Campaign

	if err := c.ShouldBindJSON(&campaign); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body",
			"error":   err.Error(),
		})
		return
	}

	data, err := services.AddCampaign(campaign)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to create campaign",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Campaign created successfully",
		"data":    data,
	})
}

func GetCampaigns(c *gin.Context) {
	data, err := services.FetchCampaigns()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to fetch campaigns",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Campaigns fetched successfully",
		"data":    data,
	})
}

func GetCampaignByID(c *gin.Context) {
	id := c.Param("id")

	data, err := services.FetchCampaignByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Campaign id does not exist in database",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Campaign fetched successfully",
		"data":    data,
	})
}

func UpdateCampaign(c *gin.Context) {
	id := c.Param("id")

	var campaign models.Campaign

	if err := c.ShouldBindJSON(&campaign); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body",
		})
		return
	}

	updated, err := services.EditCampaign(id, campaign)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to update campaign",
			"error":   err.Error(),
		})
		return
	}

	if !updated {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Campaign id does not exist in database",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Campaign updated successfully",
	})
}

func DeleteCampaign(c *gin.Context) {
	id := c.Param("id")

	deleted, err := services.RemoveCampaign(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to delete campaign",
			"error":   err.Error(),
		})
		return
	}

	if !deleted {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Campaign id does not exist in database",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Campaign deleted successfully",
	})
}
