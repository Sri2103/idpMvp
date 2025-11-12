package service

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// claims info
// IDPClaims defines the structure for all JWTs used in IDP
type IDPClaims struct {
	UserID   string `json:"user_id"`
	TenantID string `json:"tenant_id"`
	IsAdmin  bool   `json:"is_admin"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

type JwtService struct {
	secretKey       []byte
	issuer          string
	accessTokenTTL  time.Duration
	refreshTokenTTL time.Duration
}

// initiate new jwt service
func NewJWTService() *JwtService {
	secret := os.Getenv("JWT_SECRET")

	if secret == "" {
		secret = "supersecretkey"
	}

	issuer := os.Getenv("JWT_ISSUER")
	if issuer == "" {
		issuer = "idp"
	}

	return &JwtService{
		secretKey:       []byte(secret),
		issuer:          issuer,
		accessTokenTTL:  15 * time.Minute,
		refreshTokenTTL: 7 * 24 * time.Hour,
	}
}

// generate tokens
func (s *JwtService) GenerateTokens(userId, tenantId, username string, isAdmin bool) (accessToken, refreshToken string, err error) {
	accessClaims := IDPClaims{
		UserID:   userId,
		TenantID: tenantId,
		IsAdmin:  isAdmin,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(s.accessTokenTTL)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    s.issuer,
			Subject:   userId,
		},
	}

	access := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)

	accessToken, err = access.SignedString(s.secretKey)
	if err != nil {
		return "", "", err
	}

	// Refresh Token
	refreshClaims := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(s.refreshTokenTTL)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		Issuer:    s.issuer,
		Subject:   userId,
	}

	refresh := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshToken, err = refresh.SignedString(s.secretKey)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

// validate token
func (s *JwtService) ValidateToken(tokenString string) (*IDPClaims, error) {
	// parse the token
	token, err := jwt.ParseWithClaims(tokenString, IDPClaims{}, func(t *jwt.Token) (any, error) {
		return s.secretKey, nil
	})
	if err != nil {
		return nil, err
	}

	// retrieve and checking validity of the token and fetch claims
	if claims, ok := token.Claims.(*IDPClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}

// parse and validate refresh token with minimal claims
func (s *JwtService) ValidateRefreshToken(refreshTokenString string) (*jwt.RegisteredClaims, error) {
	token, err := jwt.ParseWithClaims(refreshTokenString, jwt.RegisteredClaims{}, func(t *jwt.Token) (any, error) {
		return s.secretKey, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*jwt.RegisteredClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("refresh token invalid")
}
