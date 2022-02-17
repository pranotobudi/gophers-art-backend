package controllers

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/pranotobudi/gophers-art-backend/app/common"
	"github.com/pranotobudi/gophers-art-backend/app/models"

	"golang.org/x/crypto/bcrypt"
)

type UserRegistrationResponse struct {
	ID       uint   `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}
type UserRegistrationRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func UserRegistrationResponseFormatter(user models.User) UserRegistrationResponse {
	formatter := UserRegistrationResponse{
		ID:       user.ID,
		Email:    user.Email,
		Username: user.Username,
	}
	return formatter
}

type UserLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserLoginResponse struct {
	ID        uint   `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	AuthToken string `json:"auth_token"`
}

func UserResponseFormatter(user models.User, auth_token string) UserLoginResponse {
	formatter := UserLoginResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		AuthToken: auth_token,
	}
	return formatter
}

type userHandler struct {
	repository models.UserRepository
}

func NewUserHandler() *userHandler {
	repository := models.NewUserRepository()

	return &userHandler{repository}
}

func (h *userHandler) RegisterUser(c App) common.Response {
	// Input Binding
	userReg := new(UserRegistrationRequest)
	if err := c.Params.BindJSON(userReg); err != nil {
		return common.ResponseErrorFormatter(http.StatusBadRequest, err)
	}

	// Process Input - User Registration
	user := models.User{}
	user.Username = userReg.Username
	user.Email = userReg.Email
	user.Password = common.GeneratePassword(userReg.Password)
	savedUser, err := h.repository.AddUser(user)
	if err != nil {
		return common.ResponseErrorFormatter(http.StatusInternalServerError, err)
	}

	// Success ProductResponse
	data := UserRegistrationResponseFormatter(*savedUser)

	return common.ResponseFormatter(http.StatusOK, "success", "get user successfull", data)
}

func (h *userHandler) UserLogin(c App) common.Response {
	// Input Binding
	userLogin := UserLoginRequest{}
	if err := c.Params.BindJSON(&userLogin); err != nil {
		return common.ResponseErrorFormatter(http.StatusBadRequest, err)
	}

	// Process Input
	authUser, err := h.AuthUser(userLogin)
	fmt.Println("We're IN HERE: USERLOGIN INSIDE: authUser: ", authUser)
	if err != nil {
		return common.ResponseErrorFormatter(http.StatusBadRequest, err)
	}

	// Create JWT token
	auth_token, err := h.CreateAccessToken()
	if err != nil {
		return common.ResponseErrorFormatter(http.StatusInternalServerError, err)
	}

	// Success UserLoginResponse
	data := UserResponseFormatter(*authUser, auth_token)
	return common.ResponseFormatter(http.StatusOK, "success", "user login successfull", data)
}

func (h *userHandler) AuthUser(req UserLoginRequest) (*models.User, error) {
	username := req.Username
	password := req.Password
	fmt.Println("AUTHUSER CALLED, username: ", username, " password: ", password)

	//check Author Table
	user, err := h.repository.GetUserByUsername(username)
	if err != nil {
		return nil, fmt.Errorf("username is not registered")
	}

	test, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	fmt.Printf("COMPARES: %s %s \n", user.Password, string(test))
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, fmt.Errorf("invalid password")
	}
	return user, nil
}

func (s *userHandler) CreateAccessToken() (string, error) {
	claims := jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 3).Unix(),
		IssuedAt:  time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedKey, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		return signedKey, err
	}

	return signedKey, nil
}
