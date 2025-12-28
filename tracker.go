package main

import (
	"log"
	"net/http"
	"tracker/controllers"
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
)

type RouteFunc func(w http.ResponseWriter, req *http.Request)

func PrepRoute(f RouteFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		f(w, r)
	}
}




func main() {
	r := mux.NewRouter()

	// Configure CORS options
	// Where ORIGIN_ALLOWED is like `scheme://dns[:port]`, or `*` (insecure)
	headersOk := handlers.AllowedHeaders([]string{"Content-Type"})
	originsOk := handlers.AllowedOrigins([]string{"https://danielberry.tech"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "POST"})



	r.HandleFunc("/stat", PrepRoute(controllers.Stat)).Methods("POST")


	fs := http.FileServer(http.Dir("static/"))
	r.Handle("/static/js/{$}", http.StripPrefix("/static/", fs)).Methods("GET")
	r.Handle("/static/html/{$}", http.StripPrefix("/static/", fs)).Methods("GET")


	log.Fatal(http.ListenAndServe(":3000", handlers.CORS(originsOk, headersOk, methodsOk)(r)))

	//log.Fatal(http.ListenAndServe(":3000", r))
	//log.Fatal(http.ListenAndServeTLS(":443", "certs/cert1.pem", "certs/privkey1.pem", r))
}
