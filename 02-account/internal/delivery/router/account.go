package router

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/adilsonmenechini/golabbank/internal/account/repositories"
	"github.com/adilsonmenechini/golabbank/internal/account/usecases"
	"github.com/adilsonmenechini/golabbank/internal/delivery/handler"
	"github.com/adilsonmenechini/golabbank/pkg/utils"
	"github.com/gorilla/mux"
)

type AccountRouter struct {
	hdl    handler.AccountHandler
	logger *utils.Logger
}

func NewAccountRouter(hdlr handler.AccountHandler) *AccountRouter {
	return &AccountRouter{
		hdl:    hdlr,
		logger: utils.NewLogger("Router"),
	}
}

func (ra *AccountRouter) account() http.Handler {
	r := mux.NewRouter()
	c := r.PathPrefix("/api/account").Subrouter()

	a := c.PathPrefix("/v1").Subrouter()

	a.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	})

	a.HandleFunc("/create", ra.hdl.CreateAccountHandler).Methods("POST")
	a.HandleFunc("/deposit", ra.hdl.DepositHandler).Methods("POST")
	a.HandleFunc("/withdraw", ra.hdl.WithdrawHandler).Methods("POST")
	a.HandleFunc("/transfer", ra.hdl.TransferHandler).Methods("POST")
	a.HandleFunc("/payment", ra.hdl.PaymentHandler).Methods("POST")
	a.HandleFunc("/balance", ra.hdl.PaymentLimitHandler).Methods("POST")

	return r
}

func AccountImpl(db *sql.DB) http.Handler {
	repoC := repositories.NewAccountRepository(db)
	uscC := usecases.NewAccountUseCase(repoC)
	hdlC := handler.NewAccountHandler(uscC)
	rc := NewAccountRouter(hdlC).account()

	return rc
}
