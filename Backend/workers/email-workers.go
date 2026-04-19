package workers

import (
	"bytes"
	"fmt"
	"net/smtp"
	"sync"
	"text/template"

	models "github.com/TusharChoudharykp/Email-Campaign-Tool/models"
	"github.com/TusharChoudharykp/Email-Campaign-Tool/repositories"
)

func EmailWorker(
	id int,
	wg *sync.WaitGroup,
	jobs <-chan models.Contact,
	results chan<- models.RecipientResult,
	campaign models.Campaign,
) {
	defer wg.Done()

	for contact := range jobs {
		err := sendMail(contact, campaign)

		if err != nil {
			repositories.CreateEmailLog(
				campaign.ID,
				contact.ID,
				"failed",
				err.Error(),
			)

			results <- models.RecipientResult{
				Name:   contact.Name,
				Email:  contact.Email,
				Status: "failed",
				Error:  err.Error(),
			}

			continue
		}

		repositories.CreateEmailLog(
			campaign.ID,
			contact.ID,
			"sent",
			"",
		)

		results <- models.RecipientResult{
			Name:   contact.Name,
			Email:  contact.Email,
			Status: "sent",
		}

		fmt.Printf("Worker %d sent to %s\n", id, contact.Email)
	}
}

func sendMail(contact models.Contact, campaign models.Campaign) error {
	tpl, err := template.New("email").Parse(campaign.Template)
	if err != nil {
		return err
	}

	var body bytes.Buffer

	err = tpl.Execute(&body, contact)
	if err != nil {
		return err
	}

	message := fmt.Sprintf(
		"To: %s\r\nSubject: %s\r\n\r\n%s",
		contact.Email,
		campaign.Subject,
		body.String(),
	)

	return smtp.SendMail(
		"localhost:1025",
		nil,
		"admin@example.com",
		[]string{contact.Email},
		[]byte(message),
	)
}
