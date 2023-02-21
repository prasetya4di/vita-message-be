package usecase

import (
	"github.com/gin-gonic/gin"
	"vita-message-service/data/entity"
)

type GetCurrentUser interface {
	Invoke(c *gin.Context) (*entity.User, error)
}
