package usecases

import (
	"context"

	"github.com/adilsonmenechini/golabbank/internal/customer/repositories"
	"github.com/adilsonmenechini/golabbank/internal/delivery/presenter"
	"github.com/adilsonmenechini/golabbank/internal/domain"
	"github.com/adilsonmenechini/golabbank/pkg/utils"
)

type (
	Write interface {
		Create(ctx context.Context, req presenter.SignupRequest) error
		Delete(ctx context.Context, id string) error
		UpdatePassword(ctx context.Context, req presenter.CustomerUpdate) error
	}
	Reader interface {
		FindByEmail(ctx context.Context, email string) (domain.Customer, error)
		FindByID(ctx context.Context, id string) (domain.Customer, error)
	}

	CustomerUseCase interface {
		Write
		Reader
	}

	customerUseCase struct {
		logger *utils.Logger
		repo   repositories.CustomerRepository
	}
)

func NewCustomerUseCase(repo repositories.CustomerRepository) CustomerUseCase {
	return &customerUseCase{
		logger: utils.NewLogger("usecase"),
		repo:   repo,
	}
}

func (u *customerUseCase) Create(ctx context.Context, req presenter.SignupRequest) error {

	input := domain.Customer{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

	err := utils.ValidateStruct(req)

	if err != nil {
		u.logger.Errorf("error validating request: %v", err)
		return err
	}

	err = u.repo.CreateCustomer(ctx, input)

	if err != nil {
		u.logger.Errorf("error creating customer: %v", err)
		return err
	}
	return nil
}

func (u *customerUseCase) Delete(ctx context.Context, id string) error {
	err := u.repo.DeleteCustomer(ctx, id)
	if err != nil {
		u.logger.Errorf("error deleting customer: %v", err)
		return err
	}
	return nil
}

func (u *customerUseCase) UpdatePassword(ctx context.Context, req presenter.CustomerUpdate) error {
	acc := domain.Customer{
		Email:    req.Email,
		Password: req.Password,
	}

	err := utils.ValidateStruct(acc)
	if err != nil {
		u.logger.Errorf("error validating request: %v", err)
		return err
	}

	err = u.repo.UpdatePasswordCustomer(ctx, acc)
	if err != nil {
		u.logger.Errorf("error updating password: %v", err)
		return err
	}
	return nil
}

func (u *customerUseCase) FindByEmail(ctx context.Context, email string) (domain.Customer, error) {
	return u.repo.GetEmailCustomer(ctx, email)
}

func (u *customerUseCase) FindByID(ctx context.Context, id string) (domain.Customer, error) {
	return u.repo.GetIDCustomer(ctx, id)
}
