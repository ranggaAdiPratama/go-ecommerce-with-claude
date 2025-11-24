package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"ranggaAdiPratama/go-with-claude/internal/database"
	"ranggaAdiPratama/go-with-claude/internal/requests"
	"ranggaAdiPratama/go-with-claude/internal/responses"
	"ranggaAdiPratama/go-with-claude/internal/utils"

	"github.com/google/uuid"
)

type CategoryService struct {
	store *database.Store
}

func NewCategoryService(store *database.Store) *CategoryService {
	return &CategoryService{
		store: store,
	}
}

func (s *CategoryService) Index(ctx context.Context, params database.CategoryListParams) (responses.CategoryPaginatedResponse, error) {
	var resp responses.CategoryPaginatedResponse

	categories, err := s.store.CategoryList(ctx, params)

	if err != nil {
		fmt.Println(err)

		return resp, err
	}

	countParam := database.CategoryListTotalParams{
		Search: params.Search,
		Status: params.Status,
	}

	total, err := s.store.CategoryListTotal(ctx, countParam)

	if err != nil {
		fmt.Println(err)

		return resp, err
	}

	data := make([]responses.CategoryResponse, len(categories))

	for i, category := range categories {
		data[i] = responses.CategoryResponse{
			ID:        category.ID,
			Name:      category.Name,
			Icon:      category.Icon,
			Slug:      category.Slug,
			IsActive:  category.IsActive,
			CreatedAt: category.CreatedAt.Format("2006-01-02T15:04:05Z"),
			UpdatedAt: category.UpdatedAt.Format("2006-01-02T15:04:05Z"),
		}
	}

	totalPages, currentPage := utils.Paginator(total, params.Till, params.Page)

	resp = responses.CategoryPaginatedResponse{
		Data: data,
		Pagination: responses.PaginationResponse{
			Total:       int32(total),
			CurrentPage: currentPage,
			Pages:       int32(totalPages),
			Limit:       params.Till,
		}}

	return resp, nil
}

func (s *CategoryService) GetBySlug(ctx context.Context, slug string) (responses.CategoryResponseShort, error) {
	var resp responses.CategoryResponseShort

	category, err := s.store.GetCategoryBySlug(ctx, slug)

	if err != nil {
		fmt.Println(err)

		return resp, errors.New("no data found")
	}

	if !category.IsActive {
		return resp, errors.New("no data found")
	}

	resp = responses.CategoryResponseShort{
		ID:   category.ID,
		Name: category.Name,
		Icon: category.Icon,
		Slug: category.Slug,
	}

	return resp, nil
}

func (s *CategoryService) IndexNoPagination(ctx context.Context, params database.CategoryListParams) ([]responses.CategoryResponseShort, error) {
	categories, err := s.store.CategoryList(ctx, params)

	if err != nil {
		fmt.Println(err)

		return nil, err
	}

	categoryData := make([]responses.CategoryResponseShort, len(categories))

	for i, category := range categories {
		categoryData[i] = responses.CategoryResponseShort{
			ID:   category.ID,
			Name: category.Name,
			Icon: category.Icon,
			Slug: category.Slug,
		}
	}

	return categoryData, nil
}

func (s *CategoryService) Store(ctx context.Context, body requests.StoreCategoryRequest) (resp *responses.CategoryResponse, err error) {
	tx, err := s.store.BeginTx(ctx, nil)

	if err != nil {
		return resp, errors.New("failed to start transaction")
	}

	q := s.store.Queries.WithTx(tx)

	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback()
			err = fmt.Errorf("panic occurred: %v", p)
			return
		}

		if err != nil {
			fmt.Println(err)

			_ = tx.Rollback()
			return
		}

		if commitErr := tx.Commit(); commitErr != nil {
			err = fmt.Errorf("failed to commit transaction: %w", commitErr)
		}
	}()

	existingCategory, err := s.store.GetCategoryByName(ctx, body.Name)

	if err != nil && err != sql.ErrNoRows {
		return resp, err
	}

	if existingCategory.ID != uuid.Nil {
		return resp, errors.New("category already exists")
	}

	active := false

	if body.IsActive != "" && body.IsActive == "1" {
		active = true
	}

	category, err := q.StoreCategory(ctx, database.StoreCategoryParams{
		Name:     body.Name,
		Icon:     body.Icon,
		Slug:     utils.GenerateSlug(body.Name),
		IsActive: active,
	})

	if err != nil {
		fmt.Println("Error storing category in database:", err)

		return resp, err
	}

	resp = &responses.CategoryResponse{
		ID:        category.ID,
		Name:      category.Name,
		Icon:      category.Icon,
		Slug:      category.Slug,
		IsActive:  category.IsActive,
		CreatedAt: category.CreatedAt.Format("2006-01-02T15:04:05Z"),
		UpdatedAt: category.UpdatedAt.Format("2006-01-02T15:04:05Z"),
	}

	return resp, err
}

func (s *CategoryService) Update(ctx context.Context, body database.UpdateCategoryParams) (resp *responses.CategoryResponse, err error) {
	tx, err := s.store.BeginTx(ctx, nil)

	if err != nil {
		return resp, errors.New("failed to start transaction")
	}

	q := s.store.Queries.WithTx(tx)

	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback()
			err = fmt.Errorf("panic occurred: %v", p)
			resp = nil
			return
		}

		if err != nil {
			fmt.Println(err)

			_ = tx.Rollback()
			return
		}

		if commitErr := tx.Commit(); commitErr != nil {
			err = fmt.Errorf("failed to commit transaction: %w", commitErr)
			resp = nil
		}
	}()

	existingCategory, err := s.store.GetCategoryByName(ctx, body.Name)

	if err != nil && err != sql.ErrNoRows {
		return resp, err
	}

	if existingCategory.ID != uuid.Nil && existingCategory.ID != body.ID {
		return resp, errors.New("category already exists")
	}

	category, err := q.UpdateCategory(ctx, body)

	if err != nil {
		fmt.Println(err)

		return resp, err
	}

	resp = &responses.CategoryResponse{
		ID:        category.ID,
		Name:      category.Name,
		Icon:      category.Icon,
		Slug:      category.Slug,
		IsActive:  category.IsActive,
		CreatedAt: category.CreatedAt.Format("2006-01-02T15:04:05Z"),
		UpdatedAt: category.UpdatedAt.Format("2006-01-02T15:04:05Z"),
	}

	return resp, err
}
