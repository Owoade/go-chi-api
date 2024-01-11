package main

import (
	"encoding/json"
	"net/http"
)

func (api_config *ApiConfig) create_author_handler(http_writter http.ResponseWriter, http_request *http.Request) {

	type CreateAuthParams struct {
		Name  string `json:name`
		Bio   string `json:string`
		Email string `json:email`
	}

	decoder := json.NewDecoder(http_request.Body)

	params := CreateAuthParams{}

	err := decoder.Decode(&params)

	if err != nil {
		respond_err(http_writter, 400, "Error parsing JSON")
		return
	}

	user, err := api_config.DB.	(http_request, CreateAuthParams{
		Name:  params.Name,
		Bio:   params.Bio,
		Email: params.Email,
	})

	if err != nil {
		respond_err(http_writter, 400, "Error creating user")
		return
	}

	respond_json(http_writter, 201, user)

}
