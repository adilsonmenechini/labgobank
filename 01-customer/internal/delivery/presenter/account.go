package presenter

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/adilsonmenechini/golabbank/pkg/utils"
)

type SignupRequest struct {
	Name     string `json:"name" valid:"notnull"`
	Email    string `json:"email" valid:"notnull,email"`
	Password string `json:"password" valid:"notnull"`
}

type SigninRequest struct {
	Email    string `json:"email" valid:"notnull,email"`
	Password string `json:"password" valid:"notnull"`
}

type CustomerUpdate struct {
	Email    string `json:"email" valid:"notnull,email"`
	Password string `json:"password" valid:"notnull"`
}

type CustomerResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

type CustomerPresenter struct {
	logger *utils.Logger
}

func NewCustomerPresenter() *CustomerPresenter {
	return &CustomerPresenter{
		logger: utils.NewLogger("presenter"),
	}
}

func (pa *CustomerPresenter) CustomerSuccess(w http.ResponseWriter, statusCode int, res string) {

	pa.logger.Infof("statusCode: %d, message: %s", statusCode, res)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":     "success",
		"statusCode": statusCode,
		"message":    res,
	})
}

func (pa *CustomerPresenter) CustomerError(w http.ResponseWriter, statusCode int, res string) {
	//pa.logger.Errorf("statusCode: %d, message: %s", statusCode, res)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":     "error",
		"statusCode": statusCode,
		"message":    res,
	})
}

func (pa *CustomerPresenter) CustomerErrorToken(w http.ResponseWriter, statusCode int, res string) {
	pa.logger.Errorf("statusCode: %d, message: %s", statusCode, res)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":     "error",
		"statusCode": statusCode,
		"message":    res,
	})
}
