package infrastructure

import (
	"encoding/json"
	"errors"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

const (
	sendgridKey = ""
)

type HandlerSendGrid struct {
}

func (this *HandlerSendGrid) AddRecipient(recipientStr string) error {
	host := "https://api.sendgrid.com"
	endpoint := "/v3/contactdb/recipients"
	request := sendgrid.GetRequest(sendgridKey, endpoint, host)
	request.Method = "POST"
	request.Body = []byte(recipientStr)

	response, responseErr := sendgrid.API(request)

	if responseErr != nil {
		return errors.New("Error: failed to send message that would add recipient to SendGrid")
	}

	var respData map[string]interface{}
	jsonErr := json.Unmarshal([]byte(response.Body), &respData)

	if jsonErr != nil {
		return errors.New("AddRecipient: failed to unmarshal string to JSON")
	}

	errorCount := respData["error_count"].(float64)

	if errorCount != 0 {
		return errors.New("Error: failed to create recipient on SendGrid")
	}

	return nil
}

func (this *HandlerSendGrid) EmailUser(userEmail, subject, plainEmailMessage, htmlEmailMessage string) error {
	from := mail.NewEmail("FoodAcidity API", "foodacidityapi@email.com")
	to := mail.NewEmail("Guest User", userEmail)
	message := mail.NewSingleEmail(from, subject, to, plainEmailMessage, htmlEmailMessage)
	client := sendgrid.NewSendClient(sendgridKey)
	_, sendErr := client.Send(message)

	if sendErr != nil {
		return errors.New("Failed to send email to " + userEmail)
	}

	return nil
}
