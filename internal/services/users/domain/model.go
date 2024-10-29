package domain

import "github.com/google/uuid"

type ProviderType string

const (
	ProviderNAVER  ProviderType = "NAVER"
	ProviderKAKAO  ProviderType = "KAKAO"
	ProviderGOOGLE ProviderType = "GOOGLE"
)

type User struct {
	Id        uuid.UUID      `json:"id"`
	Nickname  string         `json:"nickname"`
	Email     string         `json:"email"`
	Providers []ProviderType `json:"providers"`
}

func Of(nickname string, email string, providers []ProviderType) *User {
	uuid, err := uuid.NewV7()
	if err != nil {
		// TODO: error handling
		panic(err)
	}
	return &User{
		Id:        uuid,
		Nickname:  nickname,
		Email:     email,
		Providers: providers,
	}
}
