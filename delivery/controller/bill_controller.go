package controller

import (
	"fmt"
	"go-laundry/model"
	"go-laundry/usecase"

	"github.com/rodaine/table"
)

type BillController struct {
	useCase usecase.BillUseCase
}

func (controller *BillController) BillMenuForm() {
	fmt.Println(`
	+++++ Transaksi +++++
	1. Tambah Data
	2. Lihat Data
	3. Update Data By Id
	4. Keluar
	`)
	fmt.Print("Pilih Menu: ")
	var selectMenucustomer string
	fmt.Scanln(&selectMenucustomer)
	switch selectMenucustomer {
	case "1":
		controller.insertBillForm()
	case "2":
		controller.showListBills()
	case "3":
		controller.showListBillById()
	case "4":
		return
	}
}

func (controller *BillController) insertBillForm()  {
	
	var bill model.Bill
	fmt.Print("Masukan id transaksi: ")
	fmt.Scanln(&bill.Id)
	fmt.Print("Masukan id Employee: ")
	fmt.Scanln(&bill.Employee.Id)
	fmt.Print("Masukan id Customer: ")
	fmt.Scanln(&bill.Customer.Id)

	//transaksi detail

	var billDetails []model.BillDetail
	loop := true
	for loop {
		var billDetail model.BillDetail
		fmt.Print("Masukan id detail: ")
		fmt.Scanln(&billDetail.Id)
		fmt.Print("Masukan product id: ")
		fmt.Scanln(&billDetail.Product.Id)
		fmt.Print("Masukan product QTY: ")
		fmt.Scanln(&billDetail.Qty)
		billDetail.BillId = bill.Id
		billDetails = append(billDetails, billDetail)
		fmt.Print("Masukan data lagi? (true/false) : ")
		fmt.Scanln(&loop)
	}

	bill.BillDetail = billDetails
	err := controller.useCase.Create(bill)
	if err != nil {
		fmt.Println(err)
	}
}

func (controller *BillController) showListBills() {
	bills, err := controller.useCase.FindAll()
	if err != nil {
		fmt.Println(err)
	}

	table := table.New("Id", "Bill Date", "Entry Date", "Finish Date", "Empploye", "Customer")
	for _, bill := range bills {
		billDate := bill.BillDate.Format("MM-DD-YYYY")
		table.AddRow(bill.Id, billDate, bill.EntryDate, bill.FinishDate, bill.Employee.Name, bill.Customer.Name)
	}

	table.Print()
}

func (controller *BillController) showListBillById()  {
	var id string
	fmt.Print("Masukkan Id Bill: ")
	fmt.Scan(&id)
	bills, err := controller.useCase.FindById(id)
	if err != nil {
		fmt.Println(err)
	}

	table := table.New("Id", "Bill Date", "Entry Date", "Finish Date", "Employee", "Customer", "Product", "QTY", "Price")
	for _, bill := range bills {
		table.AddRow(bill.Id, bill.BillDate, bill.EntryDate, bill.FinishDate, bill.Employee.Name, bill.Customer.Name, bill.Product.Name, bill.Qty, bill.ProductPrice)
	}

	table.Print()
}


func NewBillController(useCase usecase.BillUseCase) *BillController  {
	return &BillController{
		useCase: useCase,
	}
}