package domain

import (
	"github.com/google/uuid"
)

type RefreshToken struct {
	Id         int32
	Value      string
	UserId     uuid.UUID
	ClientInfo string
}

func New(value string, userId uuid.UUID, clientInfo string) *RefreshToken {
	return &RefreshToken{
		Value:      value,
		UserId:     userId,
		ClientInfo: clientInfo,
	}
}
