package impl

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"vita-message-service/delivery/rest/handler"
	"vita-message-service/delivery/rest/request"
	"vita-message-service/delivery/rest/response"
	"vita-message-service/usecase"
	"vita-message-service/util/token"
)

type authHandler struct {
	login             usecase.Login
	register          usecase.Register
	addInitialMessage usecase.AddInitialMessage
}

func NewAuthHandler(login usecase.Login, register usecase.Register, initiaMessage usecase.AddInitialMessage) handler.AuthHandler {
	return &authHandler{
		login:             login,
		register:          register,
		addInitialMessage: initiaMessage,
	}
}

func (ah *authHandler) Login(c *gin.Context) {
	var loginRequest request.LoginRequest

	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	u, err := ah.login.Invoke(loginRequest)
	if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword || err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
			return
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	t, err := token.GenerateToken(u.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	loginResponse := response.LoginResponse{
		User: response.User{
			Email:     u.Email,
			FirstName: u.FirstName,
			LastName:  u.LastName,
			Nickname:  u.Nickname,
			BirthDate: u.BirthDate,
			Token:     t,
		},
	}

	c.JSON(http.StatusOK, gin.H{"data": loginResponse})
}

func (ah *authHandler) Register(c *gin.Context) {
	var registerRequest request.RegisterRequest

	if err := c.ShouldBindJSON(&registerRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	u, err := ah.register.Invoke(registerRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	t, err := token.GenerateToken(u.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	initialMessage, err := ah.addInitialMessage.Invoke(u.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	newUser := response.User{
		Email:     u.Email,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Nickname:  u.Nickname,
		BirthDate: u.BirthDate,
		Token:     t,
	}

	registerResponse := response.RegisterResponse{
		User:    newUser,
		Message: initialMessage,
	}

	c.JSON(http.StatusCreated, gin.H{"data": registerResponse})
}
