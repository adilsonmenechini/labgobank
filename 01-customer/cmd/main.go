package main

import (
	"github.com/adilsonmenechini/golabbank/config"
	"github.com/adilsonmenechini/golabbank/internal/customer/repositories"
	"github.com/adilsonmenechini/golabbank/internal/customer/usecases"
	"github.com/adilsonmenechini/golabbank/internal/delivery/handler"
	"github.com/adilsonmenechini/golabbank/internal/delivery/router"
	"github.com/adilsonmenechini/golabbank/pkg/database"
)

func main() {

	config.ParseEnvVariables()
	dbcon := database.ConnectPSQL()
	repo := repositories.NewCustomerRepository(dbcon)
	usc := usecases.NewCustomerUseCase(repo)
	hdl := handler.NewCustomerHandler(usc)
	router.NewCustomerRouter(hdl).Router()

}
