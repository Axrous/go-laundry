package controller

import (
	"fmt"
	"go-laundry/model"
	"go-laundry/usecase"

	"github.com/rodaine/table"
)

type CustomerController struct {
	customerUseCase usecase.CustomerUseCase
}

func (controller *CustomerController) CustomerMenuForm() {
	fmt.Println(`
	+++++ Master Customer +++++
	1. Tambah Data
	2. Lihat Data
	3. Update Data
	4. Hapus Data
	5. Cari Data Berdasarkan Id
	6. Cari Data Berdasarkan No. Telp
	7. Keluar
	`)
	fmt.Print("Pilih Menu (1-6): \n")
	var selectMenucustomer string
	fmt.Scanln(&selectMenucustomer)
	switch selectMenucustomer {
	case "1":
		controller.insertFormCustomer()
	case "2":
		controller.showListCustomer()
	case "3":
		controller.updateFormCustomer()
	case "4":
		controller.deleteFormCustomer()
	case "5":
		controller.showListCustomerById()
	case "6":
		controller.ShowListCustomerByPhoneNumber()
	case "7":
		return
	}
}

func (controller *CustomerController) insertFormCustomer() {
	var customer model.Customer
	fmt.Print("Inputkan Id: ")
	fmt.Scanln(&customer.Id)
	fmt.Print("Inputkan Name: ")
	fmt.Scanln(&customer.Name)
	fmt.Print("Inputkan No. Hp: ")
	fmt.Scanln(&customer.PhoneNumber)
	fmt.Print("Inputkan Alamat: ")
	fmt.Scanln(&customer.Address)
	err := controller.customerUseCase.Create(customer)
	if	err != nil {
		fmt.Println(err)
	}
}

func (controller *CustomerController) showListCustomer() {
	customers, err := controller.customerUseCase.FindAll()
	if err != nil {
		fmt.Println(err)
	}

	table := table.New("Id", "Name", "Phone Number", "Address")
	for _, customer := range customers {
		table.AddRow(customer.Id, customer.Name, customer.PhoneNumber, customer.Address)
	}

	table.Print()
}

func (controller *CustomerController) updateFormCustomer()  {
	var customer model.Customer
	fmt.Print("Inputkan customer Id: ")
	fmt.Scanln(&customer.Id)
	fmt.Print("Inputkan Nama customer Baru: ")
	fmt.Scanln(&customer.Name)
	fmt.Print("Inputkan No Hp customer Baru: ")
	fmt.Scanln(&customer.PhoneNumber)
	fmt.Print("Inputkan Alamat customer Baru: ")
	fmt.Scanln(&customer.Address)
	err := controller.customerUseCase.Update(customer)
	if	err != nil {
		fmt.Println(err)
	}
}

func (controller *CustomerController) deleteFormCustomer()  {
	var id string
	fmt.Print("Inputkan customer Id: ")
	fmt.Scanln(&id)
	err := controller.customerUseCase.Delete(id)
	if	err != nil {
		fmt.Println(err)
	}
}

func (controller *CustomerController) showListCustomerById() {
	var id string
	fmt.Print("Inputkan Id customer: ")
	fmt.Scanln(&id)
	customer, err := controller.customerUseCase.FindById(id)
	if err != nil {
		fmt.Println(err)
	}

	table := table.New("Id", "Name", "Phone Number", "Address")
	table.AddRow(customer.Id, customer.Name, customer.PhoneNumber, customer.Address)

	table.Print()
}

func (controller *CustomerController) ShowListCustomerByPhoneNumber() {
	var phoneNumber string
	fmt.Println("Inputkan No. Hp")
	fmt.Scanln(&phoneNumber)

	customer, err := controller.customerUseCase.FindByPhoneNumber(phoneNumber)
	if err != nil {
		fmt.Println(err)
		return
	}
	table := table.New("Id", "Name", "Phone Number", "Address")
	table.AddRow(customer.Id, customer.Name, customer.PhoneNumber, customer.Address)

	table.Print()
}

func NewCustomerController(customerUseCase usecase.CustomerUseCase) *CustomerController  {
	return &CustomerController{
		customerUseCase: customerUseCase,
	}
}