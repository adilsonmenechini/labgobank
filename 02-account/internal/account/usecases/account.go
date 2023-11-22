package usecases

import (
	"context"

	"github.com/adilsonmenechini/golabbank/internal/account/repositories"
	"github.com/adilsonmenechini/golabbank/internal/delivery/presenter"
	"github.com/adilsonmenechini/golabbank/pkg/utils"
)

type (
	Write interface {
		Create(ctx context.Context, req presenter.CreateAccountRequest) error
		Deposit(ctx context.Context, req presenter.OrderAccountRequest) error
		Withdraw(ctx context.Context, req presenter.OrderAccountRequest) error
		Transfer(ctx context.Context, req presenter.TransferAccountRequest) error
		Payment(ctx context.Context, req presenter.OrderAccountRequest) error
		PaymentLimit(ctx context.Context, req presenter.OrderAccountRequest) error
		Delete(ctx context.Context, req presenter.AccountNumberRequest) error
	}
	Reader interface {
		FindByAcoount(ctx context.Context, req presenter.AccountNumberRequest) (*presenter.AccountResponse, error)
		FindByCustomer(ctx context.Context, req presenter.AccountCustomerIDRequest) (*presenter.AccountResponse, error)
	}

	AccountUseCase interface {
		Write
		Reader
	}

	accountUseCase struct {
		logger *utils.Logger
		repo   repositories.AccountRepository
	}
)

// FindByCustomer implements AccountUseCase.
func (auc *accountUseCase) FindByCustomer(ctx context.Context, req presenter.AccountCustomerIDRequest) (*presenter.AccountResponse, error) {
	if err := utils.ValidateStruct(req); err != nil {
		auc.logger.Errorf("error validating request: %v", err)
		return &presenter.AccountResponse{}, err
	}

	acc, err := auc.repo.GetCustomerID(ctx, req.CustomerID)

	if err != nil {
		auc.logger.Errorf("error getting account: %v", err)
		return &presenter.AccountResponse{}, err
	}
	return &presenter.AccountResponse{
		AccountNumber: acc.AccountNumber,
		AccountType:   acc.AccountType,
		Balance:       acc.Balance,
		Limit:         acc.Limit,
		Name:          acc.Name,
	}, nil
}

// Transfer implements AccountUseCase.
func (auc *accountUseCase) Transfer(ctx context.Context, req presenter.TransferAccountRequest) error {

	if err := utils.ValidateStruct(req); err != nil {
		auc.logger.Errorf("error validating request: %v", err)
		return err
	}

	if err := auc.repo.Transfer(ctx, req.Amount, req.FromAccountNumber, req.ToAccountNumber); err != nil {
		auc.logger.Errorf("error depositing account: %v", err)
		return err
	}
	return nil
}

// Create implements AccountUseCase.
func (auc *accountUseCase) Create(ctx context.Context, req presenter.CreateAccountRequest) error {

	if err := utils.ValidateStruct(req); err != nil {
		auc.logger.Errorf("error validating request: %v", err)
		return err
	}

	if err := auc.repo.CreateAccount(ctx, req.CustomerID, req.Name, req.AccountType, req.Limit); err != nil {
		auc.logger.Errorf("error creating account: %v", err)
		return err
	}

	return nil
}

// Delete implements AccountUseCase.
func (auc *accountUseCase) Delete(ctx context.Context, req presenter.AccountNumberRequest) error {

	if err := utils.ValidateStruct(req); err != nil {
		auc.logger.Errorf("error validating request: %v", err)
		return err
	}

	if err := auc.repo.DeleteAccount(ctx, req.AccountNumber); err != nil {
		auc.logger.Errorf("error deleting account: %v", err)
		return err
	}
	return nil
}

// Deposit implements AccountUseCase.
func (auc *accountUseCase) Deposit(ctx context.Context, req presenter.OrderAccountRequest) error {

	if err := utils.ValidateStruct(req); err != nil {
		auc.logger.Errorf("error validating request: %v", err)
		return err
	}

	if err := auc.repo.Deposit(ctx, req.Amount, req.AccountNumber); err != nil {
		auc.logger.Errorf("error depositing account: %v", err)
		return err
	}
	return nil
}

// FindByAcoount implements AccountUseCase.
func (auc *accountUseCase) FindByAcoount(ctx context.Context, req presenter.AccountNumberRequest) (*presenter.AccountResponse, error) {

	if err := utils.ValidateStruct(req); err != nil {
		auc.logger.Errorf("error validating request: %v", err)
		return &presenter.AccountResponse{}, err
	}

	acc, err := auc.repo.GetAccountNumber(ctx, req.AccountNumber)

	if err != nil {
		auc.logger.Errorf("error getting account: %v", err)
		return &presenter.AccountResponse{}, err
	}
	return &presenter.AccountResponse{
		AccountNumber: acc.AccountNumber,
		AccountType:   acc.AccountType,
		Balance:       acc.Balance,
		Limit:         acc.Limit,
		Name:          acc.Name,
	}, nil
}

// Payment implements AccountUseCase.
func (auc *accountUseCase) Payment(ctx context.Context, req presenter.OrderAccountRequest) error {

	if err := utils.ValidateStruct(req); err != nil {
		auc.logger.Errorf("error validating request: %v", err)
		return err
	}

	if err := auc.repo.Payment(ctx, req.Amount, req.AccountNumber); err != nil {
		auc.logger.Errorf("error payment account: %v", err)
		return err
	}

	return nil
}

// PaymentLimit implements AccountUseCase.
func (auc *accountUseCase) PaymentLimit(ctx context.Context, req presenter.OrderAccountRequest) error {

	if err := utils.ValidateStruct(req); err != nil {
		auc.logger.Errorf("error validating request: %v", err)
		return err
	}

	if err := auc.repo.PaymentLimit(ctx, req.Amount, req.AccountNumber); err != nil {
		auc.logger.Errorf("error payment limit account: %v", err)
		return err
	}
	return nil
}

// Withdraw implements AccountUseCase.
func (auc *accountUseCase) Withdraw(ctx context.Context, req presenter.OrderAccountRequest) error {

	if err := utils.ValidateStruct(req); err != nil {
		auc.logger.Errorf("error validating request: %v", err)
		return err
	}

	if err := auc.repo.Withdraw(ctx, req.Amount, req.AccountNumber); err != nil {
		auc.logger.Errorf("error withdrawing account: %v", err)
		return err
	}
	return nil
}

func NewAccountUseCase(repo repositories.AccountRepository) AccountUseCase {
	return &accountUseCase{
		logger: utils.NewLogger("usecaseAccount"),
		repo:   repo,
	}
}
