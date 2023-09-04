package api

import (
	"go-laundry/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UomController struct {
	uomUseCase usecase.UomUseCase
	engine *gin.Engine
}

func (controller *UomController) getListUomsHandler(ctx *gin.Context)  {
	
	uoms, err := controller.uomUseCase.FindAll()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, uoms)
}


func NewUomController(uomUC usecase.UomUseCase, engine *gin.Engine) {
	controller := UomController{
		uomUseCase: uomUC,
		engine:     engine,
	}

	rg := engine.Group("/api/v1")
	rg.GET("uoms", controller.getListUomsHandler)
}