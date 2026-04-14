package controllers

import (
	"net/http"

	models "github.com/TusharChoudharykp/Email-Campaign-Tool/models"
	"github.com/TusharChoudharykp/Email-Campaign-Tool/services"

	"github.com/gin-gonic/gin"
)

func CreateContact(c *gin.Context) {
	var contact models.Contact

	err := c.ShouldBindJSON(&contact)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body",
			"error":   err.Error(),
		})
		return
	}

	data, err := services.AddContact(contact)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to create contact",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Contact created successfully",
		"data":    data,
	})
}

func GetContacts(c *gin.Context) {
	contacts, err := services.FetchContacts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to fetch contacts",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Contacts fetched successfully",
		"data":    contacts,
	})
}

func GetContactByID(c *gin.Context) {
	id := c.Param("id")

	contact, err := services.FetchContactByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Contact id does not exist in database",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Contact fetched successfully",
		"data":    contact,
	})
}

func UpdateContact(c *gin.Context) {
	id := c.Param("id")

	var contact models.Contact

	err := c.ShouldBindJSON(&contact)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body",
			"error":   err.Error(),
		})
		return
	}

	updated, err := services.EditContact(id, contact)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to update contact",
			"error":   err.Error(),
		})
		return
	}

	if !updated {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Contact id does not exist in database",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Contact updated successfully",
	})
}

func DeleteContact(c *gin.Context) {
	id := c.Param("id")

	deleted, err := services.RemoveContact(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to delete contact",
			"error":   err.Error(),
		})
		return
	}

	if !deleted {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Contact id does not exist in database",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Contact deleted successfully",
	})
}