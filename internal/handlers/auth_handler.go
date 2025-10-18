package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"ranggaAdiPratama/go-with-claude/internal/database"
	"ranggaAdiPratama/go-with-claude/internal/requests"
	"ranggaAdiPratama/go-with-claude/internal/responses"
	"ranggaAdiPratama/go-with-claude/internal/service"
	"ranggaAdiPratama/go-with-claude/internal/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type AuthHandler struct {
	user *service.UserService
}

func NewAuthHandler(user *service.UserService) *AuthHandler {
	return &AuthHandler{
		user: user,
	}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var request requests.StoreUserRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		var ve validator.ValidationErrors

		if errors.As(err, &ve) {
			out := make([]string, len(ve))

			for i, fe := range ve {
				out[i] = utils.HumanizeError(fe)
			}

			c.JSON(http.StatusBadRequest, responses.Response{
				MetaData: responses.MetaDataResponse{
					Code:    http.StatusBadRequest,
					Message: "Invalid input. Please check your data.",
				},
				Data: out,
			})

			return
		}

		c.JSON(http.StatusBadRequest, responses.Response{
			MetaData: responses.MetaDataResponse{
				Code:    http.StatusBadRequest,
				Message: "Invalid input. Please check your data.",
			},
		})

		return
	}

	hashed, err := utils.HashPassword(request.Password)

	if err != nil {
		fmt.Println(err)

		c.JSON(http.StatusInternalServerError, responses.Response{
			MetaData: responses.MetaDataResponse{
				Code:    http.StatusInternalServerError,
				Message: "Failed in hashing password",
			},
		})

		return
	}

	body := &database.StoreUserParams{
		Name:     request.Name,
		Email:    request.Email,
		Username: request.Username,
		Password: hashed,
		Role:     request.Role,
	}

	_, err = h.user.Store(c, *body)

	if err != nil {
		if errors.Is(err, errors.New("e-mail already taken")) {
			c.JSON(http.StatusBadGateway, responses.Response{
				MetaData: responses.MetaDataResponse{
					Code:    http.StatusBadGateway,
					Message: "e-mail already taken",
				},
			})

			return
		} else if errors.Is(err, errors.New("username already taken")) {
			c.JSON(http.StatusBadGateway, responses.Response{
				MetaData: responses.MetaDataResponse{
					Code:    http.StatusBadGateway,
					Message: "username already taken",
				},
			})

			return
		}

		fmt.Println(err)

		c.JSON(http.StatusInternalServerError, responses.Response{
			MetaData: responses.MetaDataResponse{
				Code:    http.StatusInternalServerError,
				Message: "Failed in storing user",
			},
		})

		return
	}

	c.JSON(http.StatusCreated, responses.Response{
		MetaData: responses.MetaDataResponse{
			Code:    http.StatusCreated,
			Message: "Registration completed",
		},
	})
}
