package token

/*
buat/generate token
ganti package token jd package main
*/
import (
	"fmt"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
)

type M map[string]interface{}
type MyClaims struct {
	jwt.StandardClaims
	Username string
	Email    string
	Group    string
}

var APPLICATION_NAME = "My Simple JWT App"
var LOGIN_EXPIRATION_DURATION = time.Duration(1) * time.Hour
var JWT_SIGNING_METHOD = jwt.SigningMethodHS256
var JWT_SIGNATURE_KEY = []byte("Private Key")

func main() {
	// generate token
	claims := MyClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    APPLICATION_NAME,
			ExpiresAt: time.Now().Add(LOGIN_EXPIRATION_DURATION).Unix(),
		},
		// data yg dimasukkan
		Username: "username",
		Email:    "email",
		Group:    "group",
	}

	token := jwt.NewWithClaims(
		JWT_SIGNING_METHOD,
		claims,
	)

	signedToken, err := token.SignedString(JWT_SIGNATURE_KEY)
	if err != nil {
		panic(err)
	}

	// tokenString, _ := json.Marshal(M{"token": signedToken})
	fmt.Println("token jwt ", signedToken)
	// fmt.Println("tidak tahu ", tokenString)
}
