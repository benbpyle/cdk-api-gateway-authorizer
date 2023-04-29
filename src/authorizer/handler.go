package main

import (
	"context"
	"fmt"
	"os"

	"github.com/lestrrat-go/jwx/jwt"
	log "github.com/sirupsen/logrus"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/lestrrat-go/jwx/jwk"
)

var (
	keySet jwk.Set
)

func handler(ctx context.Context, event events.APIGatewayCustomAuthorizerRequest) (events.APIGatewayCustomAuthorizerResponse, error) {
	bounds := len(event.AuthorizationToken)
	token := event.AuthorizationToken[7:bounds]
	parsedToken, err := jwt.Parse(
		[]byte(token),
		jwt.WithKeySet(keySet),
		jwt.WithValidate(true),
	)

	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Error("Error parsing the token")
		return events.APIGatewayCustomAuthorizerResponse{
			PrincipalID: "",
			PolicyDocument: events.APIGatewayCustomAuthorizerPolicy{
				Version: "2012-10-17",
				Statement: []events.IAMPolicyStatement{
					{
						Action:   []string{"execute-api:Invoke"},
						Effect:   "Deny",
						Resource: []string{"*"},
					},
				},
			},
			UsageIdentifierKey: "",
		}, nil
	}

	return events.APIGatewayCustomAuthorizerResponse{
		PrincipalID: "",
		PolicyDocument: events.APIGatewayCustomAuthorizerPolicy{
			Version: "2012-10-17",
			Statement: []events.IAMPolicyStatement{
				{
					Action:   []string{"execute-api:Invoke"},
					Effect:   "Allow",
					Resource: []string{"*"},
				},
			},
		},
		Context:            DumpClaims(parsedToken),
		UsageIdentifierKey: "",
	}, nil
}

func init() {
	log.SetFormatter(&log.JSONFormatter{
		PrettyPrint: false,
	})

	log.SetLevel(log.DebugLevel)

	region := "us-west-2"
	poolId := os.Getenv("USER_POOL_ID")
	var err error

	jwksUrl := fmt.Sprintf("https://cognito-idp.%s.amazonaws.com/%s/.well-known/jwks.json", region, poolId)
	keySet, err = jwk.Fetch(context.TODO(), jwksUrl)

	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
			"url":   jwksUrl,
		}).Fatal("error getting keyset")
	}
}

func main() {
	lambda.Start(handler)
}
