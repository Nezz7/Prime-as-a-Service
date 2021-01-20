package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
)

var getPrimesCounter = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_request_get_primes_count", // metric name
		Help: "Number of GET /primes request.",
	},
	[]string{"status"}, // labels
)
var getPrimeFactors = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_request_get_prime_factors_count", // metric name
		Help: "Number of GET /prime-factors request.",
	},
	[]string{"status"}, // labels
)
var getNumberOfDivisors = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_request_get_number_of_divisors_count", // metric name
		Help: "Number of GET /number-of-divisors request.",
	},
	[]string{"status"}, // labels
)

func init() {
	// must register counter on init
	prometheus.MustRegister(getPrimesCounter)
	prometheus.MustRegister(getPrimeFactors)
	prometheus.MustRegister(getNumberOfDivisors)
}
func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
func getPrimesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET /primes")

	var status string
	defer func() {
		getPrimesCounter.WithLabelValues(status).Inc()
	}()

	vars := mux.Vars(r)
	lower, err := strconv.Atoi(vars["lower"])
	if err != nil || lower < 0 || lower > MAXN {
		status = "error"
		respondWithError(w, http.StatusBadRequest, "Invalid number LOWER")
		return
	}

	upper, er := strconv.Atoi(vars["upper"])
	if er != nil || upper < 0 || upper > MAXN {
		status = "error"
		respondWithError(w, http.StatusBadRequest, "Invalid number UPPER")
		return
	}

	result := primesInRange(lower, upper)
	status = "success"

	respondWithJSON(w, http.StatusOK, result)
}

func getPrimeFactorsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET /prime-factors")

	var status string
	defer func() {
		getPrimeFactors.WithLabelValues(status).Inc()
	}()

	vars := mux.Vars(r)
	n, err := strconv.Atoi(vars["n"])
	if err != nil || n < 0 || n > MAXN {
		status = "error"
		respondWithError(w, http.StatusBadRequest, "Invalid number N")
		return
	}

	result := primeFactors(n)
	status = "success"

	respondWithJSON(w, http.StatusOK, result)
}
func getNumberOfDivisorsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET /number-of-divisors")

	var status string
	defer func() {
		getNumberOfDivisors.WithLabelValues(status).Inc()
	}()

	vars := mux.Vars(r)
	n, err := strconv.Atoi(vars["n"])
	if err != nil || n > MAXN {
		status = "error"
		respondWithError(w, http.StatusBadRequest, "Invalid number N")
		return
	}

	result := numberOfDivisors(n)
	status = "success"
	respondWithJSON(w, http.StatusOK, result)
}
