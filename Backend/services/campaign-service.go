package services

import (
	"fmt"

	models "github.com/TusharChoudharykp/Email-Campaign-Tool/models"
	"github.com/TusharChoudharykp/Email-Campaign-Tool/repositories"
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
