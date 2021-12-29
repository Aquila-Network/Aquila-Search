package service

import (
	"aquiladb/src/config"
	"aquiladb/src/model"
	"aquiladb/src/repository"
	"crypto/sha1"
	"errors"
	"fmt"
	"regexp"
	"time"

	"github.com/golang-jwt/jwt"
)

var (
	envConfig  = config.InitEnvConfig()
	salt       = envConfig.Hash.Salt
	signingKey = envConfig.Hash.SignKey
	tokenTTL   = 24 * time.Hour
)

type TokenClaims struct {
	jwt.StandardClaims
	UserId   int    `json:"user_id"`
	Email    string `json:"email"`
	IsAdmin  bool   `json:"is_admin"`
	IsActive bool   `json:"is_active"`
}

type AuthService interface {
	RegisterService() string
	LoginService() string
}

type authService struct {
	repo repository.AuthRepositoryInterface
}

func NewAuthService(repo repository.AuthRepositoryInterface) *authService {
	return &authService{
		repo: repo,
	}
}

func (a *authService) Register(user model.User) (string, error) {

	user.Password = generatePasswordHash(user.Password)

	if len(user.FirstName) < 1 {
		return "", errors.New("First Name is required 'first_name'.")
	}

	if len(user.LastName) < 1 {
		return "", errors.New("Last Name is required 'last_name'.")
	}

	if len(user.Email) < 1 {
		return "", errors.New("Email is required 'email'.")
	}

	if isEmailValid(user.Email) == false {
		return "", errors.New("Email is not valid.")
	}

	if len(user.Password) < 1 {
		return "", errors.New("Password is required 'password'.")
	}

	createdUser, err := a.repo.Register(user)
	if err != nil {
		return "", err
	}

	createdUser.IsActive = true
	token, err := GenerateToken(createdUser)
	if err != nil {
		return "", err
	}

	return token, err
}

func (a *authService) Login(user model.LoginUser) (string, error) {

	if len(user.Email) < 1 {
		return "", errors.New("Email is required 'email'.")
	}

	if isEmailValid(user.Email) == false {
		return "", errors.New("Email is not valid.")
	}

	if len(user.Password) < 1 {
		return "", errors.New("Password is required 'password'.")
	}

	userExist, err := a.repo.Login(user.Email, generatePasswordHash(user.Password))
	if err != nil {
		return "", err
	}

	token, err := GenerateToken(userExist)
	if err != nil {
		return "", err
	}

	return token, nil
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

func GenerateToken(user model.User) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &TokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		int(user.Id),
		user.Email,
		user.IsAdmin,
		user.IsActive,
	})

	return token.SignedString([]byte(signingKey))
}

func ParseToken(accessToken string) (TokenClaims, error) {

	var newTokenClaims TokenClaims

	token, err := jwt.ParseWithClaims(accessToken, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})
	if err != nil {
		return newTokenClaims, err
	}

	claims, ok := token.Claims.(*TokenClaims)
	if !ok {
		return newTokenClaims, errors.New("token claims are not of type *tokenClaims")
	}

	return *claims, nil
}

// remove this after find out how to correctly validate
func isEmailValid(e string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(e)
}
