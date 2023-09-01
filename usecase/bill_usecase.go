package usecase

import (
	"fmt"
	"go-laundry/model"
	"go-laundry/repository"
	"time"
)

type BillUseCase interface {
	Create(payload model.Bill) error
	FindAll() ([]model.Bill, error)
	FindById(id string) ([]model.BillDetailShow, error)
}

type billUseCase struct {
	repository      repository.BillRepository
	employeUseCase  EmployeeUseCase
	customerUseCase CustomerUseCase
	productUseCase  ProductUseCase
}

// FindById implements BillUseCase.
func (useCase *billUseCase) FindById(id string) ([]model.BillDetailShow, error) {
	bills, err := useCase.repository.FindById(id)
	if err != nil {
		return nil ,fmt.Errorf("failed get bills with id: %s, %v", id, err)
	}

	return bills, nil

}

// FindAll implements BillUseCase.
func (useCase *billUseCase) FindAll() ([]model.Bill, error) {
	bills, err := useCase.repository.FindAll()
	if err != nil {
		return nil, fmt.Errorf("failed to get all bills: %v", err)
	}

	return bills, nil
}

// Create implements BillUseCase.
func (useCase *billUseCase) Create(payload model.Bill) error {
	if payload.Id == "" {
		return fmt.Errorf("id tidak boleh kosong")
	}
	if payload.Employee.Id == "" {
		return fmt.Errorf("id Pegawai tidak boleh kosong")
	}
	if payload.Customer.Id == "" {
		return fmt.Errorf("id Customer tidak boleh kosong")
	}

	_, err := useCase.employeUseCase.FindById(payload.Employee.Id)
	if err != nil {
		return err
	}

	_, err = useCase.customerUseCase.FindById(payload.Customer.Id)
	if err != nil {
		return err
	}

	var billDetails []model.BillDetail

	for _, billDetail := range payload.BillDetail {
		if billDetail.Id == "" {
			return fmt.Errorf("id Bill Detail tidak boleh kosong")
		}
		if billDetail.Product.Id == "" {
			return fmt.Errorf("id Produk Detail tidak boleh kosong")
		}

		product, err := useCase.productUseCase.FindById(billDetail.Product.Id)
		if err != nil {
			return err
		}

		billDetail.ProductPrice = billDetail.Qty * product.Price
		billDetails = append(billDetails, billDetail)
	}

	payload.BillDate = time.Now()
	payload.EntryDate = time.Now()

	payload.BillDetail = billDetails

	err = useCase.repository.Save(payload)
	if err != nil {
		return fmt.Errorf("gagal insert transaksi: %v", err)
	}

	return nil

}

func NewBillUseCase(
	repository repository.BillRepository,
	employeeUseCase EmployeeUseCase,
	customerUseCase CustomerUseCase,
	productUseCase ProductUseCase,
) BillUseCase {
	return &billUseCase{
		repository:      repository,
		employeUseCase:  employeeUseCase,
		customerUseCase: customerUseCase,
		productUseCase:  productUseCase,
	}
}
