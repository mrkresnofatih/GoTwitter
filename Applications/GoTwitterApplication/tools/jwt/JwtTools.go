package jwt

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"log"
	"time"
)

const applicationJwtSecret = "WDvhlFvg1L70crAumeLM0muYg783TTWRwdVFonOsjbrpL3PFoTnAcrv5ew7B"
const applicationJwtClaimsKeyExpiresAt = "expiresAt"
const ApplicationJwtClaimsKeyUsername = "username"

type IJwtTokenBuilder interface {
	GetClaims() *jwt.MapClaims
	initialize()
	Build() (string, error)
}

type BasicJwtTokenBuilder struct {
	Claims       jwt.MapClaims
	ExpiresAfter time.Duration
}

func (s *BasicJwtTokenBuilder) initialize() {
	s.Claims = jwt.MapClaims{}
	s.Claims[applicationJwtClaimsKeyExpiresAt] = time.Now().Add(s.ExpiresAfter).Format(time.RFC3339)
}

func (s *BasicJwtTokenBuilder) Build() (string, error) {
	s.initialize()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, s.Claims)
	tokenString, err := token.SignedString([]byte(applicationJwtSecret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (s *BasicJwtTokenBuilder) GetClaims() *jwt.MapClaims {
	return &(s.Claims)
}

type UsernameJwtTokenBuilder struct {
	JwtTokenBuilder IJwtTokenBuilder
	Username        string
}

func (s *UsernameJwtTokenBuilder) initialize() {
	s.JwtTokenBuilder.initialize()
	claims := *(s.JwtTokenBuilder.GetClaims())
	claims[ApplicationJwtClaimsKeyUsername] = s.Username
}

func (s *UsernameJwtTokenBuilder) Build() (string, error) {
	s.initialize()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, s.JwtTokenBuilder.GetClaims())
	tokenString, err := token.SignedString([]byte(applicationJwtSecret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (s *UsernameJwtTokenBuilder) GetClaims() *jwt.MapClaims {
	return s.JwtTokenBuilder.GetClaims()
}

func GetValidityFromToken(tokenString string) bool {
	token, err := jwt.Parse(tokenString, func(tkn *jwt.Token) (interface{}, error) {
		if _, ok := tkn.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Unexpected_Signing_Method")
		}
		return []byte(applicationJwtSecret), nil
	})
	if err != nil {
		return false
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		claimValue := claims[applicationJwtClaimsKeyExpiresAt]
		if claimValue == nil {
			return false
		}

		tokenExpireTime, err := time.Parse(time.RFC3339, claimValue.(string))
		if err != nil {
			log.Println("error here")
			return false
		}
		expired := time.Now().After(tokenExpireTime)
		if expired {
			log.Println("token is expired!")
		}
		return !expired
	}
	return false
}

func GetClaimFromToken[T any](tokenString string, claimType string) (T, error) {
	token, err := jwt.Parse(tokenString, func(tkn *jwt.Token) (interface{}, error) {
		if _, ok := tkn.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected Signing Method")
		}
		return []byte(applicationJwtSecret), nil
	})
	if err != nil {
		return *new(T), err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		claimValue := claims[claimType]
		if claimValue != nil {
			return claimValue.(T), nil
		}
		return *new(T), errors.New("claim Not Found")
	}
	return *new(T), errors.New("token Not Valid")
}
