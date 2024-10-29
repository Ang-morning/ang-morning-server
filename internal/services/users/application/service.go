package application

import (
	User "angmorning.com/internal/services/users/domain"
	"angmorning.com/internal/services/users/infrastructure"
)

type UserService struct {
	userRepository *infrastructure.UserRepository
}

func New(userRepository *infrastructure.UserRepository) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

func (it *UserService) OAuth() {
	// TODO: OAuth 로직
	user := User.Of("nickname", "email", []User.ProviderType{User.ProviderNAVER})

	it.userRepository.Create(user)
}
