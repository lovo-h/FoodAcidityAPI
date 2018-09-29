package interfaces

import (
	"github.com/lovohh/FoodAcidityAPI/usecases"
	"net/http"
)

type HandlerDB interface {
	Query(string, []interface{}, func([][]byte)) error
}

type InteractorFood interface {
	OneFoodByNDB_No(string) ([]map[string]string, error)
	ManyLong_DescBySnippet([]string) ([]map[string]string, error)
}

type HandlerWebResponder interface {
	Success(http.ResponseWriter, usecases.M)
	Created(http.ResponseWriter, usecases.M)
	NoContent(http.ResponseWriter)

	Redirection(http.ResponseWriter)

	BadRequest(http.ResponseWriter, error)
	Unauthorized(http.ResponseWriter)
	Forbidden(http.ResponseWriter)
	NotFound(http.ResponseWriter)
	Gone(http.ResponseWriter)

	InternalServerError(http.ResponseWriter)
	ServiceUnavailable(http.ResponseWriter)
}

type HandlerSendGrid interface {
	AddRecipient(recipient string) error
	EmailUser(userEmail, subject, plainEmailMessage, htmlEmailMessage string) error
}
