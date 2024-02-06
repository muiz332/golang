/*

- html template


jadi digolang kita juga bisa memasukkan html template
atau logic dan data dinamis

langsung saja kita coba


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

		dataHtml := "<p>{{.}}</p>"

		t, err := template.New("ininamatemplate").Parse(dataHtml)
		if err != nil{
			panic(err)
		}

		t.ExecuteTemplate(w,"ininamatemplate","ini data dinamisnya")


		/*
		
		jadi kalian harus buat templatenya sekaligus menamai templatenya dengan 
		menggunakan template.New() dan kemudian kita parse
		
		*/
	})



	mux.HandleFunc("/about",func(w http.ResponseWriter, r *http.Request) {
		/*
		
		nah kita juga bisa menggunakan template pada html file dan nama template akan tergantung
		dengan nama filenya

		akan tetapi kita harus menamai html yang menggunakan template dengan nama
		about.gohtml

		kita cobq
		
		*/

		t, err := template.ParseFiles("./views/about.gohtml")
		if err != nil{
			panic(err)
		}

		t.ExecuteTemplate(w,"about.gohtml","ini html template")

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