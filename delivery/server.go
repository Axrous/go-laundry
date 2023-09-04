package delivery

import (
	"fmt"
	"go-laundry/config"
	"go-laundry/delivery/controller/api"
	"go-laundry/repository"
	"go-laundry/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server struct {
	uomUseCase usecase.UomUseCase
	engine *gin.Engine
}

func (server *Server) Run()  {

	err := server.engine.Run()
	if err != nil {
		panic(err)
	}
}

func (server *Server) InitController()  {
	api.NewUomController(server.uomUseCase, server.engine)
}



func NewServer() *Server {

	cfg, err := config.NewConfig()
	if err != nil {
		fmt.Println(err)
	}
	con, err := config.NewDbConnection(cfg)
	if err != nil {
		fmt.Println(err)
	}
	db := con.Conn()

	//Repository
	UomRepo := repository.NewUomRepository(db)
	//UseCase
	uomUseCase := usecase.NewUomUseCase(UomRepo)

	engine := gin.Default()
	http.ListenAndServe("localhost:8080", engine)
	server := &Server{
		uomUseCase: uomUseCase,
		engine:     engine,
	}

	server.InitController()
	return server
}