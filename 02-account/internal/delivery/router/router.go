package router

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func Router(db *sql.DB) {

	rcustomer := CustomerImpl(db)
	r := mux.NewRouter()
	//r.PathPrefix("/api/account/v1").Handler(accountRouter)
	r.PathPrefix("/api/customer/v1").Handler(rcustomer)

	// Crie o servidor HTTP usando o roteador principal
	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
