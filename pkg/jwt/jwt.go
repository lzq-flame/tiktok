package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

/**
 * @Description
 * @Author 拥抱漏风
 * @Date 2022/5/6 22:19
 **/

var jwtKey = []byte("a_secret_crect")

type Claims struct {
	UserId int64
	jwt.StandardClaims
}

func ReleaseToken(id int64) (string, error) {
	expirationTime := time.Now().Add(7 * 24 * time.Hour)
	claims := &Claims{
		UserId: id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "ngxs.site",
			Subject:   "user token",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return " ", err
	}
	return tokenString, err
}

func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (i interface{}, err error) {
		return jwtKey, nil
	})

	return token, claims, err
}
