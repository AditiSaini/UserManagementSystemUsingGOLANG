package authentication

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/twinj/uuid"
	"golang.org/x/crypto/bcrypt"

	MySQL "servers/MySQL"
	Structure "servers/Structure"
	Constants "servers/internal"
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
		return []byte(Constants.TOKEN_SECRET), nil
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

//Check whether the token has expired
func TokenValid(r *http.Request) error {
	token, err := VerifyToken(r)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return err
	}
	return nil
}

func ValidateLogin(username string, password string) bool {
	profile := MySQL.Show(username)
	if profile.Valid {
		err := bcrypt.CompareHashAndPassword(profile.Password, []byte(password))
		if err != nil {
			log.Println(err)
			return false
		}
		return true
	}
	return false
}

func CreateToken(username string) (*Structure.TokenDetails, error) {
	td := &Structure.TokenDetails{}
	td.AtExpires = time.Now().Add(time.Minute * 30).Unix()
	td.AccessUuid = uuid.NewV4().String()
	td.RtExpires = time.Now().Add(time.Hour * 24 * 7).Unix()

	//Craeting access token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"authorised":  true,
		"access_uuid": td.AccessUuid,
		"username":    username,
		"exp":         td.AtExpires,
	})
	td.AccessToken, _ = token.SignedString([]byte(Constants.TOKEN_SECRET))
	return td, nil
}
