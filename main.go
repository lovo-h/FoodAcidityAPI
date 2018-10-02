package main

import (
	"github.com/lovohh/FoodAcidityAPI/infrastructure"
	"github.com/lovohh/FoodAcidityAPI/interfaces"
	"github.com/lovohh/FoodAcidityAPI/usecases"
	"net/http"
	"os"
)

func main() {
	logger := new(infrastructure.HandlerLogger)
	handlerCockrach := infrastructure.NewCDBHandler()
	repoFood := new(interfaces.RepoFood)
	repoFood.HandlerDB = handlerCockrach
	interactorFood := new(usecases.InteractorFood)
	interactorFood.RepoFood = repoFood
	webresponderjson := new(infrastructure.WebResponderJSON)

	sendgridAPIKey := os.Getenv("SENDGRID_API_KEY")
	sendgrid := infrastructure.NewHandlerSendGrid(sendgridAPIKey)

	handlerWebservice := &interfaces.HandlerWebservice{
		InteractorFood:      interactorFood,
		HandlerWebResponder: webresponderjson,
		HandlerSendGrid:     sendgrid,
	}

	router := infrastructure.GetRouterWithRoutes(handlerWebservice)

	server := &http.Server{
		Addr:    "0.0.0.0:3000",
		Handler: router,
	}

	if serverErr := server.ListenAndServe(); serverErr != nil {
		logger.Log("Server failed to boot: " + serverErr.Error())
	}
}
