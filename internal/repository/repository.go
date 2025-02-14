// internal/usecase/user_usecase.go
package usecase

import "goclean/internal/domain"

type userUseCase struct {
	userRepo domain.UserRepository
}

func NewUserUseCase(userRepo domain.UserRepository) domain.UserUseCase {
	return &userUseCase{userRepo}
}

func (uc *userUseCase) Create(user *domain.User) error {
	return uc.userRepo.Create(user)
}

func (uc *userUseCase) GetByID(id uint) (*domain.User, error) {
	return uc.userRepo.GetByID(id)
}

func (uc *userUseCase) Update(user *domain.User) error {
	return uc.userRepo.Update(user)
}

func (uc *userUseCase) Delete(id uint) error {
	return uc.userRepo.Delete(id)
}

func (uc *userUseCase) GetAll() ([]domain.User, error) {
	return uc.userRepo.GetAll()
}
