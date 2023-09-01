package delivery

import (
	"fmt"
	"go-laundry/config"
	"go-laundry/delivery/controller"
	"go-laundry/repository"
	"go-laundry/usecase"
	"os"
)

type Console struct {
	uomUseCase usecase.UomUseCase
	productuseCase usecase.ProductUseCase
	customerUseCase usecase.CustomerUseCase
	employeeUseCase usecase.EmployeeUseCase
	billUseCase usecase.BillUseCase
}

func (c *Console) showMainMenu() {
	fmt.Println(`
	+++++ Enigma Laundry Menu +++++
	1. Master UOM
	2. Master Product
	3. Master Customer
	4. Master Eployee
	5. Transaksi
	6. Keluar
	`)
	fmt.Print("Pilih Menu (1-6): ")
}



func (c *Console) Run() {
	for {
		c.showMainMenu()
		var selectedMenu string
		fmt.Scanln(&selectedMenu)
		switch selectedMenu {
		case "1":
			controller.NewUomController(c.uomUseCase).UomMenuForm()
		case "2":
			controller.NewProductController(c.productuseCase, c.uomUseCase).UProductMenuForm()
		case "3":
			controller.NewCustomerController(c.customerUseCase).CustomerMenuForm()
		case "4":
			controller.NewEmployeeController(c.employeeUseCase).EmployeeMenuForm()
		case "5":
			controller.NewBillController(c.billUseCase).BillMenuForm()
		case "6":
			os.Exit(0)
		}
	}
}

func NewConsole() *Console {
	cfg, err := config.NewConfig()
	if err != nil {
		fmt.Println(err)
	}
	con, err := config.NewDbConnection(cfg)
	if err != nil {
		fmt.Println(err)
	}
	db := con.Conn()

	//repository
	uomRepo := repository.NewUomRepository(db)
	productRepo := repository.NewProductRepository(db)
	customerRepo := repository.NewCustomerRepository(db)
	employeeRepo := repository.NewEmployeeRepository(db)
	billRepo := repository.NewBillRepository(db)

	//Use Case
	uomUseCase := usecase.NewUomUseCase(uomRepo)
	productUseCase := usecase.NewProductUseCase(productRepo, uomUseCase)
	customerUseCase := usecase.NewCustomerUseCase(customerRepo)
	employeeUseCase := usecase.NewemployeeUseCase(employeeRepo)
	billUsecase := usecase.NewBillUseCase(billRepo, employeeUseCase, customerUseCase, productUseCase)

	return &Console{
		uomUseCase:      uomUseCase,
		productuseCase:  productUseCase,
		customerUseCase: customerUseCase,
		employeeUseCase: employeeUseCase,
		billUseCase:     billUsecase,
	}
}
