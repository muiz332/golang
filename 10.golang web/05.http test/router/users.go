package router

import (
	"httptest/controllers"
	"net/http"
)

func Users(mux *http.ServeMux) {

	mux.HandleFunc("/users", controllers.GetAllUsers)
}
