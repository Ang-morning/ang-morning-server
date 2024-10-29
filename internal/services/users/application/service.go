package application

import (
	"fmt"

	"angmorning.com/internal/libs/oauth"
	"angmorning.com/internal/services/users/command"
	User "angmorning.com/internal/services/users/domain"
	"angmorning.com/internal/services/users/infrastructure"
)

type UserService struct {
	userRepository *infrastructure.UserRepository
	oauthFactory   *oauth.OauthClientFactory
}

func New(userRepository *infrastructure.UserRepository, oauthFactory *oauth.OauthClientFactory) *UserService {
	return &UserService{
		userRepository: userRepository,
		oauthFactory:   oauthFactory,
	}
}

func (it *UserService) OAuth(command command.OauthCommand) {
	client := it.oauthFactory.GetClient(User.ProviderKAKAO)
	token := client.GetToken(command.Code)
	userInfo := client.GetUserInfo(token)

	fmt.Println("@@@", userInfo)
	user := it.userRepository.FindByEmail(userInfo.Email)
	fmt.Println("!!!!", user)
	if user == nil {
		user = User.Of(userInfo.Nickname, userInfo.Email, userInfo.ProfileImageUrl, []User.ProviderType{command.Provider})
	} else {
		user.SignIn(command.Provider)
	}

	fmt.Println("###", user)

	it.userRepository.Save(user)
}
