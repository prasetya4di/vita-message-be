package image

import "vita-message-service/data/entity"

type Scan struct {
	Message       entity.Message `json:"message"`
	Possibilities []Possibility  `json:"possibilities"`
}
