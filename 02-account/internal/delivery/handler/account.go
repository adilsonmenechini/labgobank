package handler

import (
	"encoding/json"
	"net/http"

	"github.com/adilsonmenechini/golabbank/internal/account/usecases"
	"github.com/adilsonmenechini/golabbank/internal/delivery/presenter"
	"github.com/adilsonmenechini/golabbank/pkg/utils"
)

type (
	accountHandler struct {
		logger *utils.Logger
		pa     *presenter.AccountPresenter
		rs     *presenter.ResponsePresenter
		us     usecases.AccountUseCase
	}
	AccountHandler interface {
		DepositHandler(w http.ResponseWriter, r *http.Request)
		WithdrawHandler(w http.ResponseWriter, r *http.Request)
		TransferHandler(w http.ResponseWriter, r *http.Request)
		CreateAccountHandler(w http.ResponseWriter, r *http.Request)
		PaymentHandler(w http.ResponseWriter, r *http.Request)
		PaymentLimitHandler(w http.ResponseWriter, r *http.Request)
		AuthorizeAccountHandler(w http.ResponseWriter, r *http.Request)
	}
)

// AuthorizeAccountHandler implements AccountHandler.
func (hac *accountHandler) AuthorizeAccountHandler(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

// CreateAccountHandler implements AccountHandler.
func (hac *accountHandler) CreateAccountHandler(w http.ResponseWriter, r *http.Request) {

	tk, err := utils.GetTokenFromCookie(r)
	if err != nil {
		hac.rs.ResponseErrorToken(w, http.StatusUnauthorized, "invalid token")
		return
	}
	req := presenter.CreateAccountRequest{
		CustomerID: tk.ID,
		Name:       tk.Email,
	}

	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		hac.rs.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	err = hac.us.Create(r.Context(), req)
	if err != nil {
		hac.rs.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	hac.rs.ResponseSuccess(w, http.StatusCreated, "Account created successfully")
}

// DepositHandler implements AccountHandler.
func (hac *accountHandler) DepositHandler(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

// PaymentHandler implements AccountHandler.
func (hac *accountHandler) PaymentHandler(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

// PaymentLimitHandler implements AccountHandler.
func (hac *accountHandler) PaymentLimitHandler(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

// TransferHandler implements AccountHandler.
func (hac *accountHandler) TransferHandler(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

// WithdrawHandler implements AccountHandler.
func (hac *accountHandler) WithdrawHandler(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

func NewAccountHandler(usa usecases.AccountUseCase) AccountHandler {
	return &accountHandler{
		logger: utils.NewLogger("AccountHandler"),
		us:     usa,
		rs:     presenter.NewResponsePresenter(),
		pa:     presenter.NewAccountPresenter(),
	}
}
