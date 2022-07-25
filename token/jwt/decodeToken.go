package jwt

/*
ganti dg package main
decode token jwt
*/
import (
	"fmt"

	jwt "github.com/golang-jwt/jwt/v4"
)

var JWT_SIGNING_METHOD = jwt.SigningMethodHS256
var JWT_SIGNATURE_KEY = []byte("Private Key")

func main() {
	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTgzMjA4NTcsImlzcyI6Ik15IFNpbXBsZSBKV1QgQXBwIiwiVXNlcm5hbWUiOiJ1c2VybmFtZSIsIkVtYWlsIjoiZW1haWwiLCJHcm91cCI6Imdyb3VwIn0.Ughm4yfyd_zEWS1peTAAtpE20pJfSijh2jcc8F8c8wE"
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Signing method invalid")
		} else if method != JWT_SIGNING_METHOD {
			return nil, fmt.Errorf("Signing method invalid")
		}

		return JWT_SIGNATURE_KEY, nil
	})
	if err != nil {
		panic(err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		panic("salah, token tidak valid")
	}
	fmt.Println(claims)
	fmt.Println(token)
	fmt.Println(claims["Email"])
	fmt.Println(claims["exp"])
}
