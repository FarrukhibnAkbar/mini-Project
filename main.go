package main

import(
	"net/http"
	"github.com/gorilla/mux"
)


func main() {
	
	r := mux.NewRouter()

	r.HandleFunc("/admin", GetAdmin).Methods("GET")
	r.HandleFunc("/admin", PostAdmin).Methods("POST")

	http.ListenAndServe(":8080", r)
}