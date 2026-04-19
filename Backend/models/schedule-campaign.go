package models

type ScheduledCampaign struct {
	ID         int    `json:"id"`
	CampaignID int    `json:"campaign_id"`
	Segment    string `json:"segment"`
	SendAt     string `json:"send_at"`
	Status     string `json:"status"`
	CreatedAt  string `json:"created_at"`
}
