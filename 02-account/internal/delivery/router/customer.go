package router

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/adilsonmenechini/golabbank/internal/customer/repositories"
	"github.com/adilsonmenechini/golabbank/internal/customer/usecases"
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

func (ra *CustomerRouter) customer() http.Handler {
	r := mux.NewRouter()
	c := r.PathPrefix("/api/customer").Subrouter()

	a := c.PathPrefix("/v1").Subrouter()

	a.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	})

	a.HandleFunc("/signin", ra.hdl.SigninHandler).Methods("GET")
	a.HandleFunc("/signup", ra.hdl.SignupHandler).Methods("POST")
	a.HandleFunc("/welcome", ra.hdl.AuthorizeCustomerHandler).Methods("GET")
	a.HandleFunc("/logout", ra.hdl.LogoutHandler).Methods("GET")

	return r
}

func CustomerImpl(db *sql.DB) http.Handler {
	repoC := repositories.NewCustomerRepository(db)
	uscC := usecases.NewCustomerUseCase(repoC)
	hdlC := handler.NewCustomerHandler(uscC)
	rc := NewCustomerRouter(hdlC).customer()

	return rc
}
