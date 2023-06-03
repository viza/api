package service

import (
	"github.com/viza/banking/app/errs"
	"github.com/viza/banking/domain"
)

type CustomerService interface {
	GetAllCustomer() ([]domain.Customer, error)
	GetAllCustomersByStatus(string) ([]domain.Customer, *errs.AppError)
	GetCustomer(string) (*domain.Customer, *errs.AppError)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomer() ([]domain.Customer, error) {
	return s.repo.FindAll()
}

func (s DefaultCustomerService) GetAllCustomersByStatus(status string) ([]domain.Customer, *errs.AppError) {
	return s.repo.ByStatus(status)
}

func (s DefaultCustomerService) GetCustomer(id string) (*domain.Customer, *errs.AppError) {
	return s.repo.ById(id)
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository}
}
