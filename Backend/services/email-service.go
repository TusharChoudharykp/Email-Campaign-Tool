package services

import (
	models "github.com/TusharChoudharykp/Email-Campaign-Tool/models"
	"github.com/TusharChoudharykp/Email-Campaign-Tool/repositories"
)

func FetchAllEmailLogs() ([]models.EmailLog, error) {
	return repositories.GetAllEmailLogs()
}

func FetchEmailLogsByCampaignID(id string) ([]models.EmailLog, error) {
	return repositories.GetEmailLogsByCampaignID(id)
}

func FetchAdvancedEmailLogs() ([]map[string]interface{}, error) {
	return repositories.GetAdvancedEmailLogs()
}
