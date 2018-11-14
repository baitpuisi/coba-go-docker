package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Response is just a very basic example.
type Response struct {
	Status    string `json:"status,omitempty"`
	FileValue string `json:"FileValue,omitempty"`
}

// GetStatus returns always the same response.
func GetStatus(w http.ResponseWriter, _ *http.Request) {
	b := Response{Status: "idle"}
	json.NewEncoder(w).Encode(b)
}

func ReadFile(w http.ResponseWriter, _ *http.Request) {

	b, err := ioutil.ReadFile("file.txt") // just pass the file name
	if err != nil {
		fmt.Print(err)
	}

	str := string(b) // convert content to a 'string'

	data := Response{FileValue: str}
	json.NewEncoder(w).Encode(data)

}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/status", GetStatus).Methods("GET")
	router.HandleFunc("/readfile", ReadFile).Methods("GET")
	log.Fatal(http.ListenAndServe(":1234", router))
}
