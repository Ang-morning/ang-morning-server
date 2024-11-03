package application

import (
	"strconv"
	"time"

	"angmorning.com/internal/config"
	httpCode "angmorning.com/internal/libs/http/http-code"
	httpError "angmorning.com/internal/libs/http/http-error"
	"angmorning.com/internal/libs/jwt"
	"angmorning.com/internal/services/auth/domain"
	"angmorning.com/internal/services/auth/infrastructure"
	"github.com/google/uuid"
)

type AuthService struct {
	authRepository *infrastructure.AuthRepository
}

func New(authRepository *infrastructure.AuthRepository) *AuthService {
	return &AuthService{
		authRepository: authRepository,
	}
}

func (service *AuthService) CreateToken(userId uuid.UUID, clientInfo string) (string, error) {
	accessTokenExpiredAfterHour, err := strconv.Atoi(config.AccessTokenExpiredAfterHour)
	if err != nil {
		return "", httpError.New(httpCode.InternalServerError, "Failed to convert access token expired after hour to integer.", "")
	}

	refreshTokenExpiredAfterHour, err := strconv.Atoi(config.RefreshTokenExpiredAfterHour)
	if err != nil {
		return "", httpError.New(httpCode.InternalServerError, "Failed to convert refresh token expired after hour to integer.", "")
	}

	accessToken, err := jwt.Sign(userId, time.Hour*time.Duration(accessTokenExpiredAfterHour))
	if err != nil {
		return "", httpError.Wrap(err)
	}

	refreshToken, err := jwt.Sign(nil, time.Hour*time.Duration(refreshTokenExpiredAfterHour))
	if err != nil {
		return "", httpError.Wrap(err)
	}

	refreshTokenModel := domain.New(refreshToken, userId, clientInfo)
	_, err = service.authRepository.Save(refreshTokenModel)

	if err != nil {
		return "", httpError.Wrap(err)
	}

	return accessToken, error(nil)
}
