package domain

import (
	"github.com/viza/banking/app/dto"
	"github.com/viza/banking/app/errs"
)

type Customer struct {
	Id          string `db:"customer_id"`
	Name        string
	City        string
	Zipcode     string
	Dateofbirth string `db:"date_of_birth"`
	Status      string
}

func (c Customer) statusAsText() string {
	statusAsText := "active"
	if c.Status == "0" {
		statusAsText = "inactive"
	}
	return statusAsText
}

func (c Customer) ToDto() dto.CustomerResponse {
	return dto.CustomerResponse{
		Id:          c.Id,
		Name:        c.Name,
		City:        c.City,
		Zipcode:     c.Zipcode,
		Dateofbirth: c.Dateofbirth,
		Status:      c.statusAsText(),
	}
}

type CustomerRepository interface {
	// status - 1 is active, 0 - is inactive, or "" - expected all statuses
	FindAll(string) ([]Customer, *errs.AppError)
	ById(string) (*Customer, *errs.AppError)
}
