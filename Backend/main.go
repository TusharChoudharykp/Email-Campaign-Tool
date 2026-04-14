package main

import (
	"bytes"
	"sync"
	"text/template"
)

type Recipient struct {
	Name  string
	Email string
}

func main() {
	recipientChannel := make(chan Recipient)

	go func() {
		loadRecipient("recipients.csv", recipientChannel)
	}()

	var wg sync.WaitGroup
	workercount := 5 

	for i := 1; i <= workercount; i++ {
		wg.Add(1)
		go emailWorker(i, recipientChannel, &wg)
	}
	
	wg.Wait()
	//time.Sleep(3 * time.Second)
}

func executeTemplate(r Recipient) (string, error) {
	t, err := template.ParseFiles("email.tmpl")
	if err != nil {
		return "", err
	}

	var tpl bytes.Buffer

	err = t.Execute(&tpl, r)
	if err != nil {
		return "", err
	}

	return tpl.String(), nil
}




// Imporvements:
// 1. Store the contact list into DB.
// 2. Add segments
// 3. Multiple campaign
// 4. Handle email fail
// 5. Customize email template
// 6. Schedule campaign
// 7. Build a frontend