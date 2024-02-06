/*

- html template data


kadang kita juga ingin mengirimkan banyak data ke dalam htmlnya
nah ini bisa kita lakukan jika data yang kita kirimkan
itu berbentuk struct atau map

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
			"Nama" : "muiz",
		})

		/*

		jadi kita bisa gunakan map
		
		*/

	})

	
	type Add struct{
		Jalan string
	}

	type Page struct{
		Title string
		Nama string
		Alamat Add
	}

	mux.HandleFunc("/about",func(w http.ResponseWriter, r *http.Request) {





		t, err := template.ParseFS(templates,"views/*.html") 
		if err != nil{
			panic(err)
		}

		t.ExecuteTemplate(w,"about.html",Page{
			Title : "home",
			Nama : "muiz",
			Alamat: Add{
				Jalan : "belum ada",
			},
		})

		/*

		dan juga kita bisa gunakan struct
		
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