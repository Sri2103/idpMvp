package repository

import (
	"errors"
	"sync"
	"time"

	"idp_mvp/internal/auth-service/v1/models"
)

// error variables
var (
	ErrorUserNotFound = errors.New("user not found")
	ErrorUserExists   = errors.New("user exists")
)

type InMemoryUserRepo struct {
	mu     sync.RWMutex
	store  map[string]*models.User
	byName map[string]string
}

// create the repo
func NewInMemoryReps() *InMemoryUserRepo {
	return &InMemoryUserRepo{
		store:  make(map[string]*models.User),
		byName: make(map[string]string),
	}
}

//

func (r *InMemoryUserRepo) Create(u *models.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	// unique username
	if _, ok := r.byName[u.Username]; ok {
		return ErrorUserExists
	}

	// assert time stamping
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
	r.store[u.ID] = u
	r.byName[u.Username] = u.ID

	return nil
}

// user findById
func (r *InMemoryUserRepo) FindById(id string) (*models.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	u, ok := r.store[id]
	if !ok {
		return nil, ErrorUserNotFound
	}
	return u, nil
}

// user findby Username
func (r *InMemoryUserRepo) FindByUsername(username string) (*models.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	id, ok := r.byName[username]
	if !ok {
		return nil, ErrorUserNotFound
	}
	u, ok := r.store[id]
	if !ok {
		return nil, ErrorUserNotFound
	}

	return u, nil
}

// update
func (r *InMemoryUserRepo) Update(u *models.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	_, err := r.FindById(u.ID)
	if err != nil {
		return err
	}

	u.UpdatedAt = time.Now()
	r.store[u.ID] = u
	r.byName[u.Username] = u.ID
	return nil
}

// delete user
func (r *InMemoryUserRepo) Delete(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	u, err := r.FindById(id)
	if err != nil {
		return err
	}

	delete(r.store, id)
	delete(r.byName, u.Username)
	return nil
}
