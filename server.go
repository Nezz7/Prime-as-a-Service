package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {

	fmt.Printf("Starting server at port 8080\n")
	route := mux.NewRouter()

	s := route.PathPrefix("").Subrouter()
	s.HandleFunc("/primes/{lower:[0-9]+}&{upper:[0-9]+}", getPrimesHandler).Methods("GET")
	s.HandleFunc("/prime-factors/{n:[0-9]+}", getPrimeFactorsHandler).Methods("GET")
	s.HandleFunc("/number-of-divisors/{n:[0-9]+}", getNumberOfDivisorsHandler).Methods("GET")
	s.Path("/metrics").Handler(promhttp.Handler())
	
	log.Println(http.ListenAndServe(":8080", s))

}
