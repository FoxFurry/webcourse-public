package controller

import (
	"context"
	"github.com/foxfurry/university/webcourse/client/service"
	"github.com/gin-gonic/gin"
)

type IController interface {
	RegisterClientRoutes(r *gin.Engine)
}

type clientController struct {
	musicService service.IService
}

func NewClientController(ctx context.Context) IController {
	return &clientController{
		musicService: service.NewMusicService(ctx),
	}
}

func (ctrl *clientController) RegisterClientRoutes(r *gin.Engine){
	r.GET("/status", ctrl.Status)
}


func (ctrl *clientController) Status(c *gin.Context){
	var currentStatus = ctrl.musicService.Session()
	if currentStatus == nil {
		c.JSON(404, gin.H{"error": "current session is empty"})
	}else{
		c.JSON(200, currentStatus)
	}
}
