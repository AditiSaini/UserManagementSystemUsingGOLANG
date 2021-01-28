package helper

import (
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/twinj/uuid"
	"golang.org/x/crypto/bcrypt"

	Structure "../Structure"
)

func ValidateLogin(username string, password string) bool {
	profile := Show(username)
	if profile.Valid {
		//>>>>>>>Need to be modified since hash is stored in the db and not plain text<<<<<<<<
		hashed, _ := HashPassword(password)
		err := bcrypt.CompareHashAndPassword(hashed, []byte(profile.Password))
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
	td.AccessToken, _ = token.SignedString([]byte("secret"))
	return td, nil
}
