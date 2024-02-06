/*

- send file embed

nah kalian bisa send file menggunakan golang embed
langsung saja kita coba

*/

package main

import (
	_ "embed"
	"net/http"
)

//go:embed views/index.html
var index []byte

//go:embed views/about.html
var about []byte

func main(){
	
	mux := http.NewServeMux()
	mux.HandleFunc("/",func(w http.ResponseWriter, r *http.Request) {
		w.Write(index)
	})
	mux.HandleFunc("/about",func(w http.ResponseWriter, r *http.Request) {
		w.Write(about)
	})


	server := http.Server{
		Addr: "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()

	if err != nil{
		panic(err)
	}


}