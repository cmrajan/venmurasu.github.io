package main

import (
	"bytes"
	"fmt"
	"io/ioutil"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/blevesearch/bleve"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	//	mapping := bleve.NewIndexMapping()
	index, err := bleve.OpenUsing("example.bleve", map[string]interface{}{
		"read_only": true,
	})
	if err != nil {
		fmt.Println(err)

	}

	// index, err := bleve.New("example.bleve", mapping)
	// if err != nil {
	// 	fmt.Println(err)

	// }

	// data := struct {
	// 	Name string
	// }{
	// 	Name: "text",
	// }

	// // index some data
	// index.Index("id", data)

	// search for some text
	query := bleve.NewMatchQuery("text")
	search := bleve.NewSearchRequest(query)
	searchResults, err := index.Search(search)
	if err != nil {
		fmt.Println(err)

	}
	fmt.Println(searchResults)

	files, err := ioutil.ReadDir(".")
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       err.Error(),
		}, err
	}

	buf := bytes.NewBuffer(nil)

	for _, file := range files {
		buf.WriteString(file.Name() + "\n----search results----------\n" + fmt.Sprintln(searchResults) + "\n----search results----------\n")
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "Hello AWS Lambda and search\n" + buf.String(),
	}, nil
}

func main() {
	// Make the handler available for Remote Procedure Call by AWS Lambda
	lambda.Start(handler)

}
