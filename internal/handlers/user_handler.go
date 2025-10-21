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
	"github.com/google/uuid"
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
	pageString := c.DefaultQuery("page", "1")
	role := c.DefaultQuery("role", "user")
	search := c.DefaultQuery("search", "")

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

	page, err := strconv.ParseInt(pageString, 10, 32)

	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.Response{
			MetaData: responses.MetaDataResponse{
				Code:    http.StatusInternalServerError,
				Message: "Error converting string to int32",
			},
		})

		return
	}

	if search != "" {
		search = utils.EscapeRegex(search)
	}

	params := &database.UserListParams{
		Page:      int32((page - 1) * limit),
		Role:      role,
		Search:    search,
		Sort:      sort,
		SortOrder: order,
		Till:      int32(limit),
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
			Message: "Users Retrieved Successfully",
		},
		Data: users,
	})
}

func (h *UserHandler) Destroy(c *gin.Context) {
	idParam := c.Param("id")

	id, err := uuid.Parse(idParam)

	if err != nil {
		c.JSON(http.StatusBadRequest, responses.Response{
			MetaData: responses.MetaDataResponse{
				Code:    http.StatusBadRequest,
				Message: "invalid ID Format",
			},
		})
	}

	err = h.service.Destroy(c, id)

	if err != nil {
		if errors.Is(err, errors.New("user not found")) {
			c.JSON(http.StatusNotFound, responses.Response{
				MetaData: responses.MetaDataResponse{
					Code:    http.StatusNotFound,
					Message: "User not found",
				},
			})

			return
		}

		fmt.Println(err)

		c.JSON(http.StatusNotFound, responses.Response{
			MetaData: responses.MetaDataResponse{
				Code:    http.StatusNotFound,
				Message: "User not found",
			},
		})

		return
	}

	c.JSON(http.StatusOK, responses.Response{
		MetaData: responses.MetaDataResponse{
			Code:    http.StatusOK,
			Message: "User Deleted Successfully",
		},
	})
}

func (h *UserHandler) Show(c *gin.Context) {
	idParam := c.Param("id")

	id, err := uuid.Parse(idParam)

	if err != nil {
		c.JSON(http.StatusBadRequest, responses.Response{
			MetaData: responses.MetaDataResponse{
				Code:    http.StatusBadRequest,
				Message: "invalid ID Format",
			},
		})
	}

	user, err := h.service.Show(c.Request.Context(), id)

	if err != nil {
		if errors.Is(err, errors.New("user not found")) {
			c.JSON(http.StatusNotFound, responses.Response{
				MetaData: responses.MetaDataResponse{
					Code:    http.StatusNotFound,
					Message: "User not found",
				},
			})

			return
		} else if errors.Is(err, errors.New("invalid User Id")) {
			c.JSON(http.StatusBadRequest, responses.Response{
				MetaData: responses.MetaDataResponse{
					Code:    http.StatusBadRequest,
					Message: "invalid User Id",
				},
			})

			return
		}

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
		Data: user,
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
			Message: "User Created Successfully",
		},
		Data: user,
	})
}

func (h *UserHandler) Update(c *gin.Context) {
	idParam := c.Param("id")

	id, err := uuid.Parse(idParam)

	if err != nil {
		c.JSON(http.StatusBadRequest, responses.Response{
			MetaData: responses.MetaDataResponse{
				Code:    http.StatusBadRequest,
				Message: "invalid ID Format",
			},
		})
	}

	var request requests.UpdateUserRequest

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

	body := &database.UpdateUserParams{
		Name:     request.Name,
		Email:    request.Email,
		Username: request.Username,
		ID:       id,
	}

	if request.Password != "" {
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

		body.Password = hashed
	}

	if request.Role != "" {
		body.Role = request.Role
	}

	user, err := h.service.Update(c, *body)

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
			Message: "User Created Successfully",
		},
		Data: user,
	})
}
