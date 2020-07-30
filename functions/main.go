package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	// files, err := ioutil.ReadDir(".")
	// if err != nil {
	// 	return events.APIGatewayProxyResponse{
	// 		StatusCode: 500,
	// 		Body:       err.Error(),
	// 	}, err
	// }

	// buf := bytes.NewBuffer(nil)

	// for _, file := range files {
	// 	buf.WriteString(file.Name() + "\n")
	// }

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "Hello AWS Lambda and Venmurasu\n",
	}, nil
}

func main() {
	// Make the handler available for Remote Procedure Call by AWS Lambda
	lambda.Start(handler)
}
