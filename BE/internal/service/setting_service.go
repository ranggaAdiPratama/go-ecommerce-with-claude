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
)

type SettingService struct {
	store      *database.Store
	cloudinary *utils.CloudinaryService
}

func NewSettingService(store *database.Store, cloudinary *utils.CloudinaryService) *SettingService {
	return &SettingService{
		store:      store,
		cloudinary: cloudinary,
	}
}

func (s *SettingService) Index(ctx context.Context) (responses.SettingResponse, error) {
	var resp responses.SettingResponse

	setting, err := s.store.GetSetting(ctx)

	if err != nil {
		return resp, err
	}

	resp = responses.SettingResponse{
		Name: setting.Name,
		Logo: setting.Logo,
	}

	return resp, nil
}

func (s *SettingService) StoreOrUpdate(ctx context.Context, body requests.SettingRequest, logoFile multipart.File, logoHeader *multipart.FileHeader) (responses.SettingResponse, error) {
	var resp responses.SettingResponse

	tx, err := s.store.BeginTx(ctx, nil)
	if err != nil {
		return resp, errors.New("failed to start transaction")
	}

	committed := true

	defer func() {
		if !committed {
			tx.Rollback()
			fmt.Println("Transaction rolled back")
		}
	}()

	q := s.store.Queries.WithTx(tx)

	existing, err := q.GetSetting(ctx)

	if err != nil && err != sql.ErrNoRows {
		committed = false

		return resp, err
	}

	if err == sql.ErrNoRows {
		if err := utils.ValidateImageFile(logoHeader); err != nil {
			committed = false

			return resp, err
		}

		logoURL, _, err := s.cloudinary.UploadImage(ctx, logoFile, logoHeader.Filename)

		if err != nil {
			fmt.Println("Error uploading image to Cloudinary:", err)

			committed = false

			return resp, errors.New("failed to upload logo")
		}

		setting, err := q.StoreSetting(ctx, database.StoreSettingParams{
			Name: body.Name,
			Logo: logoURL,
		})

		if err != nil {
			fmt.Println("Error storing shop in database:", err)

			committed = false

			_ = s.cloudinary.DeleteImage(ctx, utils.ExtractPublicID(logoURL))

			return resp, err
		}

		resp = responses.SettingResponse{
			Name: setting.Name,
			Logo: setting.Logo,
		}

		if err := tx.Commit(); err != nil {
			committed = false

			return resp, fmt.Errorf("failed to commit: %w", err)
		}

		fmt.Println("Transaction committed")

		return resp, nil
	}

	newLogoURL := ""

	updateParam := database.UpdateSettingParams{
		ID:   existing.ID,
		Name: body.Name,
		Logo: existing.Logo,
	}

	if logoFile != nil && logoHeader != nil {
		logoURL, _, err := s.cloudinary.UploadImage(ctx, logoFile, logoHeader.Filename)

		if err != nil {
			committed = false

			fmt.Println("Error uploading image to Cloudinary:", err)

			return resp, errors.New("failed to upload logo")
		}

		updateParam.Logo = logoURL
		newLogoURL = logoURL
	}

	setting, err := q.UpdateSetting(ctx, updateParam)

	if err != nil {
		committed = false

		fmt.Println("Error storing shop in database:", err)

		if newLogoURL != "" {
			_ = s.cloudinary.DeleteImage(ctx, utils.ExtractPublicID(newLogoURL))
		}

		return resp, err
	}

	if newLogoURL != "" {
		_ = s.cloudinary.DeleteImage(ctx, utils.ExtractPublicID(existing.Logo))
	}

	resp = responses.SettingResponse{
		Name: setting.Name,
		Logo: setting.Logo,
	}

	if err := tx.Commit(); err != nil {
		committed = false

		return resp, fmt.Errorf("failed to commit: %w", err)
	}

	fmt.Println("Transaction committed")

	return resp, nil

}
