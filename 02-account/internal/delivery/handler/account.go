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
	}
)

// CreateAccountHandler implements AccountHandler.
func (hac *accountHandler) CreateAccountHandler(w http.ResponseWriter, r *http.Request) {
	tk, err := utils.GetTokenAuthorization(r)
	if err != nil {
		hac.rs.ResponseErrorToken(w, http.StatusUnauthorized, err.Error())
		return
	}

	req := presenter.CreateAccountRequest{
		CustomerID: tk.ID,
		Name:       tk.Name,
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
	_, err := utils.GetTokenAuthorization(r)
	if err != nil {
		hac.rs.ResponseErrorToken(w, http.StatusUnauthorized, err.Error())
		return
	}

	var req = presenter.OrderAccountRequest{}

	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		hac.rs.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	err = hac.us.Deposit(r.Context(), req)
	if err != nil {
		hac.rs.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	hac.rs.ResponseSuccess(w, http.StatusOK, "Deposit successfully")
}

// PaymentHandler implements AccountHandler.
func (hac *accountHandler) PaymentHandler(w http.ResponseWriter, r *http.Request) {
	_, err := utils.GetTokenAuthorization(r)
	if err != nil {
		hac.rs.ResponseErrorToken(w, http.StatusUnauthorized, err.Error())
		return
	}

	var req = presenter.OrderAccountRequest{}

	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		hac.rs.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	err = hac.us.Payment(r.Context(), req)
	if err != nil {
		hac.rs.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	hac.rs.ResponseSuccess(w, http.StatusOK, "Payment successfully")
}

// PaymentLimitHandler implements AccountHandler.
func (hac *accountHandler) PaymentLimitHandler(w http.ResponseWriter, r *http.Request) {
	_, err := utils.GetTokenAuthorization(r)
	if err != nil {
		hac.rs.ResponseErrorToken(w, http.StatusUnauthorized, err.Error())
		return
	}

	var req = presenter.OrderAccountRequest{}

	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		hac.rs.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	err = hac.us.PaymentLimit(r.Context(), req)
	if err != nil {
		hac.rs.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	hac.rs.ResponseSuccess(w, http.StatusOK, "Pauyment limit successfully")
}

// TransferHandler implements AccountHandler.
func (hac *accountHandler) TransferHandler(w http.ResponseWriter, r *http.Request) {
	_, err := utils.GetTokenAuthorization(r)
	if err != nil {
		hac.rs.ResponseErrorToken(w, http.StatusUnauthorized, err.Error())
		return
	}

	var req = presenter.TransferAccountRequest{}

	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		hac.rs.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	err = hac.us.Transfer(r.Context(), req)
	if err != nil {
		hac.rs.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	hac.rs.ResponseSuccess(w, http.StatusOK, "Transfer successfully")
}

// WithdrawHandler implements AccountHandler.
func (hac *accountHandler) WithdrawHandler(w http.ResponseWriter, r *http.Request) {
	_, err := utils.GetTokenAuthorization(r)
	if err != nil {
		hac.rs.ResponseErrorToken(w, http.StatusUnauthorized, err.Error())
		return
	}

	var req = presenter.OrderAccountRequest{}

	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		hac.rs.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	err = hac.us.Withdraw(r.Context(), req)
	if err != nil {
		hac.rs.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	hac.rs.ResponseSuccess(w, http.StatusOK, "Withdraw successfully")
}

func NewAccountHandler(usa usecases.AccountUseCase) AccountHandler {
	return &accountHandler{
		logger: utils.NewLogger("AccountHandler"),
		us:     usa,
		rs:     presenter.NewResponsePresenter(),
		pa:     presenter.NewAccountPresenter(),
	}
}
