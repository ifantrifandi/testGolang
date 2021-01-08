package auth

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
)

type Service interface {
	GenerateToken (userId int) (string, error)
	ValidateToken (token string) (*jwt.Token, error)
}

type jwtService struct {}

var SECRET_KEY = []byte("TESTSECRET_sikirit_kiyi")

func NewService() *jwtService{
	return &jwtService{}
}

// claim == payload
func (s *jwtService) GenerateToken (userID int) (string, error) {
	claim := jwt.MapClaims{}
	claim["user_id"] = userID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	signedToken, err := token.SignedString(SECRET_KEY)

	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}


func (s *jwtService) ValidateToken(token string) (*jwt.Token, error) {

	tokenParse , err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error){
		_,ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("Not Authorize")
		}

		return []byte(SECRET_KEY), nil
	})

	if err != nil {
		return tokenParse, err
	}

	return tokenParse, nil
}