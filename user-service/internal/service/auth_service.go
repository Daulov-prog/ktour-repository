package service

import (
	"errors"
	"time"
	"user-service/internal/storage"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	storage storage.UserStorageInterface
	secret  string
}

func NewAuthService(storage storage.UserStorageInterface, secret string) *AuthService {
	return &AuthService{
		storage: storage,
		secret:  secret,
	}
}

func (a *AuthService) Login(email string, password string) (string, error) {
	user, err := a.storage.GetUserByEmail(email)
	if err != nil {
		return "", errors.New("Пользовател не найден")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("Неверный пароль")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	})

	tokenString, err := token.SignedString([]byte(a.secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (a *AuthService) ValidateToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(a.secret), nil
	})
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("Невалидный токен")
	}

	userID := claims["user_id"].(string)
	return userID, nil
}
