package main

import (
	"httptest/router"
	"net/http"
)

func main() {
	url := "127.0.0.1:8080"
	server := http.Server{
		Addr:    url,
		Handler: router.Routers(),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
