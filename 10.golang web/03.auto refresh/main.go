package main

import "net/http"

func main() {

	var handle http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("selamat datang kuy"))
	}

	url := "127.0.0.1:8080"
	server := http.Server{
		Addr:    url,
		Handler: handle,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}

	/*

		jadi untuk melakukan auto reload ketika ada perubahan kalian cukup
		gunakan package external yaitu

		go install github.com/cosmtrek/air@latest

		lalu kalian ketikkan air init
		dan jalankan menggunakan perintah

		air

		maka ketika ada perubahan secara otomatis server kita akan direlod
		jadi seperti itu ya

	*/

}
