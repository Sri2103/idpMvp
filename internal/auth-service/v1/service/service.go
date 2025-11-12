package service

import (
	"errors"
	"time"

	"idp_mvp/internal/auth-service/v1/models"
	"idp_mvp/internal/auth-service/v1/repository"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// errors
var (
	ErrorInvalidCredentials = errors.New("invalid credentials")
	ErrorUserExist          = errors.New("user already exists")
)

// service struct
type UserService struct {
	repo repository.UserRepository
}

// initiating userservice
func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

// registration service
func (s *UserService) Register(username, email, password, tenantId string, roles []string, isAdmin bool) (*models.User, error) {
	// check the existing user by username
	_, err := s.repo.FindByUsername(username)
	if err == nil {
		return nil, ErrorUserExist
	}

	// generate password hash
	pwHash, err := hashPassword(password)
	if err != nil {
		return nil, err
	}

	// add user to store
	user := &models.User{
		ID:           uuid.NewString(),
		Username:     username,
		PasswordHash: pwHash,
		IsAdmin:      isAdmin,
		Email:        email,
		TenantID:     tenantId,
		Roles:        roles,
		IsActive:     true,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	err = s.repo.Create(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// login route service
func (s *UserService) Authenticate(username string, password string) (*models.User, error) {
	// fetch user from store
	u, err := s.repo.FindByUsername(username)
	if err != nil {
		return nil, err
	}

	if err := comparePassword(u.PasswordHash, password); err != nil {
		return nil, ErrorInvalidCredentials
	}

	return u, nil
}

// change password
func (s *UserService) ChangePassword(userId, oldPassword, newPassword string) error {
	// get the oldpassword from store by userId
	u, err := s.repo.FindById(userId)
	if err != nil {
		return err
	}

	// compare old password with hashed
	if err := comparePassword(u.PasswordHash, oldPassword); err != nil {
		return ErrorInvalidCredentials
	}

	// generate hash for new password
	newHash, err := hashPassword(newPassword)
	if err != nil {
		return err
	}

	// assign to user and update
	u.PasswordHash = newHash
	u.UpdatedAt = time.Now()
	return s.repo.Update(u)
}

// TODO: Implement secure forgot/reset (email + token store)
func (s *UserService) ForgotPassword(email string) error             { return nil }
func (s *UserService) ResetPassword(token, newPassword string) error { return nil }

func hashPassword(password string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(b), err
}

func comparePassword(hash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
