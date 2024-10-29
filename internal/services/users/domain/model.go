package domain

import "github.com/google/uuid"

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

func Of(nickname string, email string, profileImageUrl string, providers []ProviderType) *User {
	uuid, err := uuid.NewV7()
	if err != nil {
		// TODO: error handling
		panic(err)
	}
	return &User{
		Id:              uuid,
		Nickname:        nickname,
		ProfileImageUrl: profileImageUrl,
		Email:           email,
		Providers:       providers,
	}
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
