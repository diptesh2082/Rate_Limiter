package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/didip/tollbooth"
)

type Message struct {
	Status string `json:"status"`
	Body   string `json:"body"`
}

func endpointHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	message := Message{
		Status: "Successful",
		Body:   "Hi! you have using the api",
	}
	err := json.NewEncoder(writer).Encode(&message)
	if err != nil {
		return
	}
}

func main() {
	Message := Message{
		Status: "Request Failled",
		Body:   "You are already in your capacity",
	}
	jsonMessage, _ := json.Marshal(Message)
	tlbthLimiter := tollbooth.NewLimiter(1, nil)
	tlbthLimiter.SetMessageContentType("application/json")
	tlbthLimiter.SetMessage(string(jsonMessage))
	http.Handle("/ping", tollbooth.LimitFuncHandler(tlbthLimiter, endpointHandler))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Println("There is a error listining to the port 8080", err)
	}
}