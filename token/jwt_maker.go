package token

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const minSercretKeySize = 32

type JWTMaker struct {
	secretKey string
}

var(
	ErrInvalidToken = errors.New("token is invalid")
	ErrExpiredToken = errors.New("token is expired")
)


func NewJWTMaker(secretKey string) (Maker, error) {
	if len(secretKey) < minSercretKeySize {
		return nil, ErrExpiredToken
	}
	return &JWTMaker{secretKey: secretKey}, nil
}

func (maker *JWTMaker) CreateToken(username string, duration time.Duration) (string, *Payload, error) {
	Payload, err := NewPayload(username, duration)
	if err != nil {
		return "", Payload, fmt.Errorf("error", err)
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodES256, Payload)
	token, err := jwtToken.SignedString([]byte(maker.secretKey))
	return token,Payload,err
}

func (maker *JWTMaker) VerifyToken(token string) (*Payload, error) {
		KeyFunc := func(token *jwt.Token)(interface{},error){
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok{
				return nil, ErrInvalidToken
			}
			return []byte(maker.secretKey), nil
	}	
		jwtToken23, err := jwt.ParseWithClaims(token, &Payload{},KeyFunc)
		if err != nil{
			 verr, ok := err.(*jwt.ValidationError)
			 if ok && errors.Is(verr.Inner, ErrExpiredToken){
				return nil, ErrExpiredToken
			 }
			 return nil, ErrInvalidToken
		}

		payload, ok := jwtToken23.Claims.(*Payload)

		if !ok{
			return nil, ErrInvalidToken
		}

		return payload, nil
}