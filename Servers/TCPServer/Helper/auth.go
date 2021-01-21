package helper

import (
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/twinj/uuid"
	"golang.org/x/crypto/bcrypt"

	Structure "../Structure"
)

func ValidateLogin(username string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(password), []byte("12345"))
	if err != nil {
		log.Println(err)
		return false
	}
	if username == "Aditi" {
		return true
	}
	return false
}

func CreateToken(username string) (*Structure.TokenDetails, error) {
	td := &Structure.TokenDetails{}
	td.AtExpires = time.Now().Add(time.Minute * 15).Unix()
	td.AccessUuid = uuid.NewV4().String()
	td.RtExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
	td.RefreshUuid = uuid.NewV4().String()

	//Craeting access token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"authorised":  true,
		"access_uuid": td.AccessUuid,
		"username":    username,
		"exp":         td.AtExpires,
	})
	td.AccessToken, _ = token.SignedString([]byte("secret"))

	//Creating refresh token
	refresh_token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"refresh_uuid": td.RefreshUuid,
		"username":     username,
		"exp":          td.RtExpires,
	})
	td.RefreshToken, _ = refresh_token.SignedString([]byte("secret"))

	return td, nil
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
