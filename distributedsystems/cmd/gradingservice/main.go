package main

import (
	"context"
	"distributedsystems/grades"
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

	ctx, err := service.Start(
		context.Background(),
		gradingReg,
		host,
		port,
		grades.RegisterHandlers)
	if err != nil {
		stlog.Fatal(err)
	}

	// block until context is done
	<-ctx.Done()

	fmt.Println("Shutting Down Grading Service")
}
