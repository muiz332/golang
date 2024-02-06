/*

- html template embed


jadi ketika nanti kalian membuat golang web
maka direkomendasikan menggunakan template embed ini ya

kita coba

*/

package main

import (
	"embed"
	"html/template"
	"net/http"
)

//go:embed views/*.html
var templates embed.FS

func main(){
	
	mux := http.NewServeMux()
	mux.HandleFunc("/",func(w http.ResponseWriter, r *http.Request) {


		t, err := template.ParseFS(templates,"views/*.html") 
		if err != nil{
			panic(err)
		}

		t.ExecuteTemplate(w,"index.html","halo selamat datang")

		/*
		
		kalian tinggal gunakan saja method ParseFS dan maukkan embednya
		pada parameter pertama
		
		*/

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