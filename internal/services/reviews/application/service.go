package application

import (
	httpError "angmorning.com/internal/libs/http/http-error"
	"angmorning.com/internal/services/reviews/command"
	reviewModel "angmorning.com/internal/services/reviews/domain"
	"angmorning.com/internal/services/reviews/infrastructure"
	"github.com/google/uuid"
)

type ReviewService struct {
	reviewRepository *infrastructure.ReviewRepository
}

func New(reviewRepository *infrastructure.ReviewRepository) *ReviewService {
	return &ReviewService{
		reviewRepository: reviewRepository,
	}
}

func (service *ReviewService) Write(userId uuid.UUID, command command.WriteCommand) (string, error) {
	// TODO: check if user already reviewed the hospital
	review, err := reviewModel.Of(userId, command.HospitalId, command.Content, command.Rating)
	if err != nil {
		return "", httpError.Wrap(err)
	}

	service.reviewRepository.Save(review)

	return "success", nil
}
