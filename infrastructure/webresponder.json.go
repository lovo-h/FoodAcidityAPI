package infrastructure

import (
	"net/http"
	"encoding/json"
	"github.com/lovohh/FoodAcidityAPI/usecases"
)

type WebResponderJSON struct {
	WebResponder
}

func (responder *WebResponderJSON) jsonResponse(rw http.ResponseWriter, respMap usecases.M, statusCode int) {
	resp, respErr := json.Marshal(respMap)

	if respErr != nil {
		responder.InternalServerError(rw)
		return
	}

	rw.WriteHeader(statusCode)
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(resp)
}

func (responder *WebResponderJSON) Success(rw http.ResponseWriter, respMap usecases.M) {
	responder.jsonResponse(rw, respMap, http.StatusOK)
}

func (responder *WebResponderJSON) Created(rw http.ResponseWriter, respMap usecases.M) {
	responder.jsonResponse(rw, respMap, http.StatusCreated)
}

func (responder *WebResponderJSON) BadRequest(rw http.ResponseWriter, err error) {
	responder.jsonResponse(rw, usecases.M{"error": err.Error()}, http.StatusBadRequest)
}
