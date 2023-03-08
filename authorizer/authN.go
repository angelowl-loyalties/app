package main

import (
	"context"
	"log"
	"os"

	"github.com/cs301-itsa/project-2022-23t2-g1-t7/authorizer/models"
	"github.com/cs301-itsa/project-2022-23t2-g1-t7/authorizer/utils"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/kms"
	"github.com/matelang/jwt-go-aws-kms/v2/jwtkms"
)

var authnJWTKeyID string

var authnKMSClient *kms.Client

var authnKMSConfig *jwtkms.Config

func init() {
	authnJWTKeyID = os.Getenv("JWT_KMS_KEY_ID")
	if authnJWTKeyID == "" {
		log.Fatalln("JWT_KMS_KEY_ID environment variable is empty")
	}

	awsCfg, err := config.LoadDefaultConfig(context.Background(), config.WithRegion("ap-southeast-1"))
	if err != nil {
		log.Fatalln("config error: " + err.Error())
	}

	authnKMSClient = kms.NewFromConfig(awsCfg)
	authnKMSConfig = jwtkms.NewKMSConfig(authnKMSClient, authnJWTKeyID, false)
}

// this lambda takes result from step function that starts with /auth/user in profiler
// uses the successful credentials, then creates a JWT to return via the API gateway
func handleRequest(ctx context.Context, request models.User) (*models.AuthNResponse, error) {
	jwtToken := utils.CreateJWT(request)

	signedJWT, err := jwtToken.SignedString(authnKMSConfig.WithContext(context.Background()))
	if err != nil {
		log.Fatalln("failed to sign JWT: ", err)
		return nil, err
	}

	response := &models.AuthNResponse{
		JWT:    signedJWT,
		UserID: request.ID.String(),
	}

	return response, nil
}

func main() { //nolint:all
	lambda.Start(handleRequest)
}
