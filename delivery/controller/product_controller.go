package controller

import (
	"fmt"
	"go-laundry/model"
	"go-laundry/usecase"

	"github.com/rodaine/table"
)

type ProductController struct {
	productUseCase usecase.ProductUseCase
	uomUseCase usecase.UomUseCase
}

func (p *ProductController) UProductMenuForm() {
	fmt.Println()
	fmt.Println(`
	+++++ Master Produk +++++
	1. Tambah Data
	2. Lihat Data
	3. Update Data
	4. Hapus Data
	5. Cari Data Berdasarkan Nama
	6. Keluar
	`)
	fmt.Print("Pilih Menu (1-6): \n")
	var selectMenuUom string
	fmt.Scanln(&selectMenuUom)
	switch selectMenuUom {
	case "1":
		p.insertFormproduct()
	case "2":
		p.showListProducts()
	case "3":
		p.updateFormproduct()
	case "4":
		p.deleteFormProduct()
	case "5":
		p.showListproductByName()
	case "6":
		return
	}
}

func (p *ProductController) insertFormproduct() {
	var product model.Product

	uoms, err := p.uomUseCase.FindAll()
	if err != nil  {
		fmt.Println(err)
	}

	fmt.Println(uoms)

	fmt.Print("Inputkan Id: ")
	fmt.Scanln(&product.Id)
	fmt.Print("Inputkan Nama Produk: ")
	fmt.Scanln(&product.Name)
	fmt.Print("Inputkan Harga Produk: ")
	fmt.Scanln(&product.Price)
	fmt.Print("Inputkan Id Uom: ")
	fmt.Scanln(&product.Uom.Id)

	err = p.productUseCase.Create(product)
	if	err != nil {
		fmt.Println(err)
	}
}

func (p *ProductController) showListProducts()  {
	products, err := p.productUseCase.FindAll()
	if err != nil {
		fmt.Println(err)
	}
	table := table.New("Id", "Name", "Price", "Uom")
	for _, product := range products {
		table.AddRow(product.Id, product.Name, product.Price, product.Uom.Name)
	}

	table.Print()
}

func (p *ProductController) updateFormproduct() {
	var product model.Product

	uoms, err := p.uomUseCase.FindAll()
	if err != nil  {
		fmt.Println(err)
	}

	fmt.Println(uoms)

	fmt.Print("Inputkan Id: ")
	fmt.Scanln(&product.Id)
	fmt.Print("Inputkan Nama Produk Baru: ")
	fmt.Scanln(&product.Name)
	fmt.Print("Inputkan Harga Produk Baru: ")
	fmt.Scanln(&product.Price)
	fmt.Print("Inputkan Id Uom Baru: ")
	fmt.Scanln(&product.Uom.Id)

	err = p.productUseCase.Update(product)
	if	err != nil {
		fmt.Println(err)
	}
}

func (p *ProductController) deleteFormProduct()  {
	p.showListProducts()
	var id string
	fmt.Print("Inputkan Product Id: ")
	fmt.Scanln(&id)
	err := p.productUseCase.Delete(id)
	if	err != nil {
		fmt.Println(err)
	}
}

func (p *ProductController) showListproductByName() {
	var name string
	fmt.Print("Inputkan nama produk: ")
	fmt.Scanln(&name)
	products, err := p.productUseCase.FindByName(name)
	if err != nil {
		fmt.Println(err)
	}
	table := table.New("Id", "Name", "Price", "Uom")
	for _, product := range products {
		table.AddRow(product.Id, product.Name, product.Price, product.Uom.Name)
	}

	table.Print()
}

func NewProductController(productUseCase usecase.ProductUseCase, uomUseCase usecase.UomUseCase) *ProductController  {
	return &ProductController{
		productUseCase: productUseCase,
		uomUseCase:  uomUseCase,
	}
}