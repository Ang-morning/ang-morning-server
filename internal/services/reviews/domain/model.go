package domain

import (
	httpCode "angmorning.com/internal/libs/http/http-code"
	httpError "angmorning.com/internal/libs/http/http-error"
	"github.com/google/uuid"
)

type Review struct {
	Id         uuid.UUID `json:"id"`
	UserId     uuid.UUID `json:"userId"`
	HospitalId uuid.UUID `json:"hospitalId"`
	Content    string    `json:"content"`
	Rating     int       `json:"rating"`
}

func Of(userId, hospitalId uuid.UUID, content string, rating int) (*Review, error) {
	uuid, err := uuid.NewV7()
	if err != nil {
		return nil, httpError.New(httpCode.InternalServerError, "Failed to generate uuid.", "")
	}

	return &Review{
		Id:         uuid,
		UserId:     userId,
		HospitalId: hospitalId,
		Content:    content,
		Rating:     rating,
	}, nil
}
