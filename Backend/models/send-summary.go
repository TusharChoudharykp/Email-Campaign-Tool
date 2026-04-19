package models

type RecipientResult struct {
	Name   string `json:"name"`
	Email  string `json:"email"`
	Status string `json:"status"`
	Error  string `json:"error,omitempty"`
}

type SendSummary struct {
	CampaignID    int               `json:"campaign_id"`
	CampaignName  string            `json:"campaign_name"`
	Status        string            `json:"status"`
	TotalContacts int               `json:"total_contacts"`
	SentCount     int               `json:"sent_count"`
	FailedCount   int               `json:"failed_count"`
	Recipients    []RecipientResult `json:"recipients"`
}
