package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

func handleName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprintf(w, "Hello, %v!", vars["PARAM"])
}

func handleBad(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
}

func handleData(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(fmt.Errorf("Can't read body from POST request: %w", err))
	}
	fmt.Fprintf(w, "I got message:\n%v", string(data))
}

func handleHeaders(w http.ResponseWriter, r *http.Request) {
	a := r.Header.Get("a")
	b := r.Header.Get("b")
	Aa, err := strconv.Atoi(a)
	if err != nil {
		log.Fatal(fmt.Errorf("Can't parse header 'a' to int : %w", err))
	}
	Ab, err := strconv.Atoi(b)
	if err != nil {
		log.Fatal(fmt.Errorf("Can't parse header 'b' to int : %w", err))
	}
	w.Header().Set("a+b", strconv.Itoa(Aa + Ab))
}

/**
Please note Start functions is a placeholder for you to start your own solution.
Feel free to drop gorilla.mux if you want and use any other solution available.

main function reads host/port from env just for an example, flavor it following your taste
*/

// Start /** Starts the web server listener on given host and port.
func Start(host string, port int) {
	router := mux.NewRouter()
	router.HandleFunc("/name/{PARAM}", handleName).Methods("GET")
	router.HandleFunc("/bad", handleBad).Methods("GET")
	router.HandleFunc("/data", handleData).Methods("POST")
	router.HandleFunc("/headers", handleHeaders).Methods("POST")

	log.Println(fmt.Printf("Starting API server on %s:%d\n", host, port))
	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), router); err != nil {
		log.Fatal(err)
	}
}

//main /** starts program, gets HOST:PORT param and calls Start func.
func main() {
	host := os.Getenv("HOST")
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 8081
	}
	Start(host, port)
}
