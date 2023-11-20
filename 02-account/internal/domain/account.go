package domain

import (
	"fmt"
	"sync"
	"time"

	"github.com/adilsonmenechini/golabbank/pkg/genrand"
)

type Account struct {
	AccountNumber string
	AccountType   string
	CustomerID    string
	Name          string
	Balance       float64
	Limit         float64
	Reversal      float64
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Mu            sync.Mutex
}

func checkAccountType(accType string, inLimit float64) float64 {
	var limit float64

	switch accType {
	case "bb":
		limit = 500
	case "itau":
		limit = 1000
	case "caixa":
		limit = 1000
	case "santander":
		limit = 200
	default:
		limit = 0
	}
	if inLimit > 0 {
		limit = inLimit
	}
	return limit

}
func NewAccount(cr *Customer, accType string, inLimit float64) *Account {
	limit := checkAccountType(accType, inLimit)
	acc := genrand.GenerateAcoount(accType)

	return &Account{
		AccountNumber: acc.CardNumber,
		AccountType:   string(acc.AccountType),
		CustomerID:    cr.ID,
		Name:          cr.Name,
		Balance:       0,
		Limit:         limit,
		Reversal:      limit,
		CreatedAt:     time.Now().UTC(),
		UpdatedAt:     time.Now().UTC(),
	}
}

func (a *Account) Deposit(amount float64) error {
	a.Mu.Lock()
	defer a.Mu.Unlock()
	if amount > 0 {
		a.Balance += amount
		a.UpdatedAt = time.Now().UTC()
	}
	return nil
}

func (a *Account) Withdraw(amount float64) error {
	a.Mu.Lock()
	defer a.Mu.Unlock()

	if amount >= 0 && amount > a.Balance {
		return fmt.Errorf("%d", ErrWithdrawalInsufficient)
	}

	a.Balance -= amount
	a.UpdatedAt = time.Now().UTC()
	return nil
}

func (a *Account) Payment(amount float64) error {
	a.Mu.Lock()
	defer a.Mu.Unlock()
	var amountToPay = amount
	if amountToPay > 0 {
		if a.Balance+a.Limit >= amountToPay {
			if a.Balance >= amountToPay {
				a.Balance -= amountToPay
				a.UpdatedAt = time.Now().UTC()
				return nil
			} else {
				amountToPay -= a.Balance
				a.Limit -= amountToPay
				a.Balance = 0
				a.UpdatedAt = time.Now().UTC()
				return nil
			}

		} else {
			return fmt.Errorf("%d", ErrPaymentInsufficient)
		}

	} else {
		return fmt.Errorf("%d", ErrPaymentInsufficient)
	}
}

func (a *Account) PaymentLimit(amount float64) error {
	a.Mu.Lock()
	defer a.Mu.Unlock()

	var amountToPay = amount

	if amountToPay > 0 {
		if a.Limit+amountToPay >= a.Reversal {
			a.Limit += amountToPay
			a.Limit -= a.Reversal
			a.Balance += a.Limit
			a.Limit = a.Reversal
			a.UpdatedAt = time.Now().UTC()
			return nil
		} else {
			a.Limit += amountToPay
			a.UpdatedAt = time.Now().UTC()
			return nil
		}

	} else {
		return fmt.Errorf("%d", ErrPaymentInsufficient)
	}
}

func (a *Account) GetBalance() float64 {
	a.Mu.Lock()
	defer a.Mu.Unlock()
	return a.Balance
}

func (a *Account) GetLimit() float64 {
	a.Mu.Lock()
	defer a.Mu.Unlock()
	return a.Limit
}

func (a *Account) Transfer(toAcc *Account, amount float64) (*Account, error) {
	a.Mu.Lock()
	defer a.Mu.Unlock()
	if amount > 0 {
		if a.Balance >= amount {
			a.Balance -= amount
			toAcc.Balance += amount
			a.UpdatedAt = time.Now().UTC()
			toAcc.UpdatedAt = time.Now().UTC()
			return toAcc, nil
		} else {
			return a, fmt.Errorf("%d", ErrTransferInsufficient)
		}
	}
	return toAcc, nil
}
