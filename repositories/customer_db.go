package repositories

import "github.com/yeneichen/injection-go/entities"

type CustomerRepository struct {
	CustomerRepo entities.CustomerRepository
}

func NewCustomerRepository(customerRepo entities.CustomerRepository) *CustomerRepository {
	return &CustomerRepository{
		CustomerRepo: customerRepo,
	}
}

func (r *CustomerRepository) GetAll() ([]entities.Customer, error) {
	return r.CustomerRepo.GetAll()
}

func (r *CustomerRepository) GetById(id string) (entities.Customer, error) {
	return r.CustomerRepo.GetById(id)
}
