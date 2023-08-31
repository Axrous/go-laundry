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
}

func (c *Console) showMainMenu() {
	fmt.Println(`
	|+++++ Enigma Laundry Menu +++++|
	| 1. Master UOM                 |
	| 2. Master Product             |
	| 3. Master Customer            |
	| 4. Master Eployee             |
	| 5. Transaksi                  |
	| 6. Keluar                     |
	`)
	fmt.Println("Pilih Menu (1-6): ")
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
		case "3":
		case "4":
		case "5":
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
	uomRepo := repository.NewUomRepository(db)
	uomUseCase := usecase.NewUomUseCase(uomRepo)
	return &Console{
		uomUseCase: uomUseCase,
	}
}
