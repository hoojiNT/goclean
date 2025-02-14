package domain

import "time"

type User struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Age       int       `json:"age"`
	Address   string    `json:"address"`
}

type UserRepository interface {
	Create(user *User) error
	GetByID(id uint) (*User, error)
	Update(user *User) error
	Delete(id uint) error
	GetAll() ([]User, error)
}

type UserUseCase interface {
	Create(user *User) error
	GetByID(id uint) (*User, error)
	Update(user *User) error
	Delete(id uint) error
	GetAll() ([]User, error)
}
