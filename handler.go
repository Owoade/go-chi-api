package main

import "net/http"

func error_handler(http_writter http.ResponseWriter, http_request *http.Request) {
	respond_err(http_writter, 400, "This is an error handler")
}

func handler(http_writter http.ResponseWriter, http_request *http.Request) {

	type Respnse struct {
		Message string `json:"message"`
	}

	respond_json(http_writter, 200, Respnse{
		Message: "Hello, this is a successful response",
	})
}
