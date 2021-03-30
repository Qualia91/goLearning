package main

import (
	"context"
	"distributedsystems/registry"
	"fmt"
	"log"
	"net/http"
)

func main() {

	// register handle
	http.Handle("/services", &registry.RegistryService{})

	// create context with a cancel and defer the cancel
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var srv http.Server
	srv.Addr = registry.ServerPort

	go func() {
		log.Println(srv.ListenAndServe())
		cancel()
	}()

	go func() {
		fmt.Println("Registry service started. Press any key to stop.")
		var s string
		fmt.Scanln(&s)
		srv.Shutdown(ctx)
		cancel()
	}()

	<-ctx.Done()
	fmt.Println("Shutting down registry service")
}
