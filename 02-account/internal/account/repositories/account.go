package repositories

import (
	"context"
	"database/sql"

	"github.com/adilsonmenechini/golabbank/internal/domain"
	"github.com/adilsonmenechini/golabbank/pkg/utils"
)

type (
	AccountRepository interface {
		CreateAccount(ctx context.Context, customerID, name, accType string, inLimit float64) error
		DeleteAccount(ctx context.Context, accountNumber string) error
		GetAccountNumber(ctx context.Context, accountNumber string) (*domain.Account, error)
		GetCustomerID(ctx context.Context, customer string) (*domain.Account, error)
		Deposit(ctx context.Context, amount float64, accountNumber string) error
		Withdraw(ctx context.Context, amount float64, accountNumber string) error
		Transfer(ctx context.Context, amount float64, fromAccountNumber string, toAccountNumber string) error
		Payment(ctx context.Context, amount float64, accountNumber string) error
		PaymentLimit(ctx context.Context, amount float64, accountNumber string) error
	}

	accountRepository struct {
		logger *utils.Logger
		db     *sql.DB
	}
)

// GetCustomerID implements AccountRepository.
func (acr *accountRepository) GetCustomerID(ctx context.Context, customer string) (*domain.Account, error) {
	var i domain.Account

	i.Mu.Lock()
	defer i.Mu.Unlock()

	row := acr.db.QueryRowContext(ctx, getCustomerID, customer)

	err := row.Scan(
		&i.AccountNumber,
		&i.AccountType,
		&i.CustomerID,
		&i.Name,
		&i.Balance,
		&i.Limit,
		&i.Reversal,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	if err != nil {
		return &domain.Account{}, err
	}

	return &i, nil
}

// Transfer implements AccountRepository.
func (acr *accountRepository) Transfer(ctx context.Context, amount float64, fromAccountNumber string, toAccountNumber string) error {
	fromacc, err := acr.GetAccountNumber(ctx, fromAccountNumber)
	if err != nil {
		return err
	}
	toacc, err := acr.GetAccountNumber(ctx, toAccountNumber)
	if err != nil {
		return err
	}
	toAcc, err := fromacc.Transfer(toacc, amount)
	if err != nil {
		return err
	}

	_, err = acr.db.ExecContext(ctx, updatePayment,
		fromacc.AccountNumber,
		fromacc.Balance,
		fromacc.Limit,
		fromacc.UpdatedAt,
	)
	if err != nil {
		return err
	}

	_, err = acr.db.ExecContext(ctx, updatePayment,
		toAcc.AccountNumber,
		toAcc.Balance,
		toAcc.Limit,
		toAcc.UpdatedAt,
	)
	if err != nil {
		return err
	}
	return nil
}

// Deposit implements AccountRepository.
func (acr *accountRepository) Deposit(ctx context.Context, amount float64, accountNumber string) error {
	acc, err := acr.GetAccountNumber(ctx, accountNumber)
	if err != nil {
		return err
	}
	if err := acc.Deposit(amount); err != nil {
		return err
	}

	_, err = acr.db.ExecContext(ctx, updatePayment,
		acc.AccountNumber,
		acc.Balance,
		acc.Limit,
		acc.UpdatedAt,
	)
	if err != nil {
		return err
	}
	return nil
}

// Withdraw implements AccountRepository.
func (acr *accountRepository) Withdraw(ctx context.Context, amount float64, accountNumber string) error {
	acc, err := acr.GetAccountNumber(ctx, accountNumber)
	if err != nil {
		return err
	}
	if err := acc.Withdraw(amount); err != nil {
		return err
	}

	_, err = acr.db.ExecContext(ctx, updatePayment,
		acc.AccountNumber,
		acc.Balance,
		acc.Limit,
		acc.UpdatedAt,
	)
	if err != nil {
		return err
	}
	return nil
}

// CreateAccount implements AccountRepository.
func (acr *accountRepository) CreateAccount(ctx context.Context, customerID, name, accType string, inLimit float64) error {
	input := domain.Customer{
		ID:   customerID,
		Name: name,
	}
	newAcc := domain.NewAccount(&input, accType, inLimit)

	_, err := acr.db.ExecContext(ctx, createAccount,
		newAcc.AccountNumber,
		newAcc.AccountType,
		newAcc.CustomerID,
		newAcc.Name,
		newAcc.Balance,
		newAcc.Limit,
		newAcc.Reversal,
		newAcc.CreatedAt,
		newAcc.UpdatedAt,
	)
	if err != nil {
		return err
	}
	return nil
}

// DeleteAccount implements AccountRepository.
func (acr *accountRepository) DeleteAccount(ctx context.Context, accountNumber string) error {
	_, err := acr.db.ExecContext(ctx, deleteAccount, accountNumber)
	if err != nil {
		return err
	}
	return nil
}

// GetAccountNumber implements AccountRepository.
func (acr *accountRepository) GetAccountNumber(ctx context.Context, accountNumber string) (*domain.Account, error) {
	var i domain.Account

	i.Mu.Lock()
	defer i.Mu.Unlock()

	row := acr.db.QueryRowContext(ctx, getAccountNumber, accountNumber)

	err := row.Scan(
		&i.AccountNumber,
		&i.AccountType,
		&i.CustomerID,
		&i.Name,
		&i.Balance,
		&i.Limit,
		&i.Reversal,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	if err != nil {
		return &domain.Account{}, err
	}

	return &i, nil
}

// UpdateAccount implements AccountRepository.
func (acr *accountRepository) Payment(ctx context.Context, amount float64, accountNumber string) error {
	account, err := acr.GetAccountNumber(ctx, accountNumber)
	if err != nil {
		return err
	}
	if err := account.Payment(amount); err != nil {
		return err
	}

	_, err = acr.db.ExecContext(ctx, updatePayment,
		account.AccountNumber,
		account.Balance,
		account.Limit,
		account.UpdatedAt,
	)
	if err != nil {
		return err
	}
	return nil
}
func (acr *accountRepository) PaymentLimit(ctx context.Context, amount float64, accountNumber string) error {
	account, err := acr.GetAccountNumber(ctx, accountNumber)
	if err != nil {
		return err
	}
	if err := account.PaymentLimit(amount); err != nil {
		return err
	}

	_, err = acr.db.ExecContext(ctx, updatePayment,
		account.AccountNumber,
		account.Balance,
		account.Limit,
		account.UpdatedAt,
	)
	if err != nil {
		return err
	}
	return nil
}
func NewAccountRepository(DB *sql.DB) AccountRepository {
	return &accountRepository{
		logger: utils.NewLogger("AccountRepository"),
		db:     DB,
	}
}

const (
	createAccount    = `INSERT INTO Accounts (account_number, account_type, customer_id, name, balance, acc_limit, acc_reversal, created_at, updated_at) VALUES ( $1, $2, $3, $4, $5, $6, $7, $8, $9)`
	deleteAccount    = `DELETE FROM Accounts WHERE accountNumber = $1`
	getAccountNumber = `SELECT * FROM Accounts WHERE account_number = $1`
	getCustomerID    = `SELECT * FROM Accounts WHERE customer_id = $1`
	updatePayment    = `UPDATE Accounts set balance = $2,acc_limit = $3 ,updated_at = $4 WHERE account_number = $1`
	depositWithdraw  = `UPDATE Accounts set balance = $2, updated_at = $3 WHERE account_number = $1`
)
