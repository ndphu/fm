package service

import (
	"errors"
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/ndphu/fm/model"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type TokenService struct {
	secret string
}

func NewTokenService() *TokenService {
	return &TokenService{
		secret: "pfm-secret",
	}
}

func (s *TokenService) CreateToken(user *model.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user,
		"iss": "PFM Token Service",
		"exp": time.Now().Add(time.Hour * 1),
	})

	return token.SignedString([]byte(s.secret))
}

func (s *TokenService) ValidateToken(tokenString string) (*model.User, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(s.secret), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		sub := claims["sub"].(map[string]interface{})
		return &model.User{
			Id:         bson.ObjectId(sub["id"].(string)),
			FirstName:  sub["firstName"].(string),
			LastName:   sub["lastName"].(string),
			ExternalId: sub["externalId"].(string),
		}, nil

	} else {
		return nil, errors.New("token validate failed")
	}
}
