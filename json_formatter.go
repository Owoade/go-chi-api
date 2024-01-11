package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respond_json(http_writter http.ResponseWriter, code int, payload interface{}) {

	data, err := json.Marshal(payload)

	if err != nil {

		log.Printf("Failed to Marshal payload %v", payload)

		http_writter.WriteHeader(500)

		return

	}

	http_writter.Header().Add("Content-Type", "application/json")
	http_writter.WriteHeader(code)
	http_writter.Write(data)

}

func respond_err(http_writter http.ResponseWriter, code int, msg string) {

	type ErrorResponse struct {
		Error string `json:"error"`
	}

	respond_json(http_writter, code, ErrorResponse{
		Error: msg,
	})

}
