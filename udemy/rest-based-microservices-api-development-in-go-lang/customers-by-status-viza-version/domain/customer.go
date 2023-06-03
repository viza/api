package domain

import "github.com/viza/banking/app/errs"

type Customer struct {
	Id          string
	Name        string
	City        string
	Zipcode     string
	Dateofbirth string
	Status      string
}

type CustomerRepository interface {
	FindAll() ([]Customer, error)
	ById(string) (*Customer, *errs.AppError)
	ByStatus(string) ([]Customer, *errs.AppError)
}
