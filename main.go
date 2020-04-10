package main

import(

  "github.com/aws/aws-lambda-go/lambda"
  "github.com/sirupsen/logrus"

  "github.com/shekodn/exploring-streaming-data-analysis/handlers"
)

// Create a new instance of the logger.
var log = logrus.New()

func main() {
  log.Print("Starting the service...")

  lambda.Start(handlers.LambdaHandler)

  log.Print("Done")
}
