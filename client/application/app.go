package application

import (
	"context"
	"github.com/foxfurry/university/webcourse/client/controller"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type IApp interface {
	Start()
	Shutdown()
}

type clientApp struct {
	server *http.Server
}

func NewApp(ctx context.Context) IApp {
	router := gin.New()

	ctrl := controller.NewClientController(ctx)
	ctrl.RegisterClientRoutes(router)

	app := clientApp{
		server: &http.Server{
			Addr: ":8080",
			Handler: router,
		},
	}

	return &app
}

func (a *clientApp) Start(){
	log.Println("Starting client server")
	if err := a.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Error while running server: %v", err)
	}
}

func (a *clientApp) Shutdown(){
	log.Println("Shutting down client server")
	if err := a.server.Shutdown(context.Background()); err != nil {
		log.Fatalf("Unable to shutdown client server: %v", err)
	}
}
