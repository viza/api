package domain

import "github.com/viza/banking/app/errs"

type Customer struct {
	Id          string `db:"customer_id"`
	Name        string
	City        string
	Zipcode     string
	Dateofbirth string `db:"date_of_birth"`
	Status      string
}

type CustomerRepository interface {
	// status - 1 is active, 0 - is inactive, or "" - expected all statuses
	FindAll(string) ([]Customer, *errs.AppError)
	ById(string) (*Customer, *errs.AppError)
}
