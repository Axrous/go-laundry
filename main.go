package main

import (
	"fmt"
	"go-laundry/config"
	"go-laundry/model"
	"go-laundry/repository"
)

type Customer struct{
	Id string
	Name string
	PhoneNumber string
	Address string
}

func main() {

	cfg, err := config.NewConfig()
	if err !=nil {
		fmt.Println(err)
	}

	con, err := config.NewDbConnection(cfg)
	if err != nil {
		fmt.Println(err)
	}
	db := con.Conn()

	uomRepo := repository.NewUomRepository(db)
	productRepo := repository.NewProductRepository(db)

	//save product
	// err = productRepo.Save(model.Product{
	// 	Id: "3",
	// 	Name: "Setrika Kilat",
	// 	Price: 5000,
	// 	Uom: model.Uom{
	// 		Id: "2",
	// 	},
	// })
	
	err = uomRepo.Update(model.Uom{
		Id: "",
		Name: "",
	})

	// if err != nil {
	// 	fmt.Println(err)
	// }

	//find by id product
	// product, err := productRepo.FindById("1")

	//update name product
	// err = 	productRepo.Update(model.Product{
	// 	Id: "2",
	// 	Name: "Setrika Kilat 1 Jam!",
	// })

	// delete product
	// err = productRepo.Delete("1")

	//find all product
	products, err := productRepo.FindAll()
	if err != nil {
		fmt.Println(err)
	}

	for _, product := range products {
		fmt.Printf("id: %s, Layanan: %s, Harga: %d, Satuan: %s", product.Id, product.Name, product.Price, product.Uom.Name)
		fmt.Println()
		fmt.Println("===================================================================")
	}



}