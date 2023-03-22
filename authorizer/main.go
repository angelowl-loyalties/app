package main

import (
	"context"
	"errors"
	"log"
	"os"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/kms"
	"github.com/golang-jwt/jwt/v4"
	"github.com/matelang/jwt-go-aws-kms/v2/jwtkms"
)

var JWTKeyID string

var KMSClient *kms.Client

var KMSConfig *jwtkms.Config

type CustomJWTClaims struct {
	Role string `json:"role"`
	jwt.RegisteredClaims
}

func init() {
	JWTKeyID = os.Getenv("JWT_KMS_KEY_ID")
	if JWTKeyID == "" {
		log.Fatalln("JWS_KMS_KEY_ID environment variable is empty")
	}

	awsCfg, err := config.LoadDefaultConfig(context.Background(), config.WithRegion("ap-southeast-1"))
	if err != nil {
		log.Fatalln("config error: " + err.Error())
	}

	KMSClient = kms.NewFromConfig(awsCfg)
	KMSConfig = jwtkms.NewKMSConfig(KMSClient, JWTKeyID, false)
}

// this lambda takes the JWT authorization token that was forwarded from API gateway
// it parses the credentials and appends them to the response context to be consumed by the internal REST API
func handleRequest(ctx context.Context, event events.APIGatewayCustomAuthorizerRequest) (events.APIGatewayCustomAuthorizerResponse, error) {
	token := event.AuthorizationToken
	log.Println(token)

	tokenSlice := strings.Split(token, " ")
	var bearerToken string
	if len(tokenSlice) > 1 {
		bearerToken = tokenSlice[len(tokenSlice)-1]
	}
	if bearerToken == "" || strings.ToUpper(tokenSlice[0]) != "BEARER" {
		return events.APIGatewayCustomAuthorizerResponse{}, errors.New("not a bearer token")
	}

	claims := CustomJWTClaims{}
	_, err := jwt.ParseWithClaims(bearerToken, &claims, func(token *jwt.Token) (interface{}, error) {
		return KMSConfig, nil
	})
	if err != nil {
		log.Printf("cannot parse/verify token %s", err)
		return events.APIGatewayCustomAuthorizerResponse{}, errors.New("unauthorized")
	}

	principalID := claims.Subject // is the user's UUID
	role := claims.Role
	authContext := map[string]interface{}{ // given to backend to determine if resource accessed is valid
		"id":   principalID,
		"role": role,
	}

	// TODO: determine the method ARNs that should be allowed, make use of map for users to resources
	// arn:partition:execute-api:region:account-id:api-id/authorizers/authorizer-id
	// requestARN := event.MethodArn
	var resources []string

	if role == "user" {
		resources = []string{
			//"arn:aws:execute-api:ap-southeast-1:account-id:api-id/authorizers/authorizer-id",
			"*",
		}
	} else if role == "admin" {
		resources = []string{
			//"arn:aws:execute-api:ap-southeast-1:account-id:api-id/authorizers/authorizer-id",
			"*",
		}
	} else if role == "bank" {
		resources = []string{
			//"arn:aws:execute-api:ap-southeast-1:account-id:api-id/authorizers/authorizer-id",
			"*",
		}
	} else {
		log.Printf("invalid role")
		return events.APIGatewayCustomAuthorizerResponse{}, errors.New("invalid role")
	}

	return generatePolicy(principalID, "Allow", resources, authContext), nil
}

func generatePolicy(principalID string, effect string, resources []string, context map[string]interface{}) events.APIGatewayCustomAuthorizerResponse {
	authResponse := events.APIGatewayCustomAuthorizerResponse{PrincipalID: principalID}

	if effect != "" && len(resources) > 0 {
		authResponse.PolicyDocument = events.APIGatewayCustomAuthorizerPolicy{
			Version: "2012-10-17",
			Statement: []events.IAMPolicyStatement{
				{
					Action:   []string{"execute-api:Invoke"},
					Effect:   effect,
					Resource: resources,
				},
			},
		}
	}

	authResponse.Context = context
	return authResponse
}

func main() {
	lambda.Start(handleRequest)
}
