package middleware

import (
	"github.com/douyin/cmd/user/global"
	"github.com/douyin/pkg/errno"

	"github.com/dgrijalva/jwt-go"
)

type JWT struct {
	SigningKey []byte
}

type CustomClaims struct {
	Id          int64
	AuthorityId int64
	jwt.StandardClaims
}

func NewJWT() *JWT {
	return &JWT{
		[]byte(global.JWTSetting.Secret),
	}
}

func (j *JWT) CreateToken(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//zap.S().Debugf(token.SigningString())
	return token.SignedString(j.SigningKey)

}

func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, errno.TokenMalformedErr
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, errno.TokenExpiredErr
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, errno.TokenNotValidYetErr
			} else {
				return nil, errno.TokenInvalidErr
			}

		}
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errno.TokenInvalidErr
}
