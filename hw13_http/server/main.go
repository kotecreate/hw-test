package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
}

var addressFlag = flag.String("address", ":8080", "server address")

func handleRequest(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received %s request from %s for path %s", r.Method, r.RemoteAddr, r.URL.Path)

	var response Response

	switch r.Method {
	case "GET":
		response.Message = "You sent a GET!"
		w.WriteHeader(http.StatusOK)
	case "POST":
		response.Message = "You sent a POST!"
		w.WriteHeader(http.StatusCreated)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		response.Message = "Method not allowed!"
	}

	jsonResponse, _ := json.Marshal(response)
	w.Write(jsonResponse)
}

func main() {
	flag.Parse()
	http.HandleFunc("/", handleRequest)
	log.Printf("Server listening on %s", *addressFlag)
	http.ListenAndServe(*addressFlag, nil)
}
