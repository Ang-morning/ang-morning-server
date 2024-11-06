package infrastructure

import (
	"context"
	"database/sql"

	httpCode "angmorning.com/internal/libs/http/http-code"
	httpError "angmorning.com/internal/libs/http/http-error"
	"angmorning.com/internal/services/hospitals/domain"
	"angmorning.com/internal/services/hospitals/infrastructure/internal"
	valueObject "angmorning.com/internal/services/value-object"
)

type HospitalRepository struct {
	query *internal.Queries
}

func New(con *sql.DB) *HospitalRepository {
	return &HospitalRepository{
		query: internal.New(con),
	}
}

func (repository *HospitalRepository) Save(hospital *domain.Hospital) (*domain.Hospital, error) {
	ctx := context.Background()
	savedHospital, err := repository.query.Save(ctx, internal.SaveParams{
		ID:          hospital.Id,
		Name:        hospital.Name,
		Phone:       hospital.Phone,
		RoadAddress: hospital.Address.RoadAddress,
		Latitude:    hospital.Address.Latitude,
		Longitude:   hospital.Address.Longitude,
		ZipCode:     hospital.Address.ZipCode,
	})

	if err != nil {
		return nil, httpError.New(httpCode.InternalServerError, err.Error(), "")
	}
	return &domain.Hospital{
		Id:    savedHospital.ID,
		Name:  savedHospital.Name,
		Phone: savedHospital.Phone,
		Address: valueObject.Address{
			RoadAddress: savedHospital.RoadAddress,
			Latitude:    savedHospital.Latitude,
			Longitude:   savedHospital.Longitude,
			ZipCode:     savedHospital.ZipCode,
		},
	}, nil
}

func (repository *HospitalRepository) FindByCity(cities []string) ([]*domain.Hospital, error) {
	ctx := context.Background()
	hospitals, err := repository.query.FindByCity(ctx, cities)
	if err != nil {
		return nil, httpError.New(httpCode.InternalServerError, err.Error(), "")
	}

	result := []*domain.Hospital{}
	for _, hospital := range hospitals {
		result = append(result, &domain.Hospital{
			Id:    hospital.ID,
			Name:  hospital.Name,
			Phone: hospital.Phone,
			Address: valueObject.Address{
				RoadAddress: hospital.RoadAddress,
				Latitude:    hospital.Latitude,
				Longitude:   hospital.Longitude,
				ZipCode:     hospital.ZipCode,
			},
		})
	}

	return result, nil
}

func (repository *HospitalRepository) CountByCity(cities []string) (int64, error) {
	ctx := context.Background()
	count, err := repository.query.CountByCity(ctx, cities)
	if err != nil {
		return 0, httpError.New(httpCode.InternalServerError, err.Error(), "")
	}

	return count, nil
}
