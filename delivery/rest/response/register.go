package response

import "vita-message-service/data/entity"

type RegisterResponse struct {
	User    User            `json:"user"`
	Message *entity.Message `json:"message"`
}
