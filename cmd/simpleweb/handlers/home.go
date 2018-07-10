package handlers

import (
	"net/http"
	"fmt"
)

func home(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(w, "/home: Hello! Your request was processed.")
}