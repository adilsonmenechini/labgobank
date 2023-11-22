package presenter

import (
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

type CustomersToken struct {
	Token string `json:"token" valid:"notnull,email"`
}

type CustomerPresenter struct {
	logger *utils.Logger
}

func NewCustomerPresenter() *CustomerPresenter {
	return &CustomerPresenter{
		logger: utils.NewLogger("presenter"),
	}
}
