package main

import "go-laundry/delivery"

func main() {
	delivery.NewConsole().Run()

	// cfg, err := config.NewConfig()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// con, err := config.NewDbConnection(cfg)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// db := con.Conn()

	// // //repository
	// // uomRepo := repository.NewUomRepository(db)
	// // productRepo := repository.NewProductRepository(db)
	// // customerRepo := repository.NewCustomerRepository(db)
	// // employeeRepo := repository.NewEmployeeRepository(db)

	// // //Use Case
	// // uomUseCase := usecase.NewUomUseCase(uomRepo)
	// // productUseCase := usecase.NewProductUseCase(productRepo, uomUseCase)
	// // customerUseCase := usecase.NewCustomerUseCase(customerRepo)
	// // employeeUseCase := usecase.NewemployeeUseCase(employeeRepo)

	// billRepo := repository.NewBillRepository(db)
	// billUseCase := usecase.NewBillUseCase(billRepo)

	// bill := model.Bill{
	// 	Id:         "",
	// 	BillDate:   time.Now(),
	// 	EntryDate:  time.Now(),
	// 	FinishDate: time.Now(),
	// 	Employee:   model.Employee{
	// 		Id:          "1",
	// 	},
	// 	Customer:   model.Customer{
	// 		Id:          "2",
	// 	},
	// 	BillDetail: []model.BillDetail{
	// 		{
	// 			Id:           "6",
	// 			BillId:       "4",
	// 			Product:      model.Product{
	// 				Id:    "4",
	// 			},
	// 			ProductPrice: 2000,
	// 			Qty:          3,
	// 		},
	// 		{
	// 			Id:           "7",
	// 			BillId:       "4",
	// 			Product:      model.Product{
	// 				Id:    "3",
	// 			},
	// 			ProductPrice: 5000,
	// 			Qty:          1,
	// 		},
	// 	},
	// }

	// // err = billRepo.Save(bill)
	// err = billUseCase.Create(bill)
	// if err != nil {
	// 	fmt.Println(err)
	// }


}