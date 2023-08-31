package usecase

import (
	"fmt"
	"go-laundry/model"
	"go-laundry/repository"
)

type ProductUseCase interface {
	Create(payload model.Product) error
	FindAll() ([]model.Product, error)
	FindById(id string) (model.Product, error)
	FindByName(name string) ([]model.Product, error)
	Update(payload model.Product) error
	Delete(id string) error
}

type productUseCase struct {
	repository repository.ProductRepository
	uomUseCase UomUseCase
}

// FindByName implements ProductUseCase.
func (p *productUseCase) FindByName(name string) ([]model.Product, error) {
	product, err := p.repository.FindByName(name)
	if err != nil {
		return nil, fmt.Errorf("error dulu, %v", err)
	}

	return product, nil
}

// Create implements ProductUseCase.
func (p *productUseCase) Create(payload model.Product) error {
	if payload.Id == "" {
		return fmt.Errorf("id is required")
	}

	if payload.Name == "" {
		return fmt.Errorf("name is required")
	}

	if payload.Price <= 0 {
		return fmt.Errorf("price must greater than zero")
	}

	_, err := p.uomUseCase.FindById(payload.Uom.Id)
	if err != nil {
		return err
	}

	err = p.repository.Save(payload)
	if err != nil {
		return fmt.Errorf("cannot save new product: %v", err)
	}

	return nil
}

// Delete implements ProductUseCase.
func (p *productUseCase) Delete(id string) error {
	_, err := p.FindById(id)
	if err != nil {
		return err
	}

	err = p.repository.Delete(id)
	if err != nil {
		return fmt.Errorf("failed delete product:%v ", err)
	}
	return nil
}

// FindAll implements ProductUseCase.
func (p *productUseCase) FindAll() ([]model.Product, error) {
	products, err := p.repository.FindAll()
	if err != nil {
		return nil, fmt.Errorf("failed to get products:%v ", err)
	}

	return products, nil
}

// FindById implements ProductUseCase.
func (p *productUseCase) FindById(id string) (model.Product, error) {
	product, err := p.repository.FindById(id)
	if err != nil {
		return model.Product{}, fmt.Errorf("product with id: %s not found", id)
	}

	return product, nil

}

// Update implements ProductUseCase.
func (p *productUseCase) Update(payload model.Product) error {

	if payload.Name == "" {
		return fmt.Errorf("name is required")
	}

	if payload.Price < 0 {
		return fmt.Errorf("price must greater than zero")
	}

	_, err := p.FindById(payload.Id)
	if err != nil {
		return err
	}

	err = p.repository.Update(payload)
	if err != nil {
		return fmt.Errorf("cannot update product: %v", err)
	}

	return nil
}

func NewProductUseCase(repository repository.ProductRepository, uomUseCase UomUseCase) ProductUseCase {
	return &productUseCase{
		repository: repository,
		uomUseCase: uomUseCase,
	}
}
