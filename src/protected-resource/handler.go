package main

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"os"

	log "github.com/sirupsen/logrus"
)

var isLocal bool

func handler(ctx context.Context, event events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {

	success := &Response{
		Message:   "Congrats! A Payload",
		CustomKey: event.RequestContext.Authorizer["customKey"].(string),
	}

	b, _ := json.Marshal(success)
	return &events.APIGatewayProxyResponse{
		Body:       string(b),
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}, nil

}

func main() {
	lambda.Start(handler)
}

func init() {
	isLocal, _ = strconv.ParseBool(os.Getenv("IS_LOCAL"))

	log.SetFormatter(&log.JSONFormatter{
		PrettyPrint: isLocal,
	})
}
