package helper

import (
	"fmt"
	"net/http"
	"strings"

	Structure "../Structure"

	"github.com/dgrijalva/jwt-go"
)

func ExtractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	//normally Authorization the_token_xxx
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

func VerifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString := ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("secret"), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func ExtractTokenMetadata(r *http.Request) (*Structure.AccessDetails, error) {
	token, err := VerifyToken(r)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		accessUUID, ok := claims["access_uuid"].(string)
		if !ok {
			return nil, err
		}
		username := claims["username"].(string)
		if err != nil {
			return nil, err
		}
		return &Structure.AccessDetails{
			AccessUUID: accessUUID,
			Username:   username,
		}, nil
	}
	return nil, err
}
