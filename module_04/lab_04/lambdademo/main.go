package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(Run)

}

func Run(event events.SQSEvent) {
	fmt.Println("Starting")
	for _, message := range event.Records {
		fmt.Println("Here is the SQS message: ", message.Body)
	}
}
