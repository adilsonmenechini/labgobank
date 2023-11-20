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
		rs     *presenter.ResponsePresenter
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
		logger: utils.NewLogger("CustomerHandler"),
		us:     usa,
		rs:     presenter.NewResponsePresenter(),
		pa:     presenter.NewCustomerPresenter(),
	}
}

func (hc *customerHandler) SignupHandler(w http.ResponseWriter, r *http.Request) {
	var req presenter.SignupRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		hc.rs.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}
	_, err = hc.us.FindByEmail(r.Context(), req.Email)
	if err == nil {
		hc.rs.ResponseError(w, http.StatusUnauthorized, "email already exists")
		return
	}

	err = hc.us.Create(r.Context(), req)
	if err != nil {
		hc.rs.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	hc.rs.ResponseSuccess(w, http.StatusCreated, "Customer created successfully")
}

func (hc *customerHandler) SigninHandler(w http.ResponseWriter, r *http.Request) {
	var req presenter.SigninRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		hc.rs.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	input, err := hc.us.FindByEmail(r.Context(), req.Email)
	if err != nil {
		hc.rs.ResponseError(w, http.StatusUnauthorized, "email or password incorrect")
		return
	}
	check := utils.CheckPasswordHash(req.Password, input.Password)

	if !check {
		hc.rs.ResponseError(w, http.StatusUnauthorized, "email or password incorrect")
		return
	}

	jwtg, err := utils.GenerateJWT(input.ID, input.Email)
	if err != nil {
		hc.rs.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SetTokenAsCookie(w, jwtg)

	hc.rs.ResponseSuccess(w, http.StatusOK, "login successful")

}

func (hc *customerHandler) AuthorizeCustomerHandler(w http.ResponseWriter, r *http.Request) {

	tk, err := utils.GetTokenFromCookie(r)
	if err != nil {
		hc.rs.ResponseErrorToken(w, http.StatusUnauthorized, "invalid token")
		return
	}

	utils.SetTokenAsCookie(w, tk.Email)
	hc.rs.ResponseSuccess(w, http.StatusOK, "welcome "+tk.Email)

}

func (hc *customerHandler) LogoutHandler(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Expires: time.Now(),
	})
}
