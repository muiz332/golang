package router

import "net/http"

func Routers() *http.ServeMux {

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("selamat datang"))
	})
	Users(mux)

	return mux
}
