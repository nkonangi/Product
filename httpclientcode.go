package main

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/lambda"
)

type Request struct {
	Name string `json:"name"`
}

type Response struct {
	Message string `json:"message"`
}

func handler(request Request) (Response, error) {
	reqBody, _ := json.Marshal(request)
	resp, err := http.Post("http://example.com/api", "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		return Response{}, err
	}
	defer resp.Body.Close()

	var response Response
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return Response{}, err
	}

	return response, nil
}

func main() {
	lambda.Start(handler)
}
