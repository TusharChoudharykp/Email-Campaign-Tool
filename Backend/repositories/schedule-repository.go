package repositories

import (
	"github.com/TusharChoudharykp/Email-Campaign-Tool/config"
	models "github.com/TusharChoudharykp/Email-Campaign-Tool/models"
)

func CreateSchedule(data models.ScheduledCampaign) (int64, error) {
	query := `
	INSERT INTO scheduled_campaigns(campaign_id, segment, send_at, status)
	VALUES (?, ?, ?, ?)
	`

	result, err := config.DB.Exec(
		query,
		data.CampaignID,
		data.Segment,
		data.SendAt,
		"pending",
	)

	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func GetPendingSchedules() ([]models.ScheduledCampaign, error) {
	query := `
	SELECT id, campaign_id, segment, send_at, status, created_at
	FROM scheduled_campaigns
	WHERE status = 'pending' AND send_at <= NOW()
	`

	rows, err := config.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var schedules []models.ScheduledCampaign

	for rows.Next() {
		var item models.ScheduledCampaign

		err := rows.Scan(
			&item.ID,
			&item.CampaignID,
			&item.Segment,
			&item.SendAt,
			&item.Status,
			&item.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		schedules = append(schedules, item)
	}

	return schedules, nil
}

func UpdateScheduleStatus(id int, status string) error {
	query := `
	UPDATE scheduled_campaigns
	SET status = ?
	WHERE id = ?
	`

	_, err := config.DB.Exec(query, status, id)
	return err
}

func GetScheduleByID(id string) (*models.ScheduledCampaign, error) {
	query := `
	SELECT id, campaign_id, segment, send_at, status, created_at
	FROM scheduled_campaigns
	WHERE id = ?
	`

	var item models.ScheduledCampaign

	err := config.DB.QueryRow(query, id).Scan(
		&item.ID,
		&item.CampaignID,
		&item.Segment,
		&item.SendAt,
		&item.Status,
		&item.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &item, nil
}
