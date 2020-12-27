package main

import (
	"net/http"

	"github.com/go-fx-practice/httphandler"
)

func main() {
	mux := http.NewServeMux()
	httphandler.New(mux)

	http.ListenAndServe(":8080", mux)
}
