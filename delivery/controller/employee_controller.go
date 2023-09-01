package controller

import (
	"bufio"
	"fmt"
	"go-laundry/model"
	"go-laundry/usecase"
	"os"

	"github.com/rodaine/table"
)

type EmployeeController struct {
	employeeUseCase usecase.EmployeeUseCase
}

func (controller *EmployeeController) EmployeeMenuForm() {
	fmt.Println(`
	+++++ Master employee +++++
	1. Tambah Data
	2. Lihat Data
	3. Update Data
	4. Hapus Data
	5. Cari Data Berdasarkan Id
	6. Cari Data Berdasarkan No. Telp
	7. Keluar
	`)
	fmt.Print("Pilih Menu: \n")
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
		controller.ShowListCustomerByPhoneNumber()
	case "7":
		return
	}
}

func (controller *EmployeeController) insertFormEmployee() {
	var employee model.Employee
	fmt.Print("Inputkan Id: ")
	fmt.Scanln(&employee.Id)
	fmt.Print("Inputkan Name: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	employee.Name = scanner.Text()
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
	table := table.New("Id", "Name", "Phone Number", "Address")
	for _, employee := range employees {
		table.AddRow(employee.Id, employee.Name, employee.PhoneNumber, employee.Address)
	}

	table.Print()
}

func (controller *EmployeeController) updateFormemployee()  {
	var employee model.Employee
	fmt.Print("Inputkan employee Id: ")
	fmt.Scanln(&employee.Id)
	fmt.Print("Inputkan Nama employee Baru: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	employee.Name = scanner.Text()
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
	employee, err := controller.employeeUseCase.FindById(id)
	if err != nil {
		fmt.Println(err)
	}
	table := table.New("Id", "Name", "Phone Number", "Address")
		table.AddRow(employee.Id, employee.Name, employee.PhoneNumber, employee.Address)


	table.Print()
}

func (controller *EmployeeController) ShowListCustomerByPhoneNumber() {
	var phoneNumber string
	fmt.Print("Inputkan No. Hp :")
	fmt.Scanln(&phoneNumber)

	employee, err := controller.employeeUseCase.FindByPhoneNumber(phoneNumber)
	if err != nil {
		fmt.Println(err)
		return
	}
	table := table.New("Id", "Name", "Phone Number", "Address")
		table.AddRow(employee.Id, employee.Name, employee.PhoneNumber, employee.Address)


	table.Print()
}

func NewEmployeeController(employeeUseCase usecase.EmployeeUseCase) *EmployeeController  {
	return &EmployeeController{
		employeeUseCase: employeeUseCase,
	}
}