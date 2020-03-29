package jwt

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
)

//IssueToken ...
func IssueToken(id string) (string, error) {
	type ConsumerInfo struct {
		ID string `json:"id,omitempty"`
		jwt.StandardClaims
	}

	expire := time.Now().Add(time.Second * 86400).Unix()

	// Create the Claims
	claims := &ConsumerInfo{
		id,
		jwt.StandardClaims{
			Issuer:    "Tin",
			ExpiresAt: expire,
		},
	}
	return SignToken(claims)
}

//SignToken sign claims
func SignToken(claims jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	ss, err := token.SignedString(signKey)

	if nil != err {
		log.Println("Error while signing the token")
		log.Printf("Error signing token: %v\n", err)
		return ss, err
	}

	return ss, nil
}

//VerificationToken ...
func VerificationToken(tokenString string) (string, error) {
	type ConsumerInfo struct {
		ID string `json:"id,omitempty"`
		jwt.StandardClaims
	}

	// Parse the token
	token, err := jwt.ParseWithClaims(tokenString, &ConsumerInfo{}, func(token *jwt.Token) (interface{}, error) {
		// since we only use the one private key to sign the tokens,
		// we also only use its public counter part to verify
		return verifyKey, nil
	})

	if err != nil {
		return "", err
	}
	claims, ok := token.Claims.(*ConsumerInfo)

	if !ok {
		return "", errors.New("invalid token")
	}
	fmt.Println("Claims id: ", claims.ID)
	return claims.ID, err
}
