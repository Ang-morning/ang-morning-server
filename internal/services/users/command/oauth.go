package command

import "angmorning.com/internal/services/users/domain"

type OauthCommand struct {
	Code     string              `json:"code"`
	Provider domain.ProviderType `json:"provider"`
}
