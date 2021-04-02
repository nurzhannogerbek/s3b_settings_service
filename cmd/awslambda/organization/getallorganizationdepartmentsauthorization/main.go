package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
)

func handleRequest(in interface{}) (interface{}, error) {
	fmt.Println(in)
	return true, nil
}

func main() {
	lambda.Start(handleRequest)
}
