package repositories

import (
	"context"
	"database/sql"

	"github.com/adilsonmenechini/golabbank/internal/domain"
	"github.com/adilsonmenechini/golabbank/pkg/utils"
)

type (
	CustomerRepository interface {
		CreateCustomer(ctx context.Context, customer domain.Customer) error
		DeleteCustomer(ctx context.Context, id string) error
		GetEmailCustomer(ctx context.Context, email string) (domain.Customer, error)
		GetIDCustomer(ctx context.Context, id string) (domain.Customer, error)
		UpdatePasswordCustomer(ctx context.Context, customer domain.Customer) error
	}

	customerRepository struct {
		logger *utils.Logger
		db     *sql.DB
	}
)

func NewCustomerRepository(DB *sql.DB) CustomerRepository {
	return &customerRepository{
		logger: utils.NewLogger("customerRepository"),
		db:     DB,
	}
}

const (
	createCustomer         = `INSERT INTO customers (id, name, email, password, created_at) VALUES ( $1, $2, $3, $4, $5)`
	deleteCustomer         = `DELETE FROM customers WHERE id = $1`
	getEmailCustomer       = `SELECT id, name, email, password, created_at FROM customers WHERE email = $1 LIMIT 1`
	updatePasswordCustomer = `UPDATE customers set password = $2 WHERE email = $1`
	getIDCustomer          = `SELECT id, name, email, password, created_at FROM customers WHERE id = $1 LIMIT 1`
)

func (ra *customerRepository) CreateCustomer(ctx context.Context, customer domain.Customer) error {
	idGen := utils.GenerateUUID()

	newAcc := domain.NewCustomer(idGen, customer.Name, customer.Email, customer.Password)

	_, err := ra.db.ExecContext(ctx, createCustomer,
		newAcc.ID,
		newAcc.Name,
		newAcc.Email,
		newAcc.Password,
		newAcc.CreatedAt,
	)
	if err != nil {
		return err
	}
	return nil
}

func (ra *customerRepository) DeleteCustomer(ctx context.Context, id string) error {
	_, err := ra.db.ExecContext(ctx, deleteCustomer, id)
	if err != nil {
		return err
	}
	return nil
}

func (ra *customerRepository) UpdatePasswordCustomer(ctx context.Context, customer domain.Customer) error {
	_, err := ra.db.ExecContext(ctx, updatePasswordCustomer, customer.Email, customer.Password)
	if err != nil {
		return err
	}
	return nil
}

func (ra *customerRepository) GetEmailCustomer(ctx context.Context, email string) (domain.Customer, error) {
	row := ra.db.QueryRowContext(ctx, getEmailCustomer, email)
	var i domain.Customer
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Password,
		&i.CreatedAt,
	)
	if err != nil {
		return i, err
	}

	return i, err
}

func (ra *customerRepository) GetIDCustomer(ctx context.Context, id string) (domain.Customer, error) {
	row := ra.db.QueryRowContext(ctx, getIDCustomer, id)
	var i domain.Customer
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Password,
		&i.CreatedAt,
	)
	if err != nil {
		return i, err
	}

	return i, err
}
