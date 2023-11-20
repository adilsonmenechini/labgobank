package router

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/adilsonmenechini/golabbank/internal/delivery/handler"
	"github.com/adilsonmenechini/golabbank/pkg/utils"
	"github.com/gorilla/mux"
)

type CustomerRouter struct {
	hdl    handler.CustomerHandler
	logger *utils.Logger
}

func NewCustomerRouter(hdlr handler.CustomerHandler) *CustomerRouter {
	return &CustomerRouter{
		hdl:    hdlr,
		logger: utils.NewLogger("Router"),
	}
}

func (ra *CustomerRouter) Router() {
	r := mux.NewRouter()
	s := r.PathPrefix("/api/v1").Subrouter()

	s.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	})

	a := s.PathPrefix("/customer").Subrouter()
	a.HandleFunc("/signin", ra.hdl.SigninHandler).Methods("GET")
	a.HandleFunc("/signup", ra.hdl.SignupHandler).Methods("POST")
	a.HandleFunc("/welcome", ra.hdl.AuthorizeCustomerHandler).Methods("GET")
	a.HandleFunc("/logout", ra.hdl.LogoutHandler).Methods("GET")

	ra.logger.Info("Servidor rodando na porta :8000")
	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
