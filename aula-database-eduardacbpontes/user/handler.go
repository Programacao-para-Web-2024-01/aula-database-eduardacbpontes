package user

import (
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
)

var key = []byte("TOKEN_SECRETO")
var jwtManager = jwt.New(jwt.SigningMethodHS256)

func createToken() (string, error) {
	return jwtManager.SignedString(key)
}

func validateToken(token string) (*jwt.Token, error) {
	return jwt.NewParser().Parse(token, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
}

type Controller struct {
	service *Service
}

func NewController(service *Service) *Controller {
	return &Controller{service: service}
}

func (s *Controller) Register(w http.ResponseWriter, req *http.Request) {
	var user User
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	err = s.service.Register(user)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	token, err := createToken()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	fmt.Fprint(w, token)
}

func (s *Controller) Login(w http.ResponseWriter, req *http.Request) {
	var user User
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	err = s.service.Login(user.Username, user.Password)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	token, err := createToken()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	fmt.Fprint(w, token)
}
