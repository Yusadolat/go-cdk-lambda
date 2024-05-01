package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
)

type myEvent struct {
	Username string `json:"username"`
}

func HandleRequest(event myEvent) (string, error) {
	if event.Username == "" {
		return "", fmt.Errorf("username is empty")
	}
	return fmt.Sprintf("Hello %s!", event.Username), nil
}
func main() {
	lambda.Start(HandleRequest)
}
