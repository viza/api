package service

import (
	"github.com/viza/banking/app/dto"
	"github.com/viza/banking/app/errs"
	"github.com/viza/banking/domain"
)

type CustomerService interface {
	GetAllCustomer(string) ([]domain.Customer, *errs.AppError)
	GetCustomer(string) (*dto.CustomerResponse, *errs.AppError)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomer(status string) ([]domain.Customer, *errs.AppError) {
	if status == "active" { // localhost:8000/customers?status=active
		status = "1"
	} else if status == "inactive" { // localhost:8000/customers?status=inactive
		status = "0"
	} else { // localhost:8000/customers
		status = ""
	}

	return s.repo.FindAll(status)
}

func (s DefaultCustomerService) GetCustomer(id string) (*dto.CustomerResponse, *errs.AppError) {

	c, err := s.repo.ById(id)
	if err != nil {
		return nil, err
	}
	response := c.ToDto()

	return &response, nil
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository}
}
