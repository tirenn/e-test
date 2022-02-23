package auth

import (
	"tirenn/test-efishery/auth-app/models"

	"gorm.io/gorm"
)

type RepositoryContract interface {
	Create(auth *models.Auth) error
	Login(phone, password string) (models.Auth, error)
}

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) RepositoryContract {
	return &Repository{
		db: db,
	}
}

func (r *Repository) Create(auth *models.Auth) error {
	return r.db.Create(auth).Error
}

func (r *Repository) Login(phone, password string) (models.Auth, error) {
	var auth models.Auth
	err := r.db.First(&auth, "phone = ? AND password = ?", phone, password).Error
	return auth, err
}
