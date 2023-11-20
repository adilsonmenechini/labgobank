package domain

import (
	"time"

	"github.com/adilsonmenechini/golabbank/pkg/utils"
)

type Customer struct {
	ID        string
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
}

func NewCustomer(id, name, email, password string) Customer {
	pwd := utils.HashPassword(password)
	return Customer{
		ID:        id,
		Name:      name,
		Email:     email,
		Password:  pwd,
		CreatedAt: time.Now().UTC(),
	}
}
