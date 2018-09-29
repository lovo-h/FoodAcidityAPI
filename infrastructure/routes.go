package infrastructure

import (
	"github.com/lovohh/FoodAcidityAPI/interfaces"
	gmux "github.com/gorilla/mux"
	"net/http"
)

func GetRouterWithRoutes(webservice *interfaces.HandlerWebservice) *gmux.Router {
	router := gmux.NewRouter().StrictSlash(false)
	apiRouter := router.PathPrefix("/api").Subrouter()

	apiRouter.HandleFunc("/long_descs/{snippet:[a-zA-Z0-9_]{3,200}}", func(rw http.ResponseWriter, req *http.Request) {
		webservice.LongDescs(rw, req)
	}).Methods("GET")

	apiRouter.HandleFunc("/foods/{ndb_no:[0-9]{4,5}}", func(rw http.ResponseWriter, req *http.Request) {
		webservice.Food(rw, req)
	}).Methods("GET")

	apiRouter.HandleFunc("/emails/edible_list", func(rw http.ResponseWriter, req *http.Request) {
		webservice.EmailUserEdibleList(rw, req)
	}).Methods("POST")

	apiRouter.HandleFunc("/emails/recipients", func(rw http.ResponseWriter, req *http.Request) {
		webservice.AddRecipient(rw, req)
	}).Methods("POST")

	return router
}
