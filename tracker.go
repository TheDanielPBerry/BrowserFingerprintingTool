package main

import (
	"log"
	"net/http"
	"tracker/controllers"
	"github.com/gorilla/mux"
)

type RouteFunc func(w http.ResponseWriter, req *http.Request)

func PrepRoute(f RouteFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		f(w, r)
	}
}



func main() {
	r := mux.NewRouter()


	r.HandleFunc("/stat", PrepRoute(controllers.Stat)).Methods("POST")


	fs := http.FileServer(http.Dir("static/"))
	r.Handle("/static/js/{$}", http.StripPrefix("/static/", fs)).Methods("GET")
	r.Handle("/static/html/{$}", http.StripPrefix("/static/", fs)).Methods("GET")


	log.Fatal(http.ListenAndServe(":3000", r))
	//log.Fatal(http.ListenAndServeTLS(":443", "certs/cert1.pem", "certs/privkey1.pem", r))
}
