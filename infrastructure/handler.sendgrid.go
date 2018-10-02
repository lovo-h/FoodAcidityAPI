package infrastructure

import (
	"encoding/json"
	"errors"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type HandlerSendGrid struct {
	apiKey string
}

func (this *HandlerSendGrid) AddRecipient(recipientStr string) error {

	host := "https://api.sendgrid.com"
	endpoint := "/v3/contactdb/recipients"
	request := sendgrid.GetRequest(this.apiKey, endpoint, host)
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

	var errorCount float64 = 0

	if respData["error_count"] != nil {
		errorCount = respData["error_count"].(float64)
	}

	if errorCount != 0 {
		return errors.New("Error: failed to create recipient on SendGrid")
	}

	return nil
}

func (this *HandlerSendGrid) EmailUser(userEmail, subject, plainEmailMessage, htmlEmailMessage string) error {
	from := mail.NewEmail("FoodAcidity API", "foodacidityapi@email.com")
	to := mail.NewEmail("Guest User", userEmail)
	message := mail.NewSingleEmail(from, subject, to, plainEmailMessage, htmlEmailMessage)
	client := sendgrid.NewSendClient(this.apiKey)
	_, sendErr := client.Send(message)

	if sendErr != nil {
		return errors.New("Failed to send email to " + userEmail)
	}

	return nil
}

func NewHandlerSendGrid(apiKey string) *HandlerSendGrid {
	sendgrid := new(HandlerSendGrid)
	sendgrid.apiKey = apiKey;

	return sendgrid;
}