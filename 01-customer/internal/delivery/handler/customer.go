package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/adilsonmenechini/golabbank/internal/customer/usecases"
	"github.com/adilsonmenechini/golabbank/internal/delivery/presenter"

	"github.com/adilsonmenechini/golabbank/pkg/utils"
)

type (
	customerHandler struct {
		logger *utils.Logger
		pa     *presenter.CustomerPresenter
		us     usecases.CustomerUseCase
	}
	CustomerHandler interface {
		SignupHandler(w http.ResponseWriter, r *http.Request)
		SigninHandler(w http.ResponseWriter, r *http.Request)
		AuthorizeCustomerHandler(w http.ResponseWriter, r *http.Request)
		LogoutHandler(w http.ResponseWriter, r *http.Request)
	}
)

func NewCustomerHandler(usa usecases.CustomerUseCase) CustomerHandler {
	return &customerHandler{
		logger: utils.NewLogger("Handler"),
		us:     usa,
		pa:     presenter.NewCustomerPresenter(),
	}
}

func (h *customerHandler) SignupHandler(w http.ResponseWriter, r *http.Request) {
	var req presenter.SignupRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		h.pa.CustomerError(w, http.StatusBadRequest, err.Error())
		return
	}
	_, err = h.us.FindByEmail(r.Context(), req.Email)
	if err == nil {
		h.pa.CustomerError(w, http.StatusUnauthorized, "email already exists")
		return
	}

	err = h.us.Create(r.Context(), req)
	if err != nil {
		h.pa.CustomerError(w, http.StatusBadRequest, err.Error())
		return
	}

	h.pa.CustomerSuccess(w, http.StatusCreated, "Customer created successfully")
}

func (h *customerHandler) SigninHandler(w http.ResponseWriter, r *http.Request) {
	var req presenter.SigninRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		h.pa.CustomerError(w, http.StatusBadRequest, err.Error())
		return
	}

	input, err := h.us.FindByEmail(r.Context(), req.Email)
	if err != nil {
		h.pa.CustomerError(w, http.StatusUnauthorized, "email or password incorrect")
		return
	}
	check := utils.CheckPasswordHash(req.Password, input.Password)

	if !check {
		h.pa.CustomerError(w, http.StatusUnauthorized, "email or password incorrect")
		return
	}

	jwtg, err := utils.GenerateJWT(input.ID, input.Email)
	if err != nil {
		h.pa.CustomerError(w, http.StatusInternalServerError, err.Error())
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:  "token",
		Value: jwtg,
	})

	h.pa.CustomerSuccess(w, http.StatusOK, "login successful")

}

func (h *customerHandler) AuthorizeCustomerHandler(w http.ResponseWriter, r *http.Request) {

	tk, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	acc, err := utils.ValidateToken(tk.Value)
	if err != nil {
		h.pa.CustomerErrorToken(w, http.StatusUnauthorized, "invalid token")
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:  "token",
		Value: tk.Value,
	})
	h.pa.CustomerSuccess(w, http.StatusOK, "welcome "+acc)

}

func (h *customerHandler) LogoutHandler(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Expires: time.Now(),
	})
}
