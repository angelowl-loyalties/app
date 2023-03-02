package utils

import (
	"time"

	"github.com/cs301-itsa/project-2022-23t2-g1-t7/authorizer/models"

	"github.com/golang-jwt/jwt/v4"
	"github.com/matelang/jwt-go-aws-kms/v2/jwtkms"
)

type JWTClaims struct {
	Role string `json:"role"`
	jwt.RegisteredClaims
}

func CreateJWT(user models.User) *jwt.Token {
	claims := JWTClaims{
		Role: user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:   "angelowl.com",
			Subject:  user.Email,
			Audience: []string{"api.angelowl.com"},
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