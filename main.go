package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	//log.Println("Vinayaka function invoked !!!!! ")
	lambda.Start(handler)
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	//log.Println("Vinayaka function invoked !!!!! ")

	var person Person
	err := json.Unmarshal([]byte(request.Body), &person)

	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	log.Println(" First Name is : ", *person.FirstName)
	log.Println(" Second Name is : ", *person.LastName)
	log.Println(" Second Name is : ", *person.Age)

	msg := fmt.Sprintf("Helloooo..... %v %v and age is : %v", *person.FirstName, *person.LastName, *person.Age)

	log.Println(" Response Message is: ", &msg)

	responseBody := ResponseBody{
		Message: &msg,
	}

	jbytes, err := json.Marshal(responseBody)

	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	response := events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(jbytes),
	}

	return response, nil
}

type Person struct {
	FirstName *string `json:"firstName"`
	LastName  *string `json:"lastName"`
	Age       *int64  `json:"age"`
}

type ResponseBody struct {
	Message *string `json:"message"`
}
