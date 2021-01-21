package helper

import (
	"fmt"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/shaj13/go-guardian/auth"
	"golang.org/x/crypto/bcrypt"
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

func CreateToken(username string) string {
	//Token session for 15 min
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Minute * 15).Unix(),
	})
	jwtToken, _ := token.SignedString([]byte("secret"))
	return jwtToken
}

func VerifyToken(tokenString string) (auth.Info, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("secret"), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		user := auth.NewDefaultUser(claims["sub"].(string), "", nil, nil)
		return user, nil
	}
	return nil, fmt.Errorf("Invalid token")
}
