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

func (it *UserRepository) Create(user *domain.User) {
	ctx := context.Background()
	_, err := it.query.Create(ctx, internal.CreateParams{
		ID:        user.Id,
		Email:     user.Email,
		Nickname:  user.Nickname,
		Providers: providerTypeToString(user.Providers),
	})
	if err != nil {
		fmt.Println(err)
	}
}
