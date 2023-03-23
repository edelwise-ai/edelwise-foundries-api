package http

import (
	"Foundries/models"
	"Foundries/usecases"
	"Foundries/utils"
	"github.com/gin-gonic/gin"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

// ResponseError will represent the http handler for error response
type ResponseError struct {
	Message string `json:"message"`
}

// UserHandler will represent the http handler for user
type UserHandler struct {
	UserUsecase usecases.UserUsecase
}

// NewUserHandler will initialize the users/ resources endpoint
func NewUserHandler(r *gin.Engine, us usecases.UserUsecase) {
	handler := &UserHandler{
		UserUsecase: us,
	}

	r.GET("/users", handler.Fetch)
	r.GET("/users/:id", handler.GetByID)
	r.POST("/users", handler.Store)
}

// Fetch will fetch the user data
func (u *UserHandler) Fetch(c *gin.Context) {
	users, err := u.UserUsecase.Fetch()
	if err != nil {
		c.JSON(500, ResponseError{Message: err.Error()})
		return
	}
	// Return user and message to client
	c.JSON(http.StatusOK, gin.H{
		"message": "Users fetched successfully",
		"users":   users,
	})
}

// GetByID will get user by given id
func (u *UserHandler) GetByID(c *gin.Context) {
	id := c.Param("id")
	user, err := u.UserUsecase.GetByID(id)
	if err != nil {
		c.JSON(500, ResponseError{Message: err.Error()})
		return
	}
	// Return user and message to client
	c.JSON(http.StatusOK, gin.H{
		"message": "User fetched successfully",
		"user":    user,
	})
}

// GetByEmail will get user by given email
func (u *UserHandler) GetByEmail(c *gin.Context) {
	email := c.Param("email")
	user, err := u.UserUsecase.GetByEmail(email)
	if err != nil {
		c.JSON(500, ResponseError{Message: err.Error()})
		return
	}
	// Return user and message to client
	c.JSON(http.StatusOK, gin.H{
		"message": "User fetched successfully",
		"user":    user,
	})
}

// Store will store the user by given request body
func (u *UserHandler) Store(c *gin.Context) {
	// Get user from request body
	var user models.User
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(500, ResponseError{Message: err.Error()})
		return
	}

	// Check if email or password is empty
	if user.Email == "" || user.Password == "" {
		c.JSON(http.StatusBadRequest, ResponseError{Message: "Email or Password is empty"})
		return
	}

	// check email format
	if utils.CheckFormat(user.Email) == false {
		c.JSON(http.StatusBadRequest, ResponseError{Message: "Email format is not correct"})
		return
	}

	// Check if email is already registered
	_, err = u.UserUsecase.GetByEmail(user.Email)
	if err == nil {
		c.JSON(http.StatusBadRequest, ResponseError{Message: "Email is already registered"})
		return
	}

	// Encrypt password with bcrypt
	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ResponseError{Message: err.Error()})
		return
	}
	user.Password = string(hasedPassword)

	// Generate nanoid
	user.ID, err = gonanoid.New(10)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ResponseError{Message: err.Error()})
		return
	}

	// Create user
	err = u.UserUsecase.Store(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ResponseError{Message: err.Error()})
		return
	}

	// Return user and message to client
	c.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
		"user":    user,
	})
}

// Login will login user and return user data if success
func (u *UserHandler) Login(c *gin.Context) {
	var user models.User
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(500, ResponseError{Message: err.Error()})
		return
	}

	// check if email or password is empty
	if user.Email == "" || user.Password == "" {
		c.JSON(http.StatusBadRequest, ResponseError{Message: "Email or Password is empty"})
		return
	}

	// check user input email format with regex
	if utils.CheckFormat(user.Email) == false {
		c.JSON(http.StatusBadRequest, ResponseError{Message: "Email format is not correct"})
		return
	}
	// check if email is registered
	us, err := u.UserUsecase.GetByEmail(user.Email)
	if err != nil {
		c.JSON(http.StatusNotFound, ResponseError{Message: "Email is not registered"})
		return
	}

	// check if password is correct
	if bcrypt.CompareHashAndPassword([]byte(us.Password), []byte(user.Password)) != nil {
		c.JSON(http.StatusUnauthorized, ResponseError{Message: "Password is incorrect"})
		return
	}

	//TODO: generate token with JWT

	// Return user and message to client
	c.JSON(http.StatusOK, gin.H{
		"message": "User logged in successfully",
		"user":    us,
	})
}
