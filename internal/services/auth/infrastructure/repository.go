package infrastructure

import (
	"context"
	"database/sql"

	httpCode "angmorning.com/internal/libs/http/http-code"
	httpError "angmorning.com/internal/libs/http/http-error"
	"angmorning.com/internal/services/auth/domain"
	"angmorning.com/internal/services/auth/infrastructure/internal"
)

type AuthRepository struct {
	query *internal.Queries
}

func New(con *sql.DB) *AuthRepository {
	return &AuthRepository{
		query: internal.New(con),
	}
}

func (repository *AuthRepository) Save(refreshToken *domain.RefreshToken) (*domain.RefreshToken, error) {
	ctx := context.Background()
	token, err := repository.query.Save(ctx, internal.SaveParams{
		UserId:     refreshToken.UserId,
		Value:      refreshToken.Value,
		ClientInfo: sql.NullString{String: refreshToken.ClientInfo, Valid: refreshToken.ClientInfo != ""},
	})
	if err != nil {
		return nil, httpError.New(httpCode.InternalServerError, err.Error(), "")
	}

	return &domain.RefreshToken{
		Id:         token.ID,
		UserId:     token.UserId,
		Value:      token.Value,
		ClientInfo: token.ClientInfo.String,
	}, nil
}

func (repository *AuthRepository) Delete(refreshToken *domain.RefreshToken) error {
	ctx := context.Background()
	err := repository.query.Delete(ctx, refreshToken.Id)
	if err != nil {
		return httpError.New(httpCode.InternalServerError, err.Error(), "")
	}

	return nil
}
