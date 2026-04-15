package models

type EmailLog struct {
	ID           int    `json:"id"`
	CampaignID   int    `json:"campaign_id"`
	ContactID    int    `json:"contact_id"`
	Status       string `json:"status"`
	ErrorMessage string `json:"error_message"`
	SentAt       string `json:"sent_at"`
}
