package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/gorilla/mux"
)

func main() {

	fmt.Printf("Starting server at port 8080\n")
	route := mux.NewRouter()
	
	s := route.PathPrefix("/api").Subrouter()
	s.HandleFunc("/primes/{lower:[0-9]+},{upper:[0-9]+}", getPrimesHandler).Methods("GET")
	s.HandleFunc("/primeFactors/{n:[0-9]+}", getPrimeFactorsHandler).Methods("GET")
	s.HandleFunc("/numberOfDivisors/{n:[0-9]+}", getNumberOfDivisorsHandler).Methods("GET")
	s.Path("/metrics").Handler(promhttp.Handler())
	log.Println(http.ListenAndServe(":8080", s))

}
