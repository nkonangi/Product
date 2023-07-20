package main

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
	Name string `json:"name"`
}

func handler(ctx context.Context, sfEvent events.StepFunctionsEvent) error {
	var payload MyEvent
	if err := json.Unmarshal(sfEvent.Payload, &payload); err != nil {
		return err
	}
	
	// Do something with payload
	println(payload.Name)

	return nil
}

func main() {
	lambda.Start(handler)
}
