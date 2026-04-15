package repositories

import (
	"github.com/TusharChoudharykp/Email-Campaign-Tool/config"
	models "github.com/TusharChoudharykp/Email-Campaign-Tool/models"
)

func CreateEmailLog(
	campaignID int,
	contactID int,
	status string,
	errorMessage string,
) error {

	query := `
	INSERT INTO email_logs(campaign_id, contact_id, status, error_message)
	VALUES (?, ?, ?, ?)
	`

	_, err := config.DB.Exec(
		query,
		campaignID,
		contactID,
		status,
		errorMessage,
	)

	return err
}

func GetAllEmailLogs() ([]models.EmailLog, error) {
	query := `
	SELECT id, campaign_id, contact_id, status, error_message, sent_at
	FROM email_logs
	ORDER BY id DESC
	`

	rows, err := config.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var logs []models.EmailLog

	for rows.Next() {
		var log models.EmailLog

		err := rows.Scan(
			&log.ID,
			&log.CampaignID,
			&log.ContactID,
			&log.Status,
			&log.ErrorMessage,
			&log.SentAt,
		)
		if err != nil {
			return nil, err
		}

		logs = append(logs, log)
	}

	return logs, nil
}

func GetEmailLogsByCampaignID(id string) ([]models.EmailLog, error) {
	query := `
	SELECT id, campaign_id, contact_id, status, error_message, sent_at
	FROM email_logs
	WHERE campaign_id = ?
	ORDER BY id DESC
	`

	rows, err := config.DB.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var logs []models.EmailLog

	for rows.Next() {
		var log models.EmailLog

		err := rows.Scan(
			&log.ID,
			&log.CampaignID,
			&log.ContactID,
			&log.Status,
			&log.ErrorMessage,
			&log.SentAt,
		)
		if err != nil {
			return nil, err
		}

		logs = append(logs, log)
	}

	return logs, nil
}
