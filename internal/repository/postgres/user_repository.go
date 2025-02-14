package postgres

import (
	"goclean/internal/domain"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) domain.UserRepository {
	return &userRepository{db}
}

func (r *userRepository) Create(user *domain.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) GetByID(id uint) (*domain.User, error) {
	var user domain.User
	err := r.db.First(&user, id).Error
	return &user, err
}

func (r *userRepository) Update(user *domain.User) error {
	return r.db.Save(user).Error
}

func (r *userRepository) Delete(id uint) error {
	return r.db.Delete(&domain.User{}, id).Error
}

func (r *userRepository) GetAll() ([]domain.User, error) {
	var users []domain.User
	err := r.db.Find(&users).Error
	return users, err
}
