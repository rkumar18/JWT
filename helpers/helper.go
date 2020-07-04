package helpers

import (
	"github.com/dgrijalva/jwt-go"
	"encoding/json"
	"net/http"
	"time"
	"fmt"
	
	
)

type Claims struct{
	Email string
	jwt.StandardClaims
}

var jwtkey = []byte("password")


func WriteResponse(write http.ResponseWriter, response interface{}, statusCode int){
	write.Header().Set("Content-Type", "application/json")
	write.WriteHeader(statusCode)
	json.NewEncoder(write).Encode(response)
}



func Generatetoken(str string ) string {
expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		Email: str ,
		StandardClaims: jwt.StandardClaims{
		ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtkey)
	if err != nil {
		fmt.Println("Error occured")
	}
	return tokenString
}

func Validatetoken(token string) bool {
	claims := &Claims{}
	checktoken, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token)(interface{}, error){
		return jwtkey , nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return false
		}
		return false
	}
	if !checktoken.Valid {
		return false
	}
	return true
}
