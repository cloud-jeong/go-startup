package main

import (
	"fmt"
	"net/http"
	"log"
)

func main() {
	log.Print("Starting the service ...")
	http.HandleFunc("/home", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprint(w, "/home called")
	},
	)

	log.Print("The service is ready to listen and service.")
	log.Fatal(http.ListenAndServe(":8000", nil))
}