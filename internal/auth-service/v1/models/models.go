package models

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type User struct {
	ID           string
	Username     string
	Email        string
	PasswordHash string
	IsAdmin      bool
	Roles        []string
	IsActive     bool
	CreatedAt    time.Time
	UpdatedAt    time.Time
	LastLogin    time.Time
	TenantID     string
}

type Tenant struct {
	ID          string
	Name        string
	Description string
	OwnerID     string
	CreateAt    time.Time
	UpdatedAt   time.Time
}

type Role struct {
	ID          string
	Name        string
	Description string
	Permission  []string
}

type Permission struct {
	ID          string
	Name        string
	Description string
}

type Session struct {
	ID        string
	UserID    string
	CreatedAt time.Time
	ExpiresAt time.Time
	IPAdress  string
	UserAgent string
}

type Token struct {
	AccessToken  string
	RefreshToken string
	ExpiresIn    int64
	TokenType    string
}

type IDPClaims struct {
	UserID   string
	TenantID string
	IsAdmin  bool
	Username string
	jwt.RegisteredClaims
}
