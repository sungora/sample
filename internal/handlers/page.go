package handlers

import (
	"net/http"
)

// MainPage главная страница
func PageMain(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("MainPage"))
}

// PageApi страница api
func PageApi(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("PageApi"))
}
