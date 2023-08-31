package usecase

import (
	"fmt"
	"go-laundry/model"
	"go-laundry/repository"
)

type EmployeeUseCase interface {
	Create(payload model.Employee) error
	FindAll() ([]model.Employee, error)
	FindById(id string) (model.Employee, error)
	Update(payload model.Employee) error
	Delete(id string) error
}

type employeeUseCase struct {
	repository repository.EmployeeRepository
}

// Create implements employeeUseCase.
func (useCase *employeeUseCase) Create(payload model.Employee) error {
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
		return fmt.Errorf("cannot save new employee: %v", err)
	}

	return nil
}

// Delete implements employeeUseCase.
func (useCase *employeeUseCase) Delete(id string) error {
	_, err := useCase.FindById(id)
	if err != nil {
		return err
	}

	err = useCase.repository.Delete(id)
	if err != nil {
		return fmt.Errorf("failed delete employee:%v ", err)
	}
	return nil
}

// FindAll implements employeeUseCase.
func (useCase *employeeUseCase) FindAll() ([]model.Employee, error) {
	employees, err := useCase.repository.FindAll()
	if err != nil {
		return nil, fmt.Errorf("failed to get employees:%v ", err)
	}

	return employees, nil
}

// FindById implements employeeUseCase.
func (useCase *employeeUseCase) FindById(id string) (model.Employee, error) {
	employee, err := useCase.repository.FindById(id)
	if err != nil {
		return model.Employee{}, fmt.Errorf("employee with id: %s not found", id)
	}

	return employee, nil
}

// Update implements employeeUseCase.
func (useCase *employeeUseCase) Update(payload model.Employee) error {
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

	_, err := useCase.FindById(payload.Id)
	if err != nil {
		return err
	}

	err = useCase.repository.Update(payload)
	if err != nil {
		return fmt.Errorf("cannot save new employee: %v", err)
	}

	return nil
}

func NewemployeeUseCase(repository repository.EmployeeRepository) EmployeeUseCase {
	return &employeeUseCase{
		repository: repository,
	}
}
