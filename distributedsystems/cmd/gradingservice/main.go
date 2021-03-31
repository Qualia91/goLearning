package main

import (
	"context"
	"distributedsystems/grades"
	"distributedsystems/log"
	"distributedsystems/registry"
	"distributedsystems/service"
	"fmt"
	stlog "log"
)

func main() {

	host, port := "localhost", "6000"

	serviceAddress := fmt.Sprintf("http://%v:%v", host, port)

	var gradingReg registry.Registration
	gradingReg.ServiceName = registry.GradingService
	gradingReg.ServiceURL = serviceAddress
	gradingReg.RequiredServices = []registry.ServiceName{registry.LogService}
	gradingReg.ServiceUpdateURL = gradingReg.ServiceURL + "/services"

	ctx, err := service.Start(
		context.Background(),
		gradingReg,
		host,
		port,
		grades.RegisterHandlers)
	if err != nil {
		stlog.Fatal(err)
	}

	if logProvider, err := registry.GetProvider(registry.LogService); err == nil {
		fmt.Printf("Logging service was found at %v\n", logProvider)
		log.SetClientLogger(logProvider, gradingReg.ServiceName)
	}

	// block until context is done
	<-ctx.Done()

	fmt.Println("Shutting Down Grading Service")
}
