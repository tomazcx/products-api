package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/jwtauth"
	"github.com/tomazcx/products-api/internal/data/factory"
	"github.com/tomazcx/products-api/internal/dto"
	"github.com/tomazcx/products-api/internal/infra/http/utils"
	"github.com/tomazcx/products-api/pkg/httperr"
)

type UserHandler struct {
	factory factory.UserFactory
}

func NewUserHandler(factory factory.UserFactory) *UserHandler {
	return &UserHandler{factory: factory}
}

// CreateUser godoc
// @Summary      Create an user
// @Description  Create an user
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        request   body   dto.UserDTO  true  "Create User DTO"
// @Success      201 {object} entity.User
// @Failure      422
// @Failure      500
// @Router       /users [post]
func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var createUserDTO dto.UserDTO

	err := json.NewDecoder(r.Body).Decode(&createUserDTO)

	if err != nil {
		http.Error(w, "Invalid json", http.StatusUnprocessableEntity)
		return
	}

	createUserUseCase := h.factory.CreateUserUseCase()
	user, err := createUserUseCase.Execute(createUserDTO)

	if err != nil {
		if httpErr, ok := err.(*httperr.HttpError); ok {
			http.Error(w, httpErr.Message, httpErr.StatusCode)
			return
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Location", fmt.Sprintf("/users/%s", createUserDTO.Email))
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

// Authenticate godoc
// @Summary      Authenticate
// @Description  Authenticate with email and password to get a token
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        request   body   dto.AuthDTO  true  "Auth  DTO"
// @Success      200 {object} dto.AccessTokenDTO
// @Failure      401
// @Failure      500
// @Router       /users/login [post]
func (h *UserHandler) Authenticate(w http.ResponseWriter, r *http.Request) {
	jwt := r.Context().Value("jwt").(*jwtauth.JWTAuth)
	exp := r.Context().Value("exp").(int)

	var authDTO dto.AuthDTO
	err := json.NewDecoder(r.Body).Decode(&authDTO)

	if err != nil {
		http.Error(w, "Invalid json", http.StatusUnprocessableEntity)
		return
	}

	authUserUseCase := h.factory.AuthUserUseCase()
	isCredentialsValid := authUserUseCase.Execute(authDTO)

	if !isCredentialsValid {
		http.Error(w, "Invalid Credentials", http.StatusUnauthorized)
		return
	}

	token, _ := utils.GenerateJWT(authDTO.Email, exp, jwt)

	accessToken := dto.AccessTokenDTO{
		AccessToken: token,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(accessToken)
}
