package infrastructure

import (
	"context"
	"database/sql"

	httpCode "angmorning.com/internal/libs/http/http-code"
	httpError "angmorning.com/internal/libs/http/http-error"
	"angmorning.com/internal/services/reviews/domain"
	"angmorning.com/internal/services/reviews/infrastructure/internal"
	"github.com/google/uuid"
)

type ReviewRepository struct {
	query *internal.Queries
}

func New(con *sql.DB) *ReviewRepository {
	return &ReviewRepository{
		query: internal.New(con),
	}
}

func (repository *ReviewRepository) Save(review *domain.Review) (*domain.Review, error) {
	ctx := context.Background()
	r, err := repository.query.Save(ctx, internal.SaveParams{
		ID:         review.Id,
		UserId:     review.UserId,
		HospitalId: review.HospitalId,
		Content:    review.Content,
		Rating:     int32(review.Rating),
	})
	if err != nil {
		return nil, httpError.New(httpCode.InternalServerError, err.Error(), "")
	}

	return &domain.Review{
		Id:         r.ID,
		UserId:     r.UserId,
		HospitalId: r.HospitalId,
		Content:    r.Content,
		Rating:     int(r.Rating),
	}, nil
}

func (repository *ReviewRepository) FindByUserIdAndHospitalId(userId, hospitalId uuid.UUID) (*domain.Review, error) {
	ctx := context.Background()
	r, err := repository.query.FindByUserIdAndHospitalId(ctx, internal.FindByUserIdAndHospitalIdParams{
		UserId:     userId,
		HospitalId: hospitalId,
	})
	if err != nil {
		return nil, httpError.New(httpCode.InternalServerError, err.Error(), "")
	}

	return &domain.Review{
		Id:         r.ID,
		UserId:     r.UserId,
		HospitalId: r.HospitalId,
		Content:    r.Content,
		Rating:     int(r.Rating),
	}, nil
}
