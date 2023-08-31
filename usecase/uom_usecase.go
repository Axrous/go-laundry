package usecase

import (
	"fmt"
	"go-laundry/model"
	"go-laundry/repository"
)

type UomUseCase interface {
	Create(payload model.Uom) error
	FindAll() ([]model.Uom, error)
	FindById(id string) (model.Uom, error)
	Update(payload model.Uom) error
	Delete(id string) error
}

type uomUseCase struct {
	uomRepository repository.UomRepository
}

// Create implements UomUseCase.
func (u *uomUseCase) Create(payload model.Uom) error {
	if payload.Id == "" {
		return fmt.Errorf("id is required")
	}

	if payload.Name == "" {
		return fmt.Errorf("name is required")
	}

	err := u.uomRepository.Save(payload)
	if err != nil {
		return fmt.Errorf("failed to create new uom, %v", err)
	}

	return nil
}

// Delete implements UomUseCase.
func (u *uomUseCase) Delete(id string) error {
	_, err := u.FindById(id)

	if err != nil {
		return err
	}

	err = u.uomRepository.DeleteById(id)
	if err != nil {
		return fmt.Errorf("cannot delete uom")
	}

	return nil
}

// FindAll implements UomUseCase.
func (u *uomUseCase) FindAll() ([]model.Uom, error) {
	uoms, err := u.uomRepository.FindAll()
	if err != nil {
		return nil, fmt.Errorf("cannot get all uoms")
	}

	return uoms, nil
}

// FindById implements UomUseCase.
func (u *uomUseCase) FindById(id string) (model.Uom, error) {
	uom, err := u.uomRepository.FindById(id)
	if err != nil {
		return model.Uom{}, fmt.Errorf("uom with id: %s not found", id)
	}

	return uom, nil
}

// Update implements UomUseCase.
func (u *uomUseCase) Update(payload model.Uom) error {
	if payload.Id == "" {
		return fmt.Errorf("id is required")
	}

	if payload.Name == "" {
		return fmt.Errorf("name is required")
	}

	_, err := u.FindById(payload.Id)
	if err != nil {
		return err
	}

	err = u.uomRepository.Update(payload)
	if err != nil {
		return fmt.Errorf("cannot update uom with id:%s", payload.Id)
	}

	return nil
}

func NewUomUseCase(uomRepository repository.UomRepository) UomUseCase {
	return &uomUseCase{
		uomRepository: uomRepository,
	}
}
