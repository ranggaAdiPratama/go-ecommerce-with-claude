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

type CategoryHandler struct {
	service *service.CategoryService
}

func NewCategoryHandler(service *service.CategoryService) *CategoryHandler {
	return &CategoryHandler{
		service: service,
	}
}

func (h *CategoryHandler) Index(c *gin.Context) {
	sort := c.DefaultQuery("sort", "name")
	order := c.DefaultQuery("order", "asc")
	limitString := c.DefaultQuery("limit", "15")
	status := c.DefaultQuery("status", "1")
	search := ""

	if status != "1" && status != "0" {
		c.JSON(http.StatusInternalServerError, responses.Response{
			MetaData: responses.MetaDataResponse{
				Code:    http.StatusInternalServerError,
				Message: "Error defining status",
			},
		})

		return
	}

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

	params := &database.CategoryListParams{
		Page:      0,
		Status:    status,
		Search:    search,
		Sort:      sort,
		SortOrder: order,
		Till:      int32(limit),
	}

	categories, err := h.service.IndexNoPagination(c.Request.Context(), *params)

	if err != nil {
		fmt.Println(err)

		c.JSON(http.StatusInternalServerError, responses.Response{
			MetaData: responses.MetaDataResponse{
				Code:    http.StatusInternalServerError,
				Message: "Failed to fetch categories",
			},
		})

		return
	}

	c.JSON(http.StatusOK, responses.Response{
		MetaData: responses.MetaDataResponse{
			Code:    http.StatusOK,
			Message: "Categories Retrieved Successfully",
		},
		Data: categories,
	})
}

func (h *CategoryHandler) DataTable(c *gin.Context) {
	limitString := c.DefaultQuery("limit", "15")
	order := c.DefaultQuery("order", "desc")
	pageString := c.DefaultQuery("page", "1")
	sort := c.DefaultQuery("sort", "created_at")
	status := c.DefaultQuery("status", "1")
	search := ""

	if status != "1" && status != "0" {
		c.JSON(http.StatusInternalServerError, responses.Response{
			MetaData: responses.MetaDataResponse{
				Code:    http.StatusInternalServerError,
				Message: "Error defining status",
			},
		})

		return
	}

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

	params := &database.CategoryListParams{
		Page:      int32((page - 1) * limit),
		Status:    status,
		Search:    search,
		Sort:      sort,
		SortOrder: order,
		Till:      int32(limit),
	}

	categories, err := h.service.Index(c.Request.Context(), *params)

	if err != nil {
		fmt.Println(err)

		c.JSON(http.StatusInternalServerError, responses.Response{
			MetaData: responses.MetaDataResponse{
				Code:    http.StatusInternalServerError,
				Message: "Failed to fetch categories",
			},
		})

		return
	}

	c.JSON(http.StatusOK, responses.Response{
		MetaData: responses.MetaDataResponse{
			Code:    http.StatusOK,
			Message: "Categories Retrieved Successfully",
		},
		Data: categories,
	})
}

func (h *CategoryHandler) GetBySlug(c *gin.Context) {
	slug := c.Param("slug")

	category, err := h.service.GetBySlug(c, slug)

	if err != nil {
		if err.Error() == "no data found" {
			c.JSON(http.StatusNotFound, responses.Response{
				MetaData: responses.MetaDataResponse{
					Code:    http.StatusNotFound,
					Message: utils.CapitalizeFirst(err.Error()),
				},
			})

			return
		}

		fmt.Println(err)

		c.JSON(http.StatusInternalServerError, responses.Response{
			MetaData: responses.MetaDataResponse{
				Code:    http.StatusInternalServerError,
				Message: "Failed in fetching data",
			},
		})

		return
	}

	c.JSON(http.StatusOK, responses.Response{
		MetaData: responses.MetaDataResponse{
			Code:    http.StatusOK,
			Message: "Category Retrieved Successfully",
		},
		Data: category,
	})
}

func (h *CategoryHandler) Store(c *gin.Context) {
	var request requests.StoreCategoryRequest

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

	category, err := h.service.Store(c, request)

	if err != nil {
		c.JSON(http.StatusBadRequest, responses.Response{
			MetaData: responses.MetaDataResponse{
				Code:    http.StatusBadRequest,
				Message: utils.CapitalizeFirst(err.Error()),
			},
		})

		return
	}

	c.JSON(http.StatusCreated, responses.Response{
		MetaData: responses.MetaDataResponse{
			Code:    http.StatusCreated,
			Message: "Category Created Successfully",
		},
		Data: category,
	})
}

func (h *CategoryHandler) Update(c *gin.Context) {
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

	var request requests.UpdateCategoryRequest

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

	body := &database.UpdateCategoryParams{
		Name: request.Name,
		Icon: request.Icon,
		Slug: utils.GenerateSlug(request.Name),
		ID:   id,
	}

	category, err := h.service.Update(c.Request.Context(), *body)

	if err != nil {
		c.JSON(http.StatusBadRequest, responses.Response{
			MetaData: responses.MetaDataResponse{
				Code:    http.StatusBadRequest,
				Message: utils.CapitalizeFirst(err.Error()),
			},
		})

		return
	}

	c.JSON(http.StatusCreated, responses.Response{
		MetaData: responses.MetaDataResponse{
			Code:    http.StatusCreated,
			Message: "Category Updated Successfully",
		},
		Data: category,
	})
}
