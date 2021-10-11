package controller

import (
	"github.com/foxfurry/university/webcourse/host/model"
	"github.com/gin-gonic/gin"
)

type IController interface {
	RegisterHostRoutes(r *gin.Engine)
}

type hostController struct {}

func NewHostController() IController {
	return &hostController{}
}

func (ctrl *hostController) RegisterHostRoutes(r *gin.Engine){
	r.GET("/next", ctrl.Next)
}


func (ctrl *hostController) Next(c *gin.Context){
	nextSong := model.GetRandomSong()

	c.JSON(200, &nextSong)
}
