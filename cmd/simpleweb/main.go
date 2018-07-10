package main

import (
	"net/http"
	"log"
	"os"

	"github.com/cloud-jeong/go-startup/cmd/simpleweb/handlers"
	"os/signal"
	"syscall"
	"context"
)

func main() {
	log.Print("Starting the service ...")

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("Port is not set.")
	}

	r := handlers.Router()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, os.Kill, syscall.SIGTERM)

	srv := &http.Server {
		Addr : ":" + port,
		Handler: r,
	}

	go func() {
		log.Fatal(srv.ListenAndServe())
	}()
	log.Print("The service is ready to listen and service.")

	killSignal := <-interrupt

	switch killSignal {
	case os.Kill:
		log.Print("Got SIGKILL ...")
	case os.Interrupt:
		log.Print("Got SIGINT ...")
	case syscall.SIGTERM:
		log.Print("Got SIGTERM ...")
	}

	log.Print("The service is shutting down ...")
	srv.Shutdown(context.Background())

	log.Print("Done")

}


