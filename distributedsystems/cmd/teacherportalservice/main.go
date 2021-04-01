package main

import (
	"context"
	"distributedsystems/log"
	"distributedsystems/registry"
	"distributedsystems/service"
	"distributedsystems/teacherportal"
	"fmt"
	stlog "log"
)

func main() {

	err := teacherportal.ImportTemplares()
	if err != nil {
		stlog.Fatal(err)
	}

	host, port := "localhost", "5000"

	serviceAddress := fmt.Sprintf("http://%v:%v", host, port)

	var teacherReg registry.Registration
	teacherReg.ServiceName = registry.TeacherPortal
	teacherReg.ServiceURL = serviceAddress
	teacherReg.RequiredServices = []registry.ServiceName{registry.LogService, registry.GradingService}
	teacherReg.ServiceUpdateURL = teacherReg.ServiceURL + "/services"
	teacherReg.HeartbeatURL = teacherReg.ServiceURL + "/heartbeat"

	ctx, err := service.Start(
		context.Background(),
		teacherReg,
		host,
		port,
		teacherportal.RegisterHandlers)
	if err != nil {
		stlog.Fatal(err)
	}

	if logProvider, err := registry.GetProvider(registry.LogService); err == nil {
		log.SetClientLogger(logProvider, teacherReg.ServiceName)
	}

	// block until context is done
	<-ctx.Done()

	fmt.Println("Shutting Down Teacher Portal")
}
