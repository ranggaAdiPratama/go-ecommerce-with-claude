package handlers

import (
	"database/sql"
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
	service *service.AuthService
	user    *service.UserService
}

func NewAuthHandler(service *service.AuthService, user *service.UserService) *AuthHandler {
	return &AuthHandler{
		service: service,
		user:    user,
	}
}

func (h *AuthHandler) Login(c *gin.Context) {
	var request requests.LoginRequest

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

	data, err := h.service.Login(c, request)

	if err == sql.ErrNoRows {
		c.JSON(http.StatusInternalServerError, responses.Response{
			MetaData: responses.MetaDataResponse{
				Code:    http.StatusInternalServerError,
				Message: "Username not found",
			},
		})

		return
	}

	if err != nil {
		statusCode := http.StatusInternalServerError
		message := "Invalid Password"

		if err.Error() == "refresh token has been revoked" {
			statusCode = http.StatusUnauthorized

			message = "Refresh token has been revoked. Please login again."
		}

		if err.Error() == "refresh token has been revoked" {
			statusCode = http.StatusUnauthorized

			message = "Refresh token has expired. Please login again."
		}

		c.JSON(statusCode, responses.Response{
			MetaData: responses.MetaDataResponse{
				Code:    statusCode,
				Message: message,
			},
		})

		return
	}

	c.JSON(http.StatusOK, responses.Response{
		MetaData: responses.MetaDataResponse{
			Code:    http.StatusOK,
			Message: "Login success",
		},
		Data: data,
	})
}

func (h *AuthHandler) Logout(c *gin.Context) {
	userPayload, exists := c.Get("user")

	if !exists {
		c.JSON(http.StatusUnauthorized, responses.Response{
			MetaData: responses.MetaDataResponse{
				Code:    http.StatusUnauthorized,
				Message: "Unauthorized",
			},
		})

		return
	}

	err := h.service.Logout(c, userPayload.(*utils.TokenPayload).ID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.Response{
			MetaData: responses.MetaDataResponse{
				Code:    http.StatusInternalServerError,
				Message: "Error :(",
			},
		})

		return
	}

	c.JSON(http.StatusOK, responses.Response{
		MetaData: responses.MetaDataResponse{
			Code:    http.StatusOK,
			Message: "Logout success",
		},
	})
}

func (h *AuthHandler) RefreshToken(c *gin.Context) {
	var request requests.RefreshTokenRequest

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

	data, err := h.service.RefreshToken(c, request)

	if err != nil {
		fmt.Println(err)

		c.JSON(http.StatusInternalServerError, responses.Response{
			MetaData: responses.MetaDataResponse{
				Code:    http.StatusInternalServerError,
				Message: "Error :(",
			},
		})

		return
	}

	c.JSON(http.StatusOK, responses.Response{
		MetaData: responses.MetaDataResponse{
			Code:    http.StatusOK,
			Message: "Refresh Token success",
		},
		Data: data,
	})
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
		if err.Error() == "e-mail already taken" || err.Error() == "username already taken" {
			c.JSON(http.StatusBadRequest, responses.Response{
				MetaData: responses.MetaDataResponse{
					Code:    http.StatusBadRequest,
					Message: utils.CapitalizeFirst(err.Error()),
				},
			})

			return
		}

		fmt.Println(err)

		c.JSON(http.StatusInternalServerError, responses.Response{
			MetaData: responses.MetaDataResponse{
				Code:    http.StatusInternalServerError,
				Message: "Failed in database transaction",
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
