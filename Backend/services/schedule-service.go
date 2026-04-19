package services

import (
	"errors"
	"fmt"
	"time"

	models "github.com/TusharChoudharykp/Email-Campaign-Tool/models"
	"github.com/TusharChoudharykp/Email-Campaign-Tool/repositories"
)

func AddSchedule(data models.ScheduledCampaign) (*models.ScheduledCampaign, error) {
	layout := "2006-01-02 15:04:05"

	loc := time.Now().Location()

	sendTime, err := time.ParseInLocation(layout, data.SendAt, loc)
	if err != nil {
		return nil, errors.New("invalid datetime format. use YYYY-MM-DD HH:MM:SS")
	}

	now := time.Now()

	if !sendTime.After(now) {
		return nil, errors.New("scheduled time must be in the future")
	}

	id, err := repositories.CreateSchedule(data)
	if err != nil {
		return nil, err
	}

	return repositories.GetScheduleByID(fmt.Sprintf("%d", id))
}

func ProcessScheduledCampaigns() {
	items, err := repositories.GetPendingSchedules()
	if err != nil {
		fmt.Println("Scheduler fetch error:", err)
		return
	}

	for _, item := range items {
		repositories.UpdateScheduleStatus(item.ID, "processing")

		_, err := SendCampaign(
			fmt.Sprintf("%d", item.CampaignID),
			item.Segment,
		)

		if err != nil {
			repositories.UpdateScheduleStatus(item.ID, "failed")
			continue
		}

		repositories.UpdateScheduleStatus(item.ID, "completed")
	}
}
