package auth

import (
	"fmt"
	"time"

	"github.com/backend-ids/src/schema/models"
	"github.com/dgrijalva/jwt-go"
)

type JWTManager struct {
	secretkey     string
	tokenDuration time.Duration
}

type UserClaims struct {
	jwt.StandardClaims
	UserName string `json:"username"`
	Role     string `json:"role"`
}

func NewJWTManager(secretkey string, tokenDuration time.Duration) *JWTManager {
	return &JWTManager{secretkey, tokenDuration}
}

func (manager *JWTManager) Generate(u *models.User) (string, error) {
	claims := UserClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(manager.tokenDuration).Unix(),
		},
		UserName: u.Name,
		Role:     u.Role,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(manager.secretkey))
}

func (manager *JWTManager) Verify(accessToken string) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(
		accessToken,
		&UserClaims{},
		func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, fmt.Errorf("unexpected token signing method")
			}
			return []byte(manager.secretkey), nil
		},
	)
	if err != nil {
		return nil, fmt.Errorf("invalid token: %w", err)
	}
	claims, ok := token.Claims.(*UserClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}
	return claims, nil
}
