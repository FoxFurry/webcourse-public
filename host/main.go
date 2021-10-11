package main

import (
	"context"
	"github.com/foxfurry/university/webcourse/host/application"
	"os"
	"os/signal"
	"syscall"
)

func main(){
	isDone := make(chan os.Signal)
	signal.Notify(isDone, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	ctx, cancel := context.WithCancel(context.Background())

	mainApp := application.NewApp()
	go mainApp.Start(ctx)

	<-isDone
	cancel()
	mainApp.Shutdown()
}