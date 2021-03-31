package main

import (
	"context"
	"distributedsystems/log"
	"distributedsystems/registry"
	"distributedsystems/service"
	"fmt"
	stlog "log"
)

func main() {
	log.Run("./app.log")

	host, port := "localhost", "4000"

	serviceAddress := fmt.Sprintf("http://%v:%v", host, port)

	var logReg registry.Registration
	logReg.ServiceName = registry.LogService
	logReg.ServiceURL = serviceAddress
	logReg.RequiredServices = make([]registry.ServiceName, 0)
	logReg.ServiceUpdateURL = logReg.ServiceURL + "/services"

	ctx, err := service.Start(
		context.Background(),
		logReg,
		host,
		port,
		log.RegisterHandlers)
	if err != nil {
		stlog.Fatal(err)
	}

	// block until context is done
	<-ctx.Done()

	fmt.Println("Shutting Down Log Service")
}
