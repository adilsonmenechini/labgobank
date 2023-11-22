package presenter

import (
	"time"

	"github.com/adilsonmenechini/golabbank/pkg/utils"
)

type CreateAccountRequest struct {
	CustomerID  string  `json:"customer_id" valid:"notnull"`
	Name        string  `json:"name" valid:"notnull"`
	AccountType string  `json:"account_type" valid:"notnull"`
	Limit       float64 `json:"limit" valid:"notnull"`
}

type AccountNumberRequest struct {
	AccountNumber string `json:"account_number" valid:"notnull" `
}

type AccountCustomerIDRequest struct {
	CustomerID string `json:"customer_id" valid:"notnull" `
}

type OrderAccountRequest struct {
	AccountNumber string  `json:"account_number" valid:"notnull" `
	Amount        float64 `json:"amount" valid:"notnull"`
}

type TransferAccountRequest struct {
	FromAccountNumber string  `json:"from_account" valid:"notnull" `
	ToAccountNumber   string  `json:"to_account" valid:"notnull" `
	Amount            float64 `json:"amount" valid:"notnull"`
}

type CreateAccountResponse struct {
	AccountNumber string    `json:"account_number"`
	AccountType   string    `json:"account_type"`
	AccountID     string    `json:"Account_id"`
	Name          string    `json:"name"`
	Balance       float64   `json:"balance"`
	Limit         float64   `json:"limit"`
	CreatedAt     time.Time `json:"created_at"`
}

type AccountResponse struct {
	AccountNumber string  `json:"account_number"`
	AccountType   string  `json:"account_type"`
	Name          string  `json:"name"`
	Balance       float64 `json:"balance"`
	Limit         float64 `json:"limit"`
}

type AccountPresenter struct {
	logger *utils.Logger
}

func NewAccountPresenter() *AccountPresenter {
	return &AccountPresenter{
		logger: utils.NewLogger("presenter"),
	}
}
