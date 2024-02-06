package main

import (
	"fmt"
	"net/http"
)

func main() {

	var hendler http.HandlerFunc = func(w http.ResponseWriter, req *http.Request) {
		// w.Write([]byte("hello world")) // bisa gini

		fmt.Fprint(w, "hello world")
	}

	url := "127.0.0.1:8080"
	server := http.Server{
		Addr:    url,
		Handler: hendler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
	fmt.Println("Sever is running on http://" + url)
}
