package application

import (
	"context"
	"github.com/foxfurry/university/webcourse/host/controller"
	"github.com/foxfurry/university/webcourse/host/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type IApp interface {
	Start(ctx context.Context)
	Shutdown()
}

type hostApp struct {
	server *http.Server
	service service.IService
}

func NewApp() IApp {
	router := gin.New()

	ctrl := controller.NewHostController()
	ctrl.RegisterHostRoutes(router)


	app := hostApp{
		server: &http.Server{
			Addr: ":8081",
			Handler: router,
		},
		service: service.NewListenerService(),
	}

	return &app
}

func (a *hostApp) Start(ctx context.Context){
	a.service.Start(ctx)

	log.Println("Starting host server")
	if err := a.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Error while running host: %v", err)
	}
}

func (a *hostApp) Shutdown(){
	log.Println("Shutting down host server")
	if err := a.server.Shutdown(context.Background()); err != nil {
		log.Fatalf("Unable to shutdown host server: %v", err)
	}
}
