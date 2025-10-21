package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"ranggaAdiPratama/go-with-claude/internal/requests"
	"ranggaAdiPratama/go-with-claude/internal/responses"
	"ranggaAdiPratama/go-with-claude/internal/service"
	"ranggaAdiPratama/go-with-claude/internal/utils"

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
