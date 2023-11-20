package presenter

import (
	"encoding/json"
	"net/http"

	"github.com/adilsonmenechini/golabbank/pkg/utils"
)

type ResponsePresenter struct {
	logger *utils.Logger
}

func NewResponsePresenter() *ResponsePresenter {
	return &ResponsePresenter{
		logger: utils.NewLogger("presenter"),
	}
}

func (pa *ResponsePresenter) ResponseSuccess(w http.ResponseWriter, statusCode int, res string) {

	pa.logger.Infof("statusCode: %d, message: %s", statusCode, res)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":     "success",
		"statusCode": statusCode,
		"message":    res,
	})
}

func (pa *ResponsePresenter) ResponseError(w http.ResponseWriter, statusCode int, res string) {
	//pa.logger.Errorf("statusCode: %d, message: %s", statusCode, res)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":     "error",
		"statusCode": statusCode,
		"message":    res,
	})
}

func (pa *ResponsePresenter) ResponseErrorToken(w http.ResponseWriter, statusCode int, res string) {
	pa.logger.Errorf("statusCode: %d, message: %s", statusCode, res)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":     "error",
		"statusCode": statusCode,
		"message":    res,
	})
}
