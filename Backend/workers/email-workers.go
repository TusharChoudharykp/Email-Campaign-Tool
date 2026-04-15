package workers

import (
	"bytes"
	"fmt"
	"log"
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
	campaign models.Campaign,
) {
	defer wg.Done()

	for contact := range jobs {
		err := sendMail(contact, campaign)

		if err != nil {
			log.Println("Failed:", contact.Email, err)

			repositories.CreateEmailLog(
				campaign.ID,
				contact.ID,
				"failed",
				err.Error(),
			)

			continue
		}

		repositories.CreateEmailLog(
			campaign.ID,
			contact.ID,
			"sent",
			"",
		)

		fmt.Printf("Worker %d sent to %s\n", id, contact.Email)
	}
}

func sendMail(
	contact models.Contact,
	campaign models.Campaign,
) error {

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

	smtpHost := "localhost"
	smtpPort := "1025"

	return smtp.SendMail(
		smtpHost+":"+smtpPort,
		nil,
		"admin@example.com",
		[]string{contact.Email},
		[]byte(message),
	)
}
