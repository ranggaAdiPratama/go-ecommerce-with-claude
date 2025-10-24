package utils

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"net/url"
	"path/filepath"
	"strings"
	"time"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

type CloudinaryService struct {
	cid    *cloudinary.Cloudinary
	folder string
}

func NewCloudinaryService(cloudName, apiKey, apiSecret, folder string) (*CloudinaryService, error) {
	cid, err := cloudinary.NewFromParams(cloudName, apiKey, apiSecret)

	if err != nil {
		return nil, fmt.Errorf("failed to initialize Cloudinary: %w", err)
	}

	return &CloudinaryService{
		cid:    cid,
		folder: folder,
	}, nil
}

func (s *CloudinaryService) DeleteImage(ctx context.Context, publicID string) error {
	_, err := s.cid.Upload.Destroy(ctx, uploader.DestroyParams{
		PublicID:     publicID,
		ResourceType: "image",
	})

	if err != nil {
		return fmt.Errorf("failed to delete image: %w", err)
	}

	return nil
}

func (s *CloudinaryService) UploadImage(ctx context.Context, file multipart.File, filename string) (string, string, error) {
	if seeker, ok := file.(io.Seeker); ok {
		_, err := seeker.Seek(0, io.SeekStart)

		if err != nil {
			return "", "", fmt.Errorf("failed to reset file pointer: %w", err)
		}
	}

	ext := filepath.Ext(filename)

	uniqueFileName := fmt.Sprintf("%d_%s%s", time.Now().Unix(), strings.TrimSuffix(filename, ext), ext)

	fmt.Printf("Uploading to Cloudinary - Folder: %s, Filename: %s\n", s.folder, uniqueFileName)

	uploadResult, err := s.cid.Upload.Upload(ctx, file, uploader.UploadParams{
		PublicID:       uniqueFileName,
		Folder:         s.folder,
		ResourceType:   "image",
		Transformation: "q_auto,f_auto",
	})

	if err != nil {
		return "", "", fmt.Errorf("failed to upload image: %w", err)
	}

	fmt.Printf("Successfully uploaded to Cloudinary: %s\n", uploadResult.SecureURL)
	fmt.Printf("Public ID: %s\n", uploadResult.PublicID)

	return uploadResult.SecureURL, uploadResult.PublicID, nil
}

func ExtractPublicID(urlStr string) string {
	urlStr = strings.TrimSpace(urlStr)

	if urlStr == "" {
		fmt.Println("ExtractPublicID: Empty URL provided")
		return ""
	}

	parts := strings.Split(urlStr, "/upload/")

	if len(parts) != 2 {
		fmt.Printf("ExtractPublicID: Invalid URL format: %s\n", urlStr)
		return ""
	}

	afterUpload := parts[1]
	pathParts := strings.Split(afterUpload, "/")

	if len(pathParts) < 2 {
		fmt.Printf("ExtractPublicID: Insufficient path segments: %s\n", urlStr)
		return ""
	}

	publicIDWithExt := strings.Join(pathParts[1:], "/")

	decodedPublicID, err := url.QueryUnescape(publicIDWithExt)
	if err != nil {
		fmt.Printf("ExtractPublicID: Error decoding URL: %v\n", err)
		decodedPublicID = publicIDWithExt
	}

	publicID := strings.TrimSuffix(decodedPublicID, filepath.Ext(decodedPublicID))

	fmt.Printf("ExtractPublicID: URL=%s -> PublicID=%s\n", urlStr, publicID)

	return publicID
}

func ValidateImageFile(fileHeader *multipart.FileHeader) error {
	const maxFileSize = 1024 * 1024 // 1 mb aja, gratisan aing makena, cuk :(

	if fileHeader.Size > maxFileSize {
		return fmt.Errorf("file size exceeds the maximum limit of 5MB")
	}

	ext := strings.ToLower(filepath.Ext(fileHeader.Filename))

	validExts := map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
		".gif":  true,
		".webp": true,
	}

	if !validExts[ext] {
		return fmt.Errorf("invalid file type: %s", ext)
	}

	contentType := fileHeader.Header.Get("Content-Type")

	validMimes := map[string]bool{
		"image/jpeg": true,
		"image/jpg":  true,
		"image/png":  true,
		"image/gif":  true,
		"image/webp": true,
	}

	if !validMimes[contentType] {
		return fmt.Errorf("invalid content type: %s", contentType)
	}

	return nil
}
