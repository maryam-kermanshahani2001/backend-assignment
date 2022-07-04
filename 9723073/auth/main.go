package auth

import (
	"fmt"
	jwt "github.com/golang-jwt/jwt/v4"
	"time"
)

var mySigningKey = []byte("mysupersecretphrase")

func GenerateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["user"] = "Elliot Forbes"
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		fmt.Errorf("something went wrong: %s", err.Error())
		return "", err
	}
	return tokenString, nil
}

func main() {
	tokenString, err := GenerateJWT()
	if err != nil {
		fmt.Println("Error generating token string")

	}
	fmt.Println(tokenString)

}
