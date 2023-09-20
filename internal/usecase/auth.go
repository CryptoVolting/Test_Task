package usecase

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"testProject/internal/entity"
	"testProject/internal/usecase/repository"
	"time"
)

const (
	salt       = "hjqrhjqw12gk7ajfhajs"
	signingKey = "qrkjk#4#%3ja#4353KSFjHrtth"
	tokenTTL   = 12 * time.Hour
)

type AuthUsecase struct {
	authorization repository.Authorization
}

type tokenClaims struct {
	jwt.StandardClaims
	Admin *bool `json:"admin"`
}

func NewAuthUsecase(authorization repository.Authorization) *AuthUsecase {
	return &AuthUsecase{authorization: authorization}
}

func (s *AuthUsecase) CreateUser(user entity.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.authorization.CreateUser(user)

}

func (s *AuthUsecase) GenerateToken(username, password string) (string, error) {
	user, err := s.authorization.GetUser(username, generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Admin,
	})
	return token.SignedString([]byte(signingKey))
}

func (s *AuthUsecase) ParseToken(accessToken string) (*bool, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return nil, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.Admin, nil
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
