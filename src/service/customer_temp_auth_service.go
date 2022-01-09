package service

import (
	"aquiladb/src/model"
	"aquiladb/src/repository"
	"errors"
	"math/rand"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
)

var (
	// ADJECTIVES ...
	ADJECTIVES = []string{"autumn", "hidden", "bitter", "misty", "silent", "empty", "dry", "dark", "summer",
		"icy", "delicate", "quiet", "white", "cool", "spring", "winter", "patient",
		"twilight", "dawn", "crimson", "wispy", "weathered", "blue", "billowing",
		"broken", "cold", "damp", "falling", "frosty", "green", "long", "late", "lingering",
		"bold", "little", "morning", "muddy", "old", "red", "rough", "still", "small",
		"sparkling", "throbbing", "shy", "wandering", "withered", "wild", "black",
		"young", "holy", "solitary", "fragrant", "aged", "snowy", "proud", "floral",
		"restless", "divine", "polished", "ancient", "purple", "lively", "nameless"}

	// NOUNS ...
	NOUNS = []string{"waterfall", "river", "breeze", "moon", "rain", "wind", "sea", "morning",
		"snow", "lake", "sunset", "pine", "shadow", "leaf", "dawn", "glitter", "forest",
		"hill", "cloud", "meadow", "sun", "glade", "bird", "brook", "butterfly",
		"bush", "dew", "dust", "field", "fire", "flower", "firefly", "feather", "grass",
		"haze", "mountain", "night", "pond", "darkness", "snowflake", "silence",
		"sound", "sky", "shape", "surf", "thunder", "violet", "water", "wildflower",
		"wave", "water", "resonance", "sun", "wood", "dream", "cherry", "tree", "fog",
		"frost", "voice", "paper", "frog", "smoke", "star"}
)

type TokenClaimsForTempCustomer struct {
	jwt.StandardClaims
	CustomerUuid string `json:"customer_uuid"`
	IsPermanent  bool   `json:"is_permanent"`
}

type CustomerTempAuth struct {
	repo repository.CustomerTempAuthRepositoryInterface
}

func NewCustomerTempAuthService(repo repository.CustomerTempAuthRepositoryInterface) *CustomerTempAuth {
	return &CustomerTempAuth{
		repo: repo,
	}
}

func (c CustomerTempAuth) CreateTempCustomer() (model.CustomerTemp, error) {

	var customer model.CustomerTemp

	randomAdjective := ADJECTIVES[rand.Intn(len(ADJECTIVES))]
	randomNoun := NOUNS[rand.Intn(len(NOUNS))]

	customer.FirstName = strings.Title(randomAdjective)
	customer.LastName = strings.Title(randomNoun)
	customer.SecretKey = KeyGenerate(15)

	createdCustomer, err := c.repo.RegisterTempCustomer(customer)
	if err != nil {
		return customer, err
	}

	token, err := GenerateTokenForTempCustomer(*createdCustomer)
	if err != nil {
		return customer, err
	}

	// ???
	createdCustomer.Token = token

	return *createdCustomer, err
}

func KeyGenerate(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	var seededRand *rand.Rand = rand.New(
		rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, length)

	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func GenerateTokenForTempCustomer(customer model.CustomerTemp) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &TokenClaimsForTempCustomer{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTlForTemporaryUser).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		customer.CustomerId,
		customer.IsPermanent,
	})

	return token.SignedString([]byte(signingKey))
}

func ParseTokenPermanentCustomer(accessToken string) (TokenClaimsForTempCustomer, error) {

	var newTokenClaims TokenClaimsForTempCustomer

	token, err := jwt.ParseWithClaims(accessToken, &TokenClaimsForTempCustomer{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Invalid token.")
		}

		return []byte(signingKey), nil
	})
	if err != nil {
		return newTokenClaims, err
	}

	claims, ok := token.Claims.(*TokenClaimsForTempCustomer)
	if !ok {
		return newTokenClaims, errors.New("token claims are not of type *tokenClaims")
	}

	return *claims, nil
}
