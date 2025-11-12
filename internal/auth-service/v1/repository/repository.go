package repository

import "idp_mvp/internal/auth-service/v1/models"

type UserRepository interface {
	Create(u *models.User) error
	FindByUsername(username string) (*models.User, error)
	Update(u *models.User) error
	Delete(id string) error
	FindById(id string) (*models.User, error)
}
