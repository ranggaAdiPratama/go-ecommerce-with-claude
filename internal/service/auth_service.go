package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"ranggaAdiPratama/go-with-claude/internal/config"
	"ranggaAdiPratama/go-with-claude/internal/database"
	"ranggaAdiPratama/go-with-claude/internal/requests"
	"ranggaAdiPratama/go-with-claude/internal/responses"
	"ranggaAdiPratama/go-with-claude/internal/utils"
	"time"

	"github.com/google/uuid"
)

type AuthService struct {
	store       *database.Store
	pasetoMaker *utils.PasetoMaker
	config      *config.Config
}

func NewAuthService(
	store *database.Store, pasetoMaker *utils.PasetoMaker, cfg *config.Config) *AuthService {
	return &AuthService{
		store:       store,
		pasetoMaker: pasetoMaker,
		config:      cfg,
	}
}

func (s *AuthService) Login(ctx context.Context, body requests.LoginRequest) (responses.LoginResponse, error) {
	var resp responses.LoginResponse
	var user database.User
	var err error

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

	user, err = q.GetUserByUsername(ctx, body.Username)

	if err != nil {
		fmt.Println(err)

		return resp, err
	}

	if err := utils.CheckPassword(user.Password, body.Password); err != nil {
		return resp, errors.New("invalid credentials")
	}

	token, _, err := s.pasetoMaker.CreateToken(
		user.ID, user.Username, user.Email, user.Role, s.config.AccessTokenTTL,
	)

	if err != nil {
		return resp, errors.New("failed to create token")
	}

	_ = q.RevokeAllUserRefreshTokens(ctx, user.ID)

	refreshtoken, _, err := s.pasetoMaker.CreateToken(
		user.ID, user.Username, user.Email, user.Role, s.config.RefreshTokenTTL,
	)

	if err != nil {
		return resp, errors.New("failed to create token")
	}

	refreshTokenParams := database.StoreRefreshTokenParams{
		UserID:    user.ID,
		TokenHash: utils.HashToken(refreshtoken),
		ExpiresAt: time.Now().Add(s.config.RefreshTokenTTL),
	}

	_, err = q.StoreRefreshToken(ctx, refreshTokenParams)

	if err != nil {
		return resp, errors.New("failed storing token")
	}

	resp = responses.LoginResponse{
		User: responses.UserResponse{
			ID:        user.ID,
			Name:      user.Name,
			Username:  user.Username,
			Email:     user.Email,
			Role:      user.Role,
			CreatedAt: user.CreatedAt.Format("2006-01-02T15:04:05Z"),
			UpdatedAt: user.UpdatedAt.Format("2006-01-02T15:04:05Z"),
		},
		Token:        token,
		RefreshToken: refreshtoken,
	}

	return resp, nil
}

func (s *AuthService) Logout(ctx context.Context, id uuid.UUID) error {
	var err error

	tx, err := s.store.BeginTx(ctx, nil)

	if err != nil {
		return errors.New("failed to start transaction")
	}

	q := s.store.Queries.WithTx(tx)

	err = q.RevokeAllUserRefreshTokens(ctx, id)

	if err != nil {
		fmt.Println(err)

		return errors.New("failed revoking all tokens")
	}

	return nil
}

func (s *AuthService) RefreshToken(ctx context.Context, body requests.RefreshTokenRequest) (responses.LoginResponse, error) {
	var resp responses.LoginResponse
	var user database.User
	var err error

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

	tokenPayload, err := s.pasetoMaker.VerifyToken(body.Token)

	if err != nil {
		return resp, errors.New("invalid token")
	}

	tokenHash := utils.HashToken(body.Token)

	storedToken, err := q.GetRefreshToken(ctx, tokenHash)

	if err != nil {
		if err == sql.ErrNoRows {
			return resp, errors.New("refresh token not found")
		}

		return resp, err
	}

	if storedToken.IsRevoked {
		_ = q.RevokeAllUserRefreshTokens(ctx, tokenPayload.ID)

		return resp, errors.New("refresh token has been revoked")
	}

	if time.Now().After(storedToken.ExpiresAt) {
		_ = q.RevokeRefreshToken(ctx, tokenHash)

		return resp, errors.New("refresh token has expired")
	}

	user, err = q.GetUserById(ctx, tokenPayload.ID)

	if err != nil {
		fmt.Println(err)

		return resp, err
	}

	err = q.RevokeRefreshToken(ctx, tokenHash)

	if err != nil {
		return resp, errors.New("failed to revoke old token")
	}

	token, _, err := s.pasetoMaker.CreateToken(
		user.ID, user.Username, user.Email, user.Role, s.config.AccessTokenTTL,
	)

	if err != nil {
		return resp, errors.New("failed to create token")
	}

	refreshtoken, _, err := s.pasetoMaker.CreateToken(
		user.ID, user.Username, user.Email, user.Role, s.config.RefreshTokenTTL,
	)

	if err != nil {
		return resp, errors.New("failed to create token")
	}

	refreshTokenParams := database.StoreRefreshTokenParams{
		UserID:    user.ID,
		TokenHash: utils.HashToken(refreshtoken),
		ExpiresAt: time.Now().Add(s.config.RefreshTokenTTL),
	}

	_, err = q.StoreRefreshToken(ctx, refreshTokenParams)

	if err != nil {
		return resp, errors.New("failed storing token")
	}

	resp = responses.LoginResponse{
		User: responses.UserResponse{
			ID:        user.ID,
			Name:      user.Name,
			Username:  user.Username,
			Email:     user.Email,
			Role:      user.Role,
			CreatedAt: user.CreatedAt.Format("2006-01-02T15:04:05Z"),
			UpdatedAt: user.UpdatedAt.Format("2006-01-02T15:04:05Z"),
		},
		Token:        token,
		RefreshToken: refreshtoken,
	}

	return resp, nil
}
