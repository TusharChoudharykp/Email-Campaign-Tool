package services

import (
	"fmt"
	"sync"

	models "github.com/TusharChoudharykp/Email-Campaign-Tool/models"
	"github.com/TusharChoudharykp/Email-Campaign-Tool/repositories"
	"github.com/TusharChoudharykp/Email-Campaign-Tool/workers"
)

func AddCampaign(campaign models.Campaign) (*models.Campaign, error) {
	if campaign.Status == "" {
		campaign.Status = "draft"
	}

	id, err := repositories.CreateCampaign(campaign)
	if err != nil {
		return nil, err
	}

	return repositories.GetCampaignByID(fmt.Sprintf("%d", id))
}

func FetchCampaigns() ([]models.Campaign, error) {
	return repositories.GetAllCampaigns()
}

func FetchCampaignByID(id string) (*models.Campaign, error) {
	return repositories.GetCampaignByID(id)
}

func EditCampaign(id string, campaign models.Campaign) (bool, error) {
	if campaign.Status == "" {
		campaign.Status = "draft"
	}

	return repositories.UpdateCampaign(id, campaign)
}

func RemoveCampaign(id string) (bool, error) {
	return repositories.DeleteCampaign(id)
}

func SendCampaign(id string, segment string) (*models.SendSummary, error) {
	campaign, err := repositories.GetCampaignByID(id)
	if err != nil {
		return nil, err
	}

	var contacts []models.Contact

	if segment == "" {
		contacts, err = repositories.GetContactsForCampaign()
	} else {
		contacts, err = repositories.GetContactsBySegment(segment)
	}

	if err != nil {
		return nil, err
	}

	repositories.UpdateCampaignStatus(id, "sending")

	jobs := make(chan models.Contact)
	results := make(chan models.RecipientResult, len(contacts))

	var wg sync.WaitGroup
	workerCount := 5

	for i := 1; i <= workerCount; i++ {
		wg.Add(1)

		go workers.EmailWorker(
			i,
			&wg,
			jobs,
			results,
			*campaign,
		)
	}

	for _, contact := range contacts {
		jobs <- contact
	}

	close(jobs)
	wg.Wait()
	close(results)

	repositories.UpdateCampaignStatus(id, "sent")

	summary := &models.SendSummary{
		CampaignID:    campaign.ID,
		CampaignName:  campaign.Name,
		Status:        "sent",
		TotalContacts: len(contacts),
	}

	for result := range results {
		summary.Recipients = append(summary.Recipients, result)

		if result.Status == "sent" {
			summary.SentCount++
		} else {
			summary.FailedCount++
		}
	}

	return summary, nil
}
