package domain

import (
	httpCode "angmorning.com/internal/libs/http/http-code"
	httpError "angmorning.com/internal/libs/http/http-error"
	"github.com/google/uuid"
)

type ProviderType string

const (
	ProviderNAVER  ProviderType = "NAVER"
	ProviderKAKAO  ProviderType = "KAKAO"
	ProviderGOOGLE ProviderType = "GOOGLE"
)

type User struct {
	Id               uuid.UUID      `json:"id"`
	Nickname         string         `json:"nickname"`
	Email            string         `json:"email"`
	ProfileImageUrl  string         `json:"profileImageUrl"`
	Providers        []ProviderType `json:"providers"`
	LastProviderType ProviderType   `json:"lastProviderType"`
}

func Of(nickname string, email string, profileImageUrl string, providers []ProviderType) (*User, error) {
	uuid, err := uuid.NewV7()
	if err != nil {
		return nil, httpError.New(httpCode.InternalServerError, "Failed to generate uuid.", "")
	}
	return &User{
		Id:              uuid,
		Nickname:        nickname,
		ProfileImageUrl: profileImageUrl,
		Email:           email,
		Providers:       providers,
	}, nil
}

func (user *User) SignIn(provider ProviderType) {
	hasProvider := false
	for _, p := range user.Providers {
		if p == provider {
			hasProvider = true
		}
	}

	if !hasProvider {
		user.Providers = append(user.Providers, provider)
	}

	user.LastProviderType = provider
}
