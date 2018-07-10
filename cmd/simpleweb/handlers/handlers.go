package handlers

import (
	"github.com/gorilla/mux"
	"sync/atomic"
	"time"
	"log"
)

func Router() *mux.Router {
	isReady := &atomic.Value{}
	isReady.Store(false)

	go func() {
		log.Printf("Readyz probe is negative by default ...")
		time.Sleep(10 * time.Second)

		isReady.Store(true)
		log.Printf("Readyz probe is positive.")
	}()

	r := mux.NewRouter()
	r.HandleFunc("/home", home).Methods("GET")
	r.HandleFunc("/healthz", healthz)
	r.HandleFunc("/readyz", readyz(isReady))

	return r
}