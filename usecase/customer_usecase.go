package usecase

import (
	"fmt"
	"go-laundry/model"
	"go-laundry/repository"
)

type CustomerUseCase interface {
	Create(payload model.Customer) error
	FindAll() ([]model.Customer, error)
	FindById(id string) (model.Customer, error)
	Update(payload model.Customer) error
	Delete(id string) error
}

type customerUseCase struct {
	repository repository.CustomerRepository
}

// Create implements CustomerUseCase.
func (useCase *customerUseCase) Create(payload model.Customer) error {
	if payload.Id == "" {
		return fmt.Errorf("id is required")
	}

	if payload.Name == "" {
		return fmt.Errorf("name is required")
	}

	if payload.PhoneNumber == ""  {
		return fmt.Errorf("phone number is required")
	}

	if payload.Address == ""  {
		return fmt.Errorf("address is required")
	}

	err := useCase.repository.Save(payload)
	if err != nil {
		return fmt.Errorf("cannot save new customer: %v", err)
	}

	return nil
}

// Delete implements CustomerUseCase.
func (useCase *customerUseCase) Delete(id string) error {
	_, err := useCase.FindById(id)
	if err != nil {
		return err
	}

	err = useCase.repository.Delete(id)
	if err != nil {
		return fmt.Errorf("failed delete customer:%v ", err)
	}
	return nil
}

// FindAll implements CustomerUseCase.
func (useCase *customerUseCase) FindAll() ([]model.Customer, error) {
	customers, err := useCase.repository.FindAll()
	if err != nil {
		return nil, fmt.Errorf("failed to get customers:%v ", err)
	}

	return customers, nil
}

// FindById implements CustomerUseCase.
func (useCase *customerUseCase) FindById(id string) (model.Customer, error) {
	customer, err := useCase.repository.FindById(id)
	if err != nil {
		return model.Customer{}, fmt.Errorf("customer with id: %s not found", id)
	}

	return customer, nil
}

// Update implements CustomerUseCase.
func (useCase *customerUseCase) Update(payload model.Customer) error {
	if payload.Id == "" {
		return fmt.Errorf("id is required")
	}

	// if payload.Name == "" {
	// 	return fmt.Errorf("name is required")
	// }

	// if payload.PhoneNumber == ""  {
	// 	return fmt.Errorf("phone number is required")
	// }

	// if payload.Address == ""  {
	// 	return fmt.Errorf("address is required")
	// }

	_, err := useCase.FindById(payload.Id)
	if err != nil {
		return err
	}

	err = useCase.repository.Update(payload)
	if err != nil {
		return fmt.Errorf("cannot save new customer: %v", err)
	}

	return nil
}

func NewCustomerUseCase(repository repository.CustomerRepository) CustomerUseCase {
	return &customerUseCase{
		repository: repository,
	}
}
