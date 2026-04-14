package services

import (
	"strings"

	models "github.com/TusharChoudharykp/Email-Campaign-Tool/models"
	"github.com/TusharChoudharykp/Email-Campaign-Tool/repositories"
)

func AddContact(contact models.Contact) (*models.Contact, error) {
	if contact.Segment == "" {
		contact.Segment = "general"
	}
	contact.Email = strings.ToLower(strings.TrimSpace(contact.Email))
	contact.Segment = strings.ToLower(strings.TrimSpace(contact.Segment))

	return repositories.CreateContact(contact)
}

func FetchContacts() ([]models.Contact, error) {
	return repositories.GetAllContacts()
}

func FetchContactByID(id string) (*models.Contact, error) {
	return repositories.GetContactByID(id)
}

func EditContact(id string, contact models.Contact) (bool, error) {
	if contact.Segment == "" {
		contact.Segment = "general"
	}

	return repositories.UpdateContact(id, contact)
}

func RemoveContact(id string) (bool, error) {
	return repositories.DeleteContact(id)
}