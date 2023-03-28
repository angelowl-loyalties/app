package utils

import (
	"context"
	"log"
	"time"

	"github.com/cs301-itsa/project-2022-23t2-g1-t7/profiler/models"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/kms"
	"github.com/golang-jwt/jwt/v4"
	"github.com/matelang/jwt-go-aws-kms/v2/jwtkms"
)

var KMSClient *kms.Client
var KMSConfig *jwtkms.Config

func InitKMS(AK string, SK string, JWTKeyID string) {
	awsCfg, err := config.LoadDefaultConfig(context.Background(), config.WithRegion("ap-southeast-1"))
	if err != nil {
		log.Fatalln("config error: " + err.Error())
	}

	KMSClient = kms.NewFromConfig(awsCfg)
	KMSConfig = jwtkms.NewKMSConfig(KMSClient, JWTKeyID, false)
	log.Println("configured KMS client")
}

type CustomJWTClaims struct {
	Role  string `json:"role"`
	IsNew bool   `json:"is_new"`
	jwt.RegisteredClaims
}

func CreateJWT(user *models.User) *jwt.Token {
	claims := CustomJWTClaims{
		Role:  user.Role,
		IsNew: user.IsNew,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:   "itsag1t2.com",
			Subject:  user.ID.String(),
			Audience: []string{"itsag1t2.com"},
			//ExpiresAt: jwt.NewNumericDate(time.Now().Add(20 * time.Minute)),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			NotBefore: jwt.NewNumericDate(time.Now()),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			//ID:        "",
		},
	}

	jwtToken := jwt.NewWithClaims(jwtkms.SigningMethodRS512, claims)

	return jwtToken
}
