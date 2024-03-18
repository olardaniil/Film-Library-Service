package usecase

import (
	"crypto/sha1"
	"errors"
	"film-library-service/internal/repository"
	"fmt"
	"github.com/golang-jwt/jwt"
	"time"
)

const (
	signingKey = "qrkjk#4#%35FSFJlja#4353KSFjH"
	tokenTTL   = 12 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	UserID int `json:"user_id"`
}

type AuthUsecase struct {
	repo repository.Authorization
}

func NewAuthUsecase(repo repository.Authorization) *AuthUsecase {
	return &AuthUsecase{repo: repo}
}

func (uc *AuthUsecase) GenerateToken(username string, password string) (string, error) {
	user, err := uc.repo.GetUser(username, hashPassword(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.ID,
	})

	return token.SignedString([]byte(signingKey))
}

func (uc *AuthUsecase) ParseToken(accessToken string) (int, error) {

	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserID, nil
}

func hashPassword(password string) string {
	sha := sha1.New()
	sha.Write([]byte(password))
	shaHash := sha.Sum(nil)
	shaHashString := fmt.Sprintf("%x", shaHash)

	return shaHashString
}
