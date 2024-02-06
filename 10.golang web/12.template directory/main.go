/*

- html template directory


ketika kita memiliki banyak file
nah kita bisa membuat template tersebut secara langsung berdasarkan
foldernya tanpa menyebutkan nama filenya

kita coba


*/

package main

import (
	_ "embed"
	"html/template"
	"net/http"
)

func main(){
	
	mux := http.NewServeMux()
	mux.HandleFunc("/",func(w http.ResponseWriter, r *http.Request) {


		t, err := template.ParseGlob("./views/*.html") 
		if err != nil{
			panic(err)
		}

		t.ExecuteTemplate(w,"index.html","halo selamat datang")

		/*
		
		jadi tinggal kita sebutin saja nama templatenya di
		executetemplate

		karena kan nama templatenya itu mengikuti nama filenya
		jadi kita bisa sebutkan nama templatenya pada executeTemplatenya

		jadi tanda bintang itu artinya apapun file yang extentionnya .html
		semua akan diparse ke dalam template 
		
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