package infrastructure

import (
	"context"
	"database/sql"
	"fmt"

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

func (repository *UserRepository) Save(user *domain.User) {
	ctx := context.Background()
	_, err := repository.query.Save(ctx, internal.SaveParams{
		ID:               user.Id,
		Email:            user.Email,
		Nickname:         user.Nickname,
		ProfileImageUrl:  sql.NullString{String: user.ProfileImageUrl, Valid: user.ProfileImageUrl != ""},
		Providers:        providerTypeToString(user.Providers),
		LastProviderType: string(user.LastProviderType),
	})
	if err != nil {
		fmt.Println(err)
	}
}

func (repository *UserRepository) FindByEmail(email string) *domain.User {
	ctx := context.Background()
	user, err := repository.query.FindByEmail(ctx, email)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return domain.Of(user.Nickname, user.Email, user.ProfileImageUrl.String, stringToProviderType(user.Providers))
}
