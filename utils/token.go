package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"os"
	"time"
)

const (
	accessTokenExpiration = 30 * time.Minute
)

type claims struct {
	Guid string `json:"guid"`
	*jwt.RegisteredClaims
}

func GenerateAccess(guid string) (string, error) {
	// inject this values to the jwt payload
	t := claims{
		guid,
		&jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(accessTokenExpiration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, t)

	secret := os.Getenv("JWT_SECRET_KEY")

	signedToken, err := token.SignedString([]byte(secret))

	if err != nil {
		return "no token", err
	}

	return "Bearer " + signedToken, nil
}

func GenerateRefresh() string {
	refresh := uuid.New()
	return refresh.String()
}
