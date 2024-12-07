package command

import "github.com/google/uuid"

type WriteCommand struct {
	Content    string    `json:"content"`
	HospitalId uuid.UUID `json:"hospitalId"`
	Rating     int       `json:"rating"`
}
