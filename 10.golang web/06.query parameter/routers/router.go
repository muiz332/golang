package routers

import (
	"fmt"
	"net/http"
)

func Router() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {

		nama := r.URL.Query().Get("nama") // cara mendapatkan query string

		if nama == "" {
			fmt.Fprintf(w, "hello")
		} else {
			fmt.Fprintf(w, "hello %s", nama)
		}

	})

	return mux
}
