package services

import (
	"bookstore/domain/aggregate"
	"bookstore/domain/repository"
)

type CustomerService struct {
	repository repository.ICustomerRepository
}

type ICustumers interface {
}

func NewCustomerService(repository repository.ICustomerRepository) *CustomerService {
	return &CustomerService{repository: repository}
}

func (s *CustomerService) CreateCustomers(entity *aggregate.Customer) error {
	if err := s.repository.Create(entity); err != nil {
		return err
	}

	return nil
}

func (s *CustomerService) GetCustomers() ([]*aggregate.Customer, error) {
	costumers, err := s.repository.FindAll()
	if err != nil {
		return nil, err
	}
	return costumers, nil
}
