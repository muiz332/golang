/*

- http cookie


http itu didesign stateless artinya dia tidak tahu request sebelumnya dan request yang
akan datang

nah cookie digunakan untuk informasi bahwa user telah melakukan request sebelumnya

contoh ketika user login, maka akan diberikan cookie sebagai penanda
bahwa user tersebut telah login

dan cookie tersebut akan digunakan setiap user mengirimkan request
dan server akan tahu bahwa user tersebut telah login

dan cookie tersebut akan disimpan diclient




*/

package main

import (
	"fmt"
	"net/http"
)




func handleSetCookie(w http.ResponseWriter, r *http.Request){

	cookie := new(http.Cookie) // membuat cookie

	cookie.Name = "login"
	cookie.Value = "sudah login"
	cookie.Path = "/"    // cookie nempel diurl apa
	cookie.MaxAge = 60 * 60 * 24 // 1 hari kadaluarsa

	// kalo / pathnya maka aktif disemua url 
	// karena defaultnya itu aktif diurl saat kita membuat cookie


	http.SetCookie(w,cookie) // untuk set cookie
	fmt.Fprintf(w,"hello")

}


func handleGetCookie(w http.ResponseWriter, r *http.Request){

	cookie, err := r.Cookie("login")
	if err != nil{
		fmt.Fprintf(w,"tidak ada cookie")
	}else{
		fmt.Fprintf(w,"cookie %s",cookie.Value)
	}
}

func main (){

	mux := http.NewServeMux()

	mux.HandleFunc("/",handleSetCookie)
	mux.HandleFunc("/about",handleGetCookie)

	server := http.Server{
		Addr: "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil{
		panic(err)
	}
}


/*

nah kalo untuk unitestnya cara mengambil cookie itu cukup dari 


cookies := result().Cookies()

for _, cookie := range cookies{
	cookie.Name dan cookie.Value
}


nah untuk mengirim cookie dari request kalian bisa buat cookie terlebih dahulu
seperti yang diatas kemudian 

req.AddCookie(cookie)

jadi itu cookie

*/