package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"mime/multipart"
	"ranggaAdiPratama/go-with-claude/internal/database"
	"ranggaAdiPratama/go-with-claude/internal/requests"
	"ranggaAdiPratama/go-with-claude/internal/responses"
	"ranggaAdiPratama/go-with-claude/internal/utils"

	"github.com/google/uuid"
)

type ShopService struct {
	store      *database.Store
	cloudinary *utils.CloudinaryService
}

func NewShopService(store *database.Store, cloudinary *utils.CloudinaryService) *ShopService {
	return &ShopService{
		store:      store,
		cloudinary: cloudinary,
	}
}

func (s *ShopService) Index(ctx context.Context, params database.ShopListParams) (responses.ShopDisplayResponse, error) {
	var resp responses.ShopDisplayResponse

	shops, err := s.store.ShopList(ctx, params)

	if err != nil {
		fmt.Println(err)

		return resp, err
	}

	countParam := database.ShopListTotalParams{
		Rank:   params.Rank,
		Search: params.Search,
	}

	total, err := s.store.ShopListTotal(ctx, countParam)

	if err != nil {
		fmt.Println(err)

		return resp, err
	}

	data := make([]responses.ShopResponseShort, len(shops))

	for i, shop := range shops {
		data[i] = responses.ShopResponseShort{
			ID:   shop.ID,
			Name: shop.Name,
			Logo: shop.Logo,
			Rank: shop.Rank,
			Slug: shop.Slug,
		}
	}

	totalPages := (total + int64(params.Till) - 1) / int64(params.Till)

	currentPage := int32((params.Page / params.Till) + 1)

	resp = responses.ShopDisplayResponse{
		Data: data,
		Pagination: responses.PaginationResponse{
			Total:       int32(total),
			CurrentPage: currentPage,
			Pages:       int32(totalPages),
			Limit:       params.Till,
		}}

	return resp, nil
}

func (s *ShopService) IndexNoPagination(ctx context.Context, params database.ShopListParams) ([]responses.ShopResponseShort, error) {
	shops, err := s.store.ShopList(ctx, params)

	if err != nil {
		fmt.Println(err)

		return nil, err
	}

	shopData := make([]responses.ShopResponseShort, len(shops))

	for i, shop := range shops {
		shopData[i] = responses.ShopResponseShort{
			ID:   shop.ID,
			Name: shop.Name,
			Logo: shop.Logo,
			Rank: shop.Rank,
			Slug: shop.Slug,
		}
	}

	return shopData, nil
}

func (s *ShopService) Store(ctx context.Context, userID uuid.UUID, body requests.StoreShopRequest, logoFile multipart.File, logoHeader *multipart.FileHeader) (responses.ShopResponseForUser, error) {
	var resp responses.ShopResponseForUser

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

	existingShop, err := s.store.GetShopByUserId(ctx, userID)

	if err != nil && err != sql.ErrNoRows {
		return resp, err
	}

	if existingShop.ID != uuid.Nil {
		return resp, errors.New("user already has a shop")
	}

	shopByName, err := s.store.GetShopByName(ctx, body.Name)

	if err != nil && err != sql.ErrNoRows {
		return resp, err
	}

	if shopByName.ID != uuid.Nil {
		return resp, errors.New("shop name is already taken")
	}

	if err := utils.ValidateImageFile(logoHeader); err != nil {
		return resp, err
	}

	logoURL, _, err := s.cloudinary.UploadImage(ctx, logoFile, logoHeader.Filename)

	if err != nil {
		fmt.Println("Error uploading image to Cloudinary:", err)

		return resp, errors.New("failed to upload logo")
	}

	shop, err := q.StoreShop(ctx, database.StoreShopParams{
		UserID: userID,
		Name:   body.Name,
		Logo:   logoURL,
		Slug:   utils.GenerateSlug(body.Name),
	})

	if err != nil {
		fmt.Println("Error storing shop in database:", err)

		_ = s.cloudinary.DeleteImage(ctx, utils.ExtractPublicID(logoURL))

		return resp, err
	}

	resp = responses.ShopResponseForUser{
		ID:        shop.ID,
		Name:      shop.Name,
		Logo:      shop.Logo,
		Rank:      shop.Rank,
		Slug:      shop.Slug,
		CreatedAt: shop.CreatedAt.Format("2006-01-02T15:04:05Z"),
		UpdatedAt: shop.UpdatedAt.Format("2006-01-02T15:04:05Z"),
	}

	return resp, nil
}

func (s *ShopService) UpdatePersonal(ctx context.Context, userID uuid.UUID, req requests.StoreShopRequest, logoFile multipart.File, logoHeader *multipart.FileHeader) (responses.ShopResponseForUser, error) {
	var resp responses.ShopResponseForUser

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

	oldData, err := s.store.GetShopByUserId(ctx, userID)

	if err != nil {
		return resp, err
	}

	id := oldData.ID

	shopByName, err := q.GetShopByName(ctx, req.Name)

	if err != nil && err != sql.ErrNoRows {
		return resp, err
	}

	if shopByName.ID != uuid.Nil && shopByName.ID != id {
		return resp, errors.New("shop name is already taken")
	}

	newLogoURL := ""

	body := database.UpdateShopParams{
		ID:   id,
		Name: req.Name,
		Logo: oldData.Logo,
		Slug: utils.GenerateSlug(req.Name),
	}

	if logoFile != nil && logoHeader != nil {
		logoURL, _, err := s.cloudinary.UploadImage(ctx, logoFile, logoHeader.Filename)

		if err != nil {
			fmt.Println("Error uploading image to Cloudinary:", err)

			return resp, errors.New("failed to upload logo")
		}

		body.Logo = logoURL
		newLogoURL = logoURL
	}

	shop, err := q.UpdateShop(ctx, body)

	if err != nil {
		fmt.Println("Error storing shop in database:", err)

		if newLogoURL != "" {
			_ = s.cloudinary.DeleteImage(ctx, utils.ExtractPublicID(newLogoURL))
		}

		return resp, err
	}

	if newLogoURL != "" {
		_ = s.cloudinary.DeleteImage(ctx, utils.ExtractPublicID(oldData.Logo))
	}

	resp = responses.ShopResponseForUser{
		ID:        shop.ID,
		Name:      shop.Name,
		Logo:      shop.Logo,
		Rank:      shop.Rank,
		Slug:      shop.Slug,
		CreatedAt: shop.CreatedAt.Format("2006-01-02T15:04:05Z"),
		UpdatedAt: shop.UpdatedAt.Format("2006-01-02T15:04:05Z"),
	}

	return resp, nil
}
