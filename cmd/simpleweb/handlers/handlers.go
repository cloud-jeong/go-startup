package handlers

import (
	"github.com/gorilla/mux"
	"sync/atomic"
	"time"
	"log"
	"net/http"
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
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./techmag/")))
	//r.HandleFunc("/", home).Methods("GET")
	r.HandleFunc("/healthz", healthz)
	r.HandleFunc("/readyz", readyz(isReady))

	return r
}