package repositories

import (
	"fmt"

	"github.com/TusharChoudharykp/Email-Campaign-Tool/config"
	models "github.com/TusharChoudharykp/Email-Campaign-Tool/models"
)

func CreateContact(contact models.Contact) (*models.Contact, error) {
	query := "INSERT INTO contacts(name, email, segment) VALUES (?, ?, ?)"

	result, err := config.DB.Exec(query, contact.Name, contact.Email, contact.Segment)

	if err != nil {
		return nil, err
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	return GetContactByID(fmt.Sprintf("%d", lastID))
}

func GetAllContacts() ([]models.Contact, error) {
	query := "SELECT id, name, email, segment, created_at FROM contacts"

	rows, err := config.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var contacts []models.Contact

	for rows.Next() {
		var contact models.Contact

		err := rows.Scan(
			&contact.ID,
			&contact.Name,
			&contact.Email,
			&contact.Segment,
			&contact.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		contacts = append(contacts, contact)
	}

	return contacts, nil
}

func GetContactByID(id string) (*models.Contact, error) {
	query := "SELECT id, name, email, segment, created_at FROM contacts WHERE id = ?"

	var contact models.Contact

	err := config.DB.QueryRow(query, id).Scan(
		&contact.ID,
		&contact.Name,
		&contact.Email,
		&contact.Segment,
		&contact.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &contact, nil
}

func UpdateContact(id string, contact models.Contact) (bool, error) {
	query := "UPDATE contacts SET name = ?, email = ?, segment = ? WHERE id = ?"

	result, err := config.DB.Exec(
		query,
		contact.Name,
		contact.Email,
		contact.Segment,
		id,
	)

	if err != nil {
		return false, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}

	if rowsAffected == 0 {
		return false, nil
	}

	return true, nil
}

func DeleteContact(id string) (bool, error) {
	query := "DELETE FROM contacts WHERE id = ?"

	result, err := config.DB.Exec(query, id)
	if err != nil {
		return false, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}

	if rowsAffected == 0 {
		return false, nil
	}

	return true, nil
}

func GetContactsForCampaign() ([]models.Contact, error) {
	query := `
	SELECT id, name, email, segment, created_at
	FROM contacts
	`

	rows, err := config.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var contacts []models.Contact

	for rows.Next() {
		var contact models.Contact

		err := rows.Scan(
			&contact.ID,
			&contact.Name,
			&contact.Email,
			&contact.Segment,
			&contact.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		contacts = append(contacts, contact)
	}

	return contacts, nil
}
