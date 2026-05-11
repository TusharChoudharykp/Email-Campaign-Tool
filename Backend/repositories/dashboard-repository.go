package repositories

import "github.com/TusharChoudharykp/Email-Campaign-Tool/config"

func GetDashboardStats() (map[string]int, error) {

	stats := map[string]int{}

	var totalContacts int
	var totalCampaigns int
	var sentEmails int
	var failedEmails int

	// Total Contacts
	err := config.DB.QueryRow(
		"SELECT COUNT(*) FROM contacts",
	).Scan(&totalContacts)

	if err != nil {
		return nil, err
	}

	// Total Campaigns
	err = config.DB.QueryRow(
		"SELECT COUNT(*) FROM campaigns",
	).Scan(&totalCampaigns)

	if err != nil {
		return nil, err
	}

	// Sent Emails
	err = config.DB.QueryRow(
		"SELECT COUNT(*) FROM email_logs WHERE status='sent'",
	).Scan(&sentEmails)

	if err != nil {
		return nil, err
	}

	// Failed Emails
	err = config.DB.QueryRow(
		"SELECT COUNT(*) FROM email_logs WHERE status='failed'",
	).Scan(&failedEmails)

	if err != nil {
		return nil, err
	}

	stats["total_contacts"] = totalContacts
	stats["total_campaigns"] = totalCampaigns
	stats["sent_emails"] = sentEmails
	stats["failed_emails"] = failedEmails

	return stats, nil
}
