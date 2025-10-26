package handlers

import (
	"database/sql"
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

type SettingHandler struct {
	service *service.SettingService
}

func NewSettingHandler(service *service.SettingService) *SettingHandler {
	return &SettingHandler{
		service: service,
	}
}

func (h *SettingHandler) Index(c *gin.Context) {
	data, err := h.service.Index(c)

	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, responses.Response{
				MetaData: responses.MetaDataResponse{
					Code:    http.StatusNotFound,
					Message: "Setting is not set",
				},
			})
			return
		}

		fmt.Println(err)

		c.JSON(http.StatusInternalServerError, responses.Response{
			MetaData: responses.MetaDataResponse{
				Code:    http.StatusInternalServerError,
				Message: "Failed to fetch settings",
			},
		})

		return
	}

	c.JSON(http.StatusOK, responses.Response{
		MetaData: responses.MetaDataResponse{
			Code:    http.StatusOK,
			Message: "Settings successfullyFetched",
		},
		Data: data,
	})
}

func (h *SettingHandler) StoreOrUpdate(c *gin.Context) {
	var request requests.SettingRequest

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

	_, err := h.service.Index(c)
	isNewRecord := err == sql.ErrNoRows

	if err != nil && err != sql.ErrNoRows {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, responses.Response{
			MetaData: responses.MetaDataResponse{
				Code:    http.StatusInternalServerError,
				Message: "Failed to fetch setting",
			},
		})
		return
	}

	logoFile, logoHeader, fileErr := c.Request.FormFile("logo")

	if fileErr != nil && fileErr != http.ErrMissingFile {
		c.JSON(http.StatusBadRequest, responses.Response{
			MetaData: responses.MetaDataResponse{
				Code:    http.StatusBadRequest,
				Message: "Failed to read logo file",
			},
		})
		return
	}

	if isNewRecord && fileErr == http.ErrMissingFile {
		c.JSON(http.StatusBadRequest, responses.Response{
			MetaData: responses.MetaDataResponse{
				Code:    http.StatusBadRequest,
				Message: "Logo is required",
			},
		})
		return
	}

	if logoFile != nil {
		defer logoFile.Close()
	}

	data, err := h.service.StoreOrUpdate(c, request, logoFile, logoHeader)

	if err != nil {
		fmt.Println(err)

		statusCode := http.StatusInternalServerError
		message := "Failed to save setting"

		switch err.Error() {
		case "failed to upload logo":
			statusCode = http.StatusBadRequest
			message = "Failed to upload logo image"
		case "invalid file type. allowed: jpg, jpeg, png, gif, webp":
			statusCode = http.StatusBadRequest
			message = err.Error()
		case "file size exceeds 1MB limit":
			statusCode = http.StatusBadRequest
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

	successCode := http.StatusOK
	successMessage := "Setting successfully updated"

	c.JSON(successCode, responses.Response{
		MetaData: responses.MetaDataResponse{
			Code:    successCode,
			Message: successMessage,
		},
		Data: data,
	})
}
