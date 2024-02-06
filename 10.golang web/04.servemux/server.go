/*

- serveMux

saat kita membuat web, kita biasanya ingin membuat lebih dari satu endpoint url
handlerFunc tidak mendukung hal tersebut

untungnya digolang ada serveMuxc yang digunakan untuk membuat lebih dari satu url
dimana serveMux kita bisa buat dari package http

jadi sebenarnya serveMux itu mengumpulkan lebih dari satu handlerFunc
yang dijadikan satu kesatuan

oke langsung saja kita coba




*/

package main

import (
	"fmt"
	"net/http"
)

func main() {
	url := "127.0.0.1:8080"

	// multiple endpoint url
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("ini adalah halaman utama"))
	})
	mux.HandleFunc("/about", func(w http.ResponseWriter, req *http.Request) {
		fmt.Println(req.Method)
		w.Write([]byte("ini adalah halaman about"))
	})

	server := http.Server{
		Addr:    url,
		Handler: mux,
	}
	fmt.Println(fmt.Sprintf("server is running on http://%s", url))
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}

}
