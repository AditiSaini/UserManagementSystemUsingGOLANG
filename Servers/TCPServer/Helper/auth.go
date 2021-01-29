package helper

import (
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/twinj/uuid"
	"golang.org/x/crypto/bcrypt"

	Structure "servers/TCPServer/Structure"
	Constants "servers/internal"
)

func ValidateLogin(username string, password string) bool {
	profile := Show(username)
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
