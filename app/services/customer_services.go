package services

import (
	"bookstore/domain/aggregate"
	"bookstore/domain/repository"
)

type ICustomerService interface {
	Create(entity *aggregate.Customer) error
	Update(entity *aggregate.Customer) error
	GetById(id string) (*aggregate.Customer, error)
	GetAll() ([]*aggregate.Customer, error)
	Remove(id string) error
}

type customerService struct {
	ICustomerRepository repository.ICustomerRepository
}

func NewCustomerServices(IcustomerRepository repository.ICustomerRepository) ICustomerService {
	return &customerService{
		ICustomerRepository: IcustomerRepository,
	}
}

func (c *customerService) Create(entity *aggregate.Customer) error {
	return c.ICustomerRepository.Create(entity)
}

func (c *customerService) Update(entity *aggregate.Customer) error {
	return c.ICustomerRepository.Update(entity)
}

func (c *customerService) GetById(id string) (*aggregate.Customer, error) {
	result, err := c.ICustomerRepository.FindByID(id)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *customerService) GetAll() ([]*aggregate.Customer, error) {
	result, err := c.ICustomerRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *customerService) Remove(id string) error {
	return c.ICustomerRepository.Delete(id)
}
