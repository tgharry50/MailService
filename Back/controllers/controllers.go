package controllers

import (
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fs := http.FileServer(http.Dir("public"))
	fs.ServeHTTP(w, r)
}

func ImageHandler(w http.ResponseWriter, r *http.Request) {
	http.StripPrefix("/api/image/", http.FileServer(http.Dir("public/image"))).ServeHTTP(w, r)
}

func OtherHandler(w http.ResponseWriter, r *http.Request) {
	http.StripPrefix("/api/other/", http.FileServer(http.Dir("public/other"))).ServeHTTP(w, r)
}

func Routes() {
	http.HandleFunc("/api/image/", ImageHandler) // Image
	http.HandleFunc("/api/other/", OtherHandler) // Other
	http.HandleFunc("/", IndexHandler)           // Front
}
