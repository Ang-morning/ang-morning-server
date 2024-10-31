package infrastructure

import (
	"context"
	"database/sql"

	httpCode "angmorning.com/internal/libs/http/http-code"
	httpError "angmorning.com/internal/libs/http/http-error"
	"angmorning.com/internal/services/users/domain"
	"angmorning.com/internal/services/users/infrastructure/internal"
)

type UserRepository struct {
	query *internal.Queries
}

func New(con *sql.DB) *UserRepository {
	return &UserRepository{
		query: internal.New(con),
	}
}

func (repository *UserRepository) Save(user *domain.User) (*domain.User, error) {
	ctx := context.Background()
	u, err := repository.query.Save(ctx, internal.SaveParams{
		ID:               user.Id,
		Email:            user.Email,
		Nickname:         user.Nickname,
		ProfileImageUrl:  sql.NullString{String: user.ProfileImageUrl, Valid: user.ProfileImageUrl != ""},
		Providers:        providerTypeToString(user.Providers),
		LastProviderType: string(user.LastProviderType),
	})
	if err != nil {
		return nil, httpError.New(httpCode.InternalServerError, err.Error(), "")
	}

	return &domain.User{
		Id:               u.ID,
		Nickname:         u.Nickname,
		Email:            u.Email,
		ProfileImageUrl:  u.ProfileImageUrl.String,
		Providers:        stringToProviderType(u.Providers),
		LastProviderType: domain.ProviderType(u.LastProviderType),
	}, nil
}

func (repository *UserRepository) FindByEmail(email string) (*domain.User, error) {
	ctx := context.Background()
	user, err := repository.query.FindByEmail(ctx, email)
	if err != nil {
		return nil, httpError.New(httpCode.InternalServerError, err.Error(), "")
	}
	return &domain.User{
		Id:               user.ID,
		Nickname:         user.Nickname,
		Email:            user.Email,
		ProfileImageUrl:  user.ProfileImageUrl.String,
		Providers:        stringToProviderType(user.Providers),
		LastProviderType: domain.ProviderType(user.LastProviderType),
	}, nil
}
