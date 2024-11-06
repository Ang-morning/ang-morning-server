package application

import (
	httpError "angmorning.com/internal/libs/http/http-error"
	"angmorning.com/internal/libs/oauth"
	auth "angmorning.com/internal/services/auth/application"
	"angmorning.com/internal/services/users/command"
	User "angmorning.com/internal/services/users/domain"
	"angmorning.com/internal/services/users/infrastructure"
	"angmorning.com/internal/services/users/response"
)

type UserService struct {
	userRepository *infrastructure.UserRepository
	oauthFactory   *oauth.OauthClientFactory
	authService    *auth.AuthService
}

func New(userRepository *infrastructure.UserRepository, oauthFactory *oauth.OauthClientFactory, authService *auth.AuthService) *UserService {
	return &UserService{
		userRepository: userRepository,
		oauthFactory:   oauthFactory,
		authService:    authService,
	}
}

func (it *UserService) OAuth(command command.OauthCommand, clientInfo string) (*response.OAuthResponse, error) {
	client := it.oauthFactory.GetClient(User.ProviderKAKAO)
	token, err := client.GetToken(command.Code)
	if err != nil {
		return nil, httpError.Wrap(err)
	}

	userInfo, err := client.GetUserInfo(token)
	if err != nil {
		return nil, httpError.Wrap(err)
	}

	user, err := it.userRepository.FindByEmail(userInfo.Email)
	if err != nil {
		return nil, httpError.Wrap(err)
	}

	if user == nil {
		user, err = User.Of(userInfo.Nickname, userInfo.Email, userInfo.ProfileImageUrl, []User.ProviderType{command.Provider})
		if err != nil {
			return nil, httpError.Wrap(err)
		}
	} else {
		user.SignIn(command.Provider)
	}

	user, err = it.userRepository.Save(user)
	if err != nil {
		return nil, httpError.Wrap(err)
	}

	accessToken, err := it.authService.CreateToken(user.Id, clientInfo)
	if err != nil {
		return nil, httpError.Wrap(err)
	}

	return &response.OAuthResponse{AccessToken: accessToken}, nil
}
