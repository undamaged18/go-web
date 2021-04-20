package jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
	"webapp/config"
)

type Claims struct {
	UserId string
	Data   string
	jwt.StandardClaims
}

var conf = config.New()
var jwtKey = []byte(conf.Keys.JWT)

func Create(user string, data string, duration string) (string, error) {
	if duration == "" {
		duration = "15m"
	}
	parsedDuration, err := time.ParseDuration(duration)
	if err != nil {
		return "", err
	}
	expTime := time.Now().Add(parsedDuration)

	claims := &Claims{
		UserId: user,
		Data:   data,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func Get(token string) (newToken string, id string, data string, err error) {
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return "", "", "", err
	}
	if !tkn.Valid {
		return "", "", "", errors.New("TOKEN_NOT_VALID")
	}
	if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) < 5*time.Minute {
		token, err = Create(claims.UserId, claims.Data, "60m")
		if err != nil {
			return "", "", "", err
		}
	}
	return token, claims.UserId, claims.Data, nil
}
