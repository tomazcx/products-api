package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
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

// ShowAccount godoc
// @Summary      Show an account
// @Description  Get user data by email
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        email   path   string  true  "Account Email"
// @Success      200
// @Failure      400
// @Failure      404
// @Failure      500
// @Router       /users/{email} [get]
func (h *UserHandler) GetByEmail(w http.ResponseWriter, r *http.Request) {
	email := chi.URLParam(r, "email")

	if email == "" {
		http.Error(w, "Error: invalid email", http.StatusBadRequest)
		return
	}

	findUserByEmailUseCase := h.factory.FindUserByEmailUseCase()
	user, err := findUserByEmailUseCase.Execute(email)

	if err != nil {
		if httpErr, ok := err.(*httperr.HttpError); ok {
			http.Error(w, httpErr.Message, httpErr.StatusCode)
			return
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

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
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

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

	accessToken := struct {
		AccessToken string `json:"access_token"`
	}{
		AccessToken: token,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(accessToken)
}
