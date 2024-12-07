package response

import valueObject "angmorning.com/internal/services/value-object"

type ListResponse struct {
	Items []*HospitalListItem `json:"items"`
	Count int                 `json:"count"`
}

type HospitalListItem struct {
	Id      string              `json:"id"`
	Name    string              `json:"name"`
	Phone   string              `json:"phone"`
	Address valueObject.Address `json:"address"`
}
