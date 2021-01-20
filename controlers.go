package prime

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"

)



var getUsersCounter = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_request_get_users_count", // metric name
		Help: "Number of get_users request.",
	},
	[]string{"status"}, // labels
)
var postUserCounter = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_request_post_user_count", // metric name
		Help: "Number of post_user request.",
	},
	[]string{"status"}, // labels
)

func init() {
    // must register counter on init
	prometheus.MustRegister(getUsersCounter)
	prometheus.MustRegister(postUserCounter)
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

	vars := mux.Vars(r)
    lower, err := strconv.Atoi(vars["lower"])
    if err != nil || lower < 0  || lower >  MAXN {
        respondWithError(w, http.StatusBadRequest, "Invalid number LOWER")
        return
	}
	upper, er := strconv.Atoi(vars["upper"])
    if er != nil || upper < 0 || upper > MAXN {
        respondWithError(w, http.StatusBadRequest, "Invalid number UPPER")
        return
    }
	var status string
	defer func() {
		getUsersCounter.WithLabelValues(status).Inc()
	}()
	result := primesInRange(lower, upper)
	respondWithJSON(w, http.StatusOK, result)

}

func getPrimeFactorsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET /primeFactors")

	vars := mux.Vars(r)
    n, err := strconv.Atoi(vars["n"])
    if err != nil || n < 0 || n > MAXN {
        respondWithError(w, http.StatusBadRequest, "Invalid number N")
        return
    }
	var status string
	defer func() {
		getUsersCounter.WithLabelValues(status).Inc()
	}()
	result := primeFactors(n)
	respondWithJSON(w, http.StatusOK, result)

}
func getNumberOfDivisorsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET /numberOfDivisors")

	vars := mux.Vars(r)
    n, err := strconv.Atoi(vars["n"])
    if err != nil || n > MAXN {
        respondWithError(w, http.StatusBadRequest, "Invalid number N")
        return
    }
	var status string
	defer func() {
		getUsersCounter.WithLabelValues(status).Inc()
	}()
	result := numberOfDivisors(n)
	respondWithJSON(w, http.StatusOK, result)
	
}