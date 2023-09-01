package controller

import (
	"fmt"
	"go-laundry/model"
	"go-laundry/usecase"

	"github.com/rodaine/table"
)

type UomController struct {
	uomUseCase usecase.UomUseCase
}

func (u *UomController) UomMenuForm() {
	fmt.Println()
	fmt.Println(`
	+++++ Master UOM +++++
	1. Tambah Data
	2. Lihat Data
	3. Update Data
	4. Hapus Data
	5. Cari Data Berdasarkan Id
	6. Keluar
	`)
	fmt.Print("Pilih Menu: ")
	var selectMenuUom string
	fmt.Scanln(&selectMenuUom)
	switch selectMenuUom {
	case "1":
		u.insertFormUom()
	case "2":
		u.ShowListUom()
	case "3":
		u.UpdateFormUom()
	case "4":
		u.DeleteFormUom()
	case "5":
		u.ShowListUomById()
	case "6":
		return
	}
}

func (u *UomController) insertFormUom() {
	var uom model.Uom
	fmt.Print("Inputkan Id: ")
	fmt.Scanln(&uom.Id)
	fmt.Print("Inputkan Name: ")
	fmt.Scanln(&uom.Name)
	err := u.uomUseCase.Create(uom)
	if	err != nil {
		fmt.Println(err)
	}
}

func (u *UomController) ShowListUom() {
	uoms, err := u.uomUseCase.FindAll()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("DATA UOM")
	table := table.New("Id", "Name")
	for _, uom := range uoms {
		table.AddRow(uom.Id, uom.Name)
	}

	table.Print()
}

func (u *UomController) UpdateFormUom()  {
	var uom model.Uom
	fmt.Print("Inputkan Uom Id: ")
	fmt.Scanln(&uom.Id)
	fmt.Print("Inputkan Nama Uom Baru: ")
	fmt.Scanln(&uom.Name)
	err := u.uomUseCase.Update(uom)
	if	err != nil {
		fmt.Println(err)
	}
}

func (u *UomController) DeleteFormUom()  {
	var id string
	fmt.Print("Inputkan Uom Id: ")
	fmt.Scanln(&id)
	err := u.uomUseCase.Delete(id)
	if	err != nil {
		fmt.Println(err)
	}
}

func (u *UomController) ShowListUomById() {
	var id string
	fmt.Print("Inputkan Id Uom: ")
	fmt.Scanln(&id)
	uoms, err := u.uomUseCase.FindById(id)
	if err != nil {
		fmt.Println(err)
	}

	table := table.New("Id", "Name")
	table.AddRow(uoms.Id, uoms.Name)

	table.Print()
}

func NewUomController(uomUseCase usecase.UomUseCase) *UomController  {
	return &UomController{
		uomUseCase: uomUseCase,
	}
}