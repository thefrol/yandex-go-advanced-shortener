package main

import (
	"github.com/dsbasko/yandex-go-advanced-shortener/internal/handler"
	"net/http"
)

func main() {
	err := run()
	if err != nil {
		panic(err)
	}
}

func run() error {
	mux := handler.NewMux()
	return http.ListenAndServe("localhost:8080", mux)
}
