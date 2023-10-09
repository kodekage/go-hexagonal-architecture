package domain

import (
	"github.com/kodekage/banking/errs"
)

type Customer struct {
	Id      string `db:"customer_id"`
	Name    string
	City    string
	Zipcode string
	DOB     string `db:"date_of_birth"`
	Status  string
}

type Filters struct {
	Status string
}

type CustomerRepository interface {
	FindAll(filters Filters) ([]Customer, *errs.AppError)
	FindById(string) (*Customer, *errs.AppError)
}
