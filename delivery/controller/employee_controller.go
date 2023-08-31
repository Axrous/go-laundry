package controller

import (
	"fmt"
	"go-laundry/model"
	"go-laundry/usecase"
)

type EmployeeController struct {
	employeeUseCase usecase.EmployeeUseCase
}

func (controller *EmployeeController) EmployeeMenuForm() {
	fmt.Println(`
	|		+++++ Master employee +++++		|
	| 1. Tambah Data					|
	| 2. Lihat Data						|
	| 3. Update Data					|
	| 4. Hapus Data						|
	| 5. Cari Data Berdasarkan Id		|
	| 6. Keluar                     	|
	`)
	fmt.Print("Pilih Menu (1-6): \n")
	var selectMenuemployee string
	fmt.Scanln(&selectMenuemployee)
	switch selectMenuemployee {
	case "1":
		controller.insertFormEmployee()
	case "2":
		controller.showListemployee()
	case "3":
		controller.updateFormemployee()
	case "4":
		controller.deleteFormemployee()
	case "5":
		controller.showListemployeeById()
	case "6":
		return
	}
}

func (controller *EmployeeController) insertFormEmployee() {
	var employee model.Employee
	fmt.Print("Inputkan Id: ")
	fmt.Scanln(&employee.Id)
	fmt.Print("Inputkan Name: ")
	fmt.Scanln(&employee.Name)
	fmt.Print("Inputkan No. Hp: ")
	fmt.Scanln(&employee.PhoneNumber)
	fmt.Print("Inputkan Alamat: ")
	fmt.Scanln(&employee.Address)
	err := controller.employeeUseCase.Create(employee)
	if	err != nil {
		fmt.Println(err)
	}
}

func (controller *EmployeeController) showListemployee() {
	employees, err := controller.employeeUseCase.FindAll()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(employees)
}

func (controller *EmployeeController) updateFormemployee()  {
	var employee model.Employee
	fmt.Print("Inputkan employee Id: ")
	fmt.Scanln(&employee.Id)
	fmt.Print("Inputkan Nama employee Baru: ")
	fmt.Scanln(&employee.Name)
	fmt.Print("Inputkan No Hp employee Baru: ")
	fmt.Scanln(&employee.PhoneNumber)
	fmt.Print("Inputkan Alamat employee Baru: ")
	fmt.Scanln(&employee.Address)
	err := controller.employeeUseCase.Update(employee)
	if	err != nil {
		fmt.Println(err)
	}
}

func (controller *EmployeeController) deleteFormemployee()  {
	var id string
	fmt.Print("Inputkan employee Id: ")
	fmt.Scanln(&id)
	err := controller.employeeUseCase.Delete(id)
	if	err != nil {
		fmt.Println(err)
	}
}

func (controller *EmployeeController) showListemployeeById() {
	var id string
	fmt.Print("Inputkan Id employee: ")
	fmt.Scanln(&id)
	employees, err := controller.employeeUseCase.FindById(id)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(employees)
}

func NewEmployeeController(employeeUseCase usecase.EmployeeUseCase) *EmployeeController  {
	return &EmployeeController{
		employeeUseCase: employeeUseCase,
	}
}