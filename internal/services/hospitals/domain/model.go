package domain

import (
	httpCode "angmorning.com/internal/libs/http/http-code"
	httpError "angmorning.com/internal/libs/http/http-error"
	valueObject "angmorning.com/internal/services/value-object"
	"github.com/google/uuid"
)

type Hospital struct {
	Id      uuid.UUID           `json:"id"`
	Name    string              `json:"name"`
	Phone   string              `json:"phone"`
	Address valueObject.Address `json:"address"`
}

func New(name, phone string, address valueObject.Address) (*Hospital, error) {
	uuid, err := uuid.NewV7()
	if err != nil {
		return nil, httpError.New(httpCode.InternalServerError, "Failed to generate uuid.", "")
	}

	return &Hospital{
		Id:      uuid,
		Phone:   phone,
		Name:    name,
		Address: address,
	}, nil
}
