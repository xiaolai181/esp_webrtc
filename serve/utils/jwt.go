package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var JwtSecret = "1234lee"

var jwtSecert = []byte(JwtSecret)

type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

//生成token
func GenerateToken(username, password string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour) //过期时间
	claims := Claims{
		username,
		password,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "lee",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecert)
	return token, err
}

//解析token

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return jwtSecert, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
