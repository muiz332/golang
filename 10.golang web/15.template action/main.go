/*

- html template data


kadang kita juga ingin memasukkan control flow didalam htmlnya seperti
percabangan maupun perulangan

nah ini bisa kita lakukan
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

		t.ExecuteTemplate(w,"index.html",map[string]any{
			"Title" : "home",
			"Name" : "muiz",
			"Alamat" : "banyuwangi",
		})

	})



	server := http.Server{
		Addr: "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()

	if err != nil{
		panic(err)
	}


/*

      nah untuk pengkondisiannya itu agak berbeda, untuk operatornya kalian tuliskan setelah if
      dan bentuk operatornya itu adalah tulisan bukan simbol

          {{if eq .Alamat "banyuwangi"}}

      contohnya kyk == menjadi eq
      atau > menjadi gt

      dan lain lain

      tetapi kalo hanya saya tulis .Alamat artinya dia akan mengecheck .Alamat != ""

*/

}