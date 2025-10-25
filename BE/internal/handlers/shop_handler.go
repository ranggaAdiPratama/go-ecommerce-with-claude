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

type ShopHandler struct {
	service *service.ShopService
}

func NewShopHandler(service *service.ShopService) *ShopHandler {
	return &ShopHandler{
		service: service,
	}
}

func (h *ShopHandler) Index(c *gin.Context) {
	sort := c.DefaultQuery("sort", "created_at")
	order := c.DefaultQuery("order", "desc")
	limitString := c.DefaultQuery("limit", "15")
	rank := ""
	search := ""

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

	if search != "" {
		search = utils.EscapeRegex(search)
	}

	params := &database.ShopListParams{
		Page:      0,
		Rank:      rank,
		Search:    search,
		Sort:      sort,
		SortOrder: order,
		Till:      int32(limit),
	}

	shops, err := h.service.IndexNoPagination(c.Request.Context(), *params)

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
			Message: "Shops Retrieved Successfully",
		},
		Data: shops,
	})
}

func (h *ShopHandler) Store(c *gin.Context) {
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

	payload := userPayload.(*utils.TokenPayload)

	var request requests.StoreShopRequest

	if err := c.ShouldBind(&request); err != nil {
		fmt.Println(err)

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

	logoFile, logoHeader, err := c.Request.FormFile("logo")

	if err != nil {
		c.JSON(http.StatusBadRequest, responses.Response{
			MetaData: responses.MetaDataResponse{
				Code:    http.StatusBadRequest,
				Message: "Logo is required",
			},
		})
	}

	defer logoFile.Close()

	data, err := h.service.Store(c.Request.Context(), payload.ID, request, logoFile, logoHeader)

	if err != nil {
		fmt.Println(err)

		statusCode := http.StatusInternalServerError
		message := "Error :("

		if err.Error() == "user already has a shop" {
			statusCode = http.StatusForbidden

			message = "You already have a shop"
		} else if err.Error() == "shop name is already taken" {
			statusCode = http.StatusForbidden

			message = err.Error()
		}

		c.JSON(statusCode, responses.Response{
			MetaData: responses.MetaDataResponse{
				Code:    statusCode,
				Message: message,
			},
		})

		return
	}

	c.JSON(http.StatusCreated, responses.Response{
		MetaData: responses.MetaDataResponse{
			Code:    http.StatusCreated,
			Message: "Shop successfully created",
		},
		Data: data,
	})
}

func (h *ShopHandler) UpdatePersonal(c *gin.Context) {
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

	payload := userPayload.(*utils.TokenPayload)

	var request requests.StoreShopRequest

	if err := c.ShouldBind(&request); err != nil {
		fmt.Println(err)

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

	var data responses.ShopResponseForUser

	logoFile, logoHeader, err := c.Request.FormFile("logo")

	if logoFile != nil {
		defer logoFile.Close()
	}

	if err != nil {
		data, err = h.service.UpdatePersonal(c.Request.Context(), payload.ID, request, nil, nil)
	} else {
		data, err = h.service.UpdatePersonal(c.Request.Context(), payload.ID, request, logoFile, logoHeader)
	}

	if err != nil {
		fmt.Println(err)

		statusCode := http.StatusInternalServerError
		message := "Error :("

		if err.Error() == "user already has a shop" {
			statusCode = http.StatusForbidden

			message = "You already have a shop"
		} else if err.Error() == "shop name is already taken" {
			statusCode = http.StatusForbidden

			message = err.Error()
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
			Message: "Shop successfully updated",
		},
		Data: data,
	})
}
