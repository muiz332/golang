package controllers

import "net/http"

var GetAllUsers http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("semua data users"))
}
var GetSingleUser http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("satu data user"))
}
