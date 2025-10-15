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
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (h *UserHandler) Index(c *gin.Context) {
	sort := c.DefaultQuery("sort", "name")
	order := c.DefaultQuery("order", "asc")
	limitString := c.DefaultQuery("limit", "15")
	role := c.DefaultQuery("role", "")

	limit, err := strconv.ParseInt(limitString, 10, 32)

	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.Response{
			MetaData: responses.MetaDataResponse{
				Code:    http.StatusInternalServerError,
				Message: "Error converting string to int32",
			},
		})

		return
	}

	params := &database.UserListParams{
		Sort:      sort,
		SortOrder: order,
		Till:      int32(limit),
		Role:      role,
	}

	users, err := h.service.Index(c.Request.Context(), *params)

	if err != nil {
		fmt.Println(err)

		c.JSON(http.StatusInternalServerError, responses.Response{
			MetaData: responses.MetaDataResponse{
				Code:    http.StatusInternalServerError,
				Message: "Failed to fetch users",
			},
		})

		return
	}

	c.JSON(http.StatusOK, responses.Response{
		MetaData: responses.MetaDataResponse{
			Code:    http.StatusOK,
			Message: "User Retrieved Successfully",
		},
		Data: users,
	})
}

func (h *UserHandler) Store(c *gin.Context) {
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

	user, err := h.service.Store(c, *body)

	if err != nil {
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
			Message: "User Created Successfully",
		},
		Data: user,
	})
}
