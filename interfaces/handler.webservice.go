package interfaces

import (
	"encoding/json"
	"fmt"
	gmux "github.com/gorilla/mux"
	"github.com/lovohh/FoodAcidityAPI/usecases"
	"html"
	"net/http"
	"strings"
)

type HandlerWebservice struct {
	InteractorFood InteractorFood

	HandlerWebResponder HandlerWebResponder
	HandlerSendGrid     HandlerSendGrid
}

func (handler *HandlerWebservice) LongDescs(rw http.ResponseWriter, req *http.Request) {
	longDescSnippets, _ := gmux.Vars(req)["snippet"]

	longDescs, getErr := handler.InteractorFood.ManyLong_DescBySnippet(strings.Split(longDescSnippets, "_"))

	if getErr != nil {
		handler.HandlerWebResponder.InternalServerError(rw)
		return
	}

	handler.HandlerWebResponder.Success(rw, usecases.M{"data": longDescs})
}

func (handler *HandlerWebservice) Food(rw http.ResponseWriter, req *http.Request) {
	ndbNo, _ := gmux.Vars(req)["ndb_no"]

	food, getErr := handler.InteractorFood.OneFoodByNDB_No(ndbNo)

	if getErr != nil {
		handler.HandlerWebResponder.InternalServerError(rw)
		return
	}

	handler.HandlerWebResponder.Success(rw, usecases.M{"data": food})
}

func (handler *HandlerWebservice) AddRecipient(rw http.ResponseWriter, req *http.Request) {
	var data map[string]string
	decoder := json.NewDecoder(req.Body)
	decodeErr := decoder.Decode(&data)

	if decodeErr != nil {
		handler.HandlerWebResponder.InternalServerError(rw)
		return
	}

	for _, key := range []string{"first_name", "last_name", "email"} {
		if _, hasKey := data[key]; !hasKey {
			handler.HandlerWebResponder.InternalServerError(rw)
			return
		}
	}

	jsonData, marshalErr := json.Marshal(data)

	if marshalErr != nil {
		handler.HandlerWebResponder.InternalServerError(rw)
		return
	}

	// recipientStr: [string]
	// format: "[{'first_name': 'name', 'last_name': 'last', 'email': 'example@email.com'}]"
	recipientsStr := "[" + string(jsonData) + "]"
	addErr := handler.HandlerSendGrid.AddRecipient(recipientsStr)

	if addErr != nil {
		handler.HandlerWebResponder.InternalServerError(rw)
		return
	}

	handler.HandlerWebResponder.NoContent(rw)
}

func (handler *HandlerWebservice) EmailUserEdibleList(rw http.ResponseWriter, req *http.Request) {
	var data map[string][]string
	decoder := json.NewDecoder(req.Body)
	decodeErr := decoder.Decode(&data)

	if decodeErr != nil {
		handler.HandlerWebResponder.InternalServerError(rw)
		return
	}

	foodList := data["edible_foods"]

	var emailMsg string
	var emailHTMLMsg string

	for foodIdx, food := range foodList {
		safeFood := html.EscapeString(strings.Replace(food, "%", "pct", -1))
		emailMsg += fmt.Sprintf("%d %s\n", foodIdx+1, safeFood)
		emailHTMLMsg += fmt.Sprintf("<li>%s</li>", safeFood)
	}

	emailMsg = fmt.Sprintf("Personalized Food List:\n%s", emailMsg)
	emailHTMLMsg = fmt.Sprintf(
		"<span style='font-weight: bold'>Personalized Food List:</span><br/><ul>%s</ul>", emailHTMLMsg)

	userEmail := data["email"][0]
	subject := "Personalized Food List To Decrease Mortality Rate"

	/* 	args:
			userEmail -> user@email.com
			subject -> whatever subject the API needs it to be
			plainEmailMessage -> list of food items user can eat
			htmlEmailMessage -> same as plainEmailMessage but with HTML formatting
	*/
	sendErr := handler.HandlerSendGrid.EmailUser(userEmail, subject, emailMsg, emailHTMLMsg)

	if sendErr != nil {
		handler.HandlerWebResponder.InternalServerError(rw)
		return
	}

	handler.HandlerWebResponder.NoContent(rw)
}
