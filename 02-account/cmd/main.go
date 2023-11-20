package main

import (
	"github.com/adilsonmenechini/golabbank/config"
	"github.com/adilsonmenechini/golabbank/internal/delivery/router"
	"github.com/adilsonmenechini/golabbank/pkg/database"
)

func main() {

	config.ParseEnvVariables()
	dbcon := database.ConnectPSQL()
	router.Router(dbcon)

}
