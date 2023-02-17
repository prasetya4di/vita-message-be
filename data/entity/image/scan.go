package image

import "vita-message-service/data/entity"

type Scan struct {
	Messages      []entity.Message `json:"messages"`
	Possibilities []Possibility    `json:"possibilities"`
}
