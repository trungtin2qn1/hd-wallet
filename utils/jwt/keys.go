package jwt

import (
	"crypto/rsa"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/dgrijalva/jwt-go"
)

var (
	verifyKey *rsa.PublicKey
	signKey   *rsa.PrivateKey
)

//LoadRSAKeys read the key files before starting http handlers
func LoadRSAKeys() {
	signBytes, err := ioutil.ReadFile("./keys/key.rsa")
	if nil != err {
		fmt.Println("Can't read private key")
		log.Fatal(err)
	}

	signKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if nil != err {
		fmt.Println("Can't parse private key")
		log.Fatal(err)
	}

	verifyBytes, err := ioutil.ReadFile("./keys/key.rsa.pub")
	if nil != err {
		fmt.Println("Can't read public key")
		log.Fatal(err)
	}

	verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	if nil != err {
		fmt.Println("Can't parse public key")
		log.Fatal(err)
	}
}
