package application

import (
	httpError "angmorning.com/internal/libs/http/http-error"
	"angmorning.com/internal/services/hospitals/command"
	"angmorning.com/internal/services/hospitals/infrastructure"
	"angmorning.com/internal/services/hospitals/response"
)

type HospitalService struct {
	hopsitalRepository *infrastructure.HospitalRepository
}

func New(hopsitalRepository *infrastructure.HospitalRepository) *HospitalService {
	return &HospitalService{
		hopsitalRepository: hopsitalRepository,
	}
}

func (service *HospitalService) List(command command.ListCommand) (*response.ListResponse, error) {
	//TODO: pagination
	hospitals, err := service.hopsitalRepository.FindByCity(command.Cities)
	if err != nil {
		return nil, httpError.Wrap(err)
	}

	count, err := service.hopsitalRepository.CountByCity(command.Cities)
	if err != nil {
		return nil, httpError.Wrap(err)
	}

	hospitalListItems := []*response.HospitalListItem{}
	for _, hospital := range hospitals {
		hospitalListItems = append(hospitalListItems, &response.HospitalListItem{
			Id:      hospital.Id.String(),
			Name:    hospital.Name,
			Phone:   hospital.Phone,
			Address: hospital.Address,
		})
	}

	return &response.ListResponse{
		Items: hospitalListItems,
		Count: int(count),
	}, nil
}
