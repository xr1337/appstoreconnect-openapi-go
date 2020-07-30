package main

import (
	"context"
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"time"

	api "github.com/xr1337/appstoreconnect-openapi-go/generated"

	"github.com/dgrijalva/jwt-go"
)

// Variables from the appstore
// from https://developer.apple.com/documentation/appstoreconnectapi/generating_tokens_for_api_requests
var authKeyP8FilePath = "<path to p8 file>"
var iss = "<Issuer ID>"
var kid = "<Key identifier>"

func main() {
	expireTime := time.Now().Add(time.Minute * 10).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"iss": iss,
		"exp": expireTime,
		"aud": "appstoreconnect-v1",
	})
	token.Header["kid"] = kid

	key, err := getPrivateKey(authKeyP8FilePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	signedToken, err := token.SignedString(key)
	if err != nil {
		fmt.Println(err)
		return
	}

	cfg := api.NewConfiguration()
	// cfg.Debug = false
	auth := context.WithValue(context.Background(), api.ContextAccessToken, signedToken)
	client := api.NewAPIClient(cfg)
	var response api.UsersResponse
	if response, _, err = client.UsersApi.UsersGetCollection(auth, nil); err != nil {
		fmt.Println(err)
		return
	}
	for _, user := range response.Data {
		fmt.Println(user.Attributes.Username)
	}
}

// Reads a p8 file and returns the ecdsa private key
func getPrivateKey(fileP8 string) (*ecdsa.PrivateKey, error) {
	var fileData []byte
	var err error
	if fileData, err = ioutil.ReadFile(fileP8); err != nil {
		return nil, err
	}
	var parsedKey interface{}
	var key *ecdsa.PrivateKey
	var ok bool
	block, _ := pem.Decode(fileData)
	if parsedKey, err = x509.ParsePKCS8PrivateKey(block.Bytes); err != nil {
		return nil, err
	}
	if key, ok = parsedKey.(*ecdsa.PrivateKey); !ok {
		return nil, fmt.Errorf("Not a EC private key file")
	}
	return key, nil
}
