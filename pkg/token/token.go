package token

import (
	"campus-recruiting-api/configs"
	"github.com/dgrijalva/jwt-go"
)

type UserClaims struct {
	Id           int32  `json:"id"`
	NickName     string `json:"username"`
	HeadPortrait string `json:"headPortrait"`
	Sex          int32  `json:"sex"`
	jwt.StandardClaims
}

func MakeToken(obj UserClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, obj)
	tokenString, err := token.SignedString([]byte(configs.JwtSigningKey))
	return tokenString, err
}
func ParseToken(token string) (*UserClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(configs.JwtSigningKey), nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*UserClaims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
