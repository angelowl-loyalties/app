package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/cs301-itsa/project-2022-23t2-g1-t7/authorizer/models"
	"github.com/cs301-itsa/project-2022-23t2-g1-t7/authorizer/utils"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/kms"
	"github.com/matelang/jwt-go-aws-kms/v2/jwtkms"
)

var keyID string

var kmsConfig *jwtkms.Config

func init() {
	keyID = os.Getenv("JWT_KMS_KEY_ID")
	if keyID == "" {
		log.Fatalln("JWT_KMS_KEY_ID environment variable is empty")
	}

	awsCfg, err := config.LoadDefaultConfig(context.Background(), config.WithRegion("ap-southeast-1"))
	if err != nil {
		log.Fatalln("config error: " + err.Error())
	}

	kmsConfig = jwtkms.NewKMSConfig(kms.NewFromConfig(awsCfg), keyID, false)
	// TODO: setup JWE
}

// this lambda takes result from step function that starts with /auth/user in profiler
// uses the successful credentials, then creates a JWT to return via the API gateway
func handleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	testUserJSON := "{\n    \"id\": \"eef42ade-5c56-470c-a485-7573cc23dafc\",\n    \"first_name\": \"justin\",\n    \"last_name\": \"lam\",\n    \"phone\": \"+6597924661\",\n    \"email\": \"justin.100600@gmail.com\",\n    \"Role\": \"user\",\n    \"CreditCards\": null,\n    \"CreatedAt\": \"2023-01-26T02:51:39.15073+08:00\",\n    \"UpdatedAt\": \"2023-01-26T21:06:45.612496+08:00\"\n}\n"

	var user models.User
	err := json.Unmarshal([]byte(testUserJSON), &user)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	jwtToken := utils.CreateJWT(user)
	signedJWT, err := jwtToken.SignedString(kmsConfig.WithContext(context.Background()))
	if err != nil {
		log.Fatalln("failed to sign JWT: ", err)
	}

	resp := &models.AuthNResponse{
		JWT:    signedJWT,
		UserID: user.ID.String(),
	}

	body, err := json.Marshal(resp)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	return events.APIGatewayProxyResponse{Body: string(body), StatusCode: 200}, nil
}

func main() {
	lambda.Start(handleRequest)
}
