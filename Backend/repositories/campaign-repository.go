package repositories

import (
	"github.com/TusharChoudharykp/Email-Campaign-Tool/config"
	models "github.com/TusharChoudharykp/Email-Campaign-Tool/models"
)

func CreateCampaign(campaign models.Campaign) (int64, error) {
	query := `
	INSERT INTO campaigns(name, subject, template, status)
	VALUES (?, ?, ?, ?)
	`

	result, err := config.DB.Exec(
		query,
		campaign.Name,
		campaign.Subject,
		campaign.Template,
		campaign.Status,
	)

	if err != nil {
		return 0, err
	}

	insertedID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return insertedID, nil
}

func GetAllCampaigns() ([]models.Campaign, error) {
	query := `
	SELECT id, name, subject, template, status, created_at
	FROM campaigns
	`

	rows, err := config.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var campaigns []models.Campaign

	for rows.Next() {
		var campaign models.Campaign

		err := rows.Scan(
			&campaign.ID,
			&campaign.Name,
			&campaign.Subject,
			&campaign.Template,
			&campaign.Status,
			&campaign.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		campaigns = append(campaigns, campaign)
	}

	return campaigns, nil
}

func GetCampaignByID(id string) (*models.Campaign, error) {
	query := `
	SELECT id, name, subject, template, status, created_at
	FROM campaigns
	WHERE id = ?
	`

	var campaign models.Campaign

	err := config.DB.QueryRow(query, id).Scan(
		&campaign.ID,
		&campaign.Name,
		&campaign.Subject,
		&campaign.Template,
		&campaign.Status,
		&campaign.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &campaign, nil
}

func UpdateCampaign(id string, campaign models.Campaign) (bool, error) {
	query := `
	UPDATE campaigns
	SET name = ?, subject = ?, template = ?, status = ?
	WHERE id = ?
	`

	result, err := config.DB.Exec(
		query,
		campaign.Name,
		campaign.Subject,
		campaign.Template,
		campaign.Status,
		id,
	)

	if err != nil {
		return false, err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return false, err
	}

	if rows == 0 {
		return false, nil
	}

	return true, nil
}

func DeleteCampaign(id string) (bool, error) {
	query := "DELETE FROM campaigns WHERE id = ?"

	result, err := config.DB.Exec(query, id)
	if err != nil {
		return false, err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return false, err
	}

	if rows == 0 {
		return false, nil
	}

	return true, nil
}
