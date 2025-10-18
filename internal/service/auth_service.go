package service

import (
	"context"
	"errors"
	"fmt"
	"ranggaAdiPratama/go-with-claude/internal/config"
	"ranggaAdiPratama/go-with-claude/internal/database"
	"ranggaAdiPratama/go-with-claude/internal/requests"
	"ranggaAdiPratama/go-with-claude/internal/responses"
	"ranggaAdiPratama/go-with-claude/internal/utils"
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

	refreshtoken, _, err := s.pasetoMaker.CreateToken(
		user.ID, user.Username, user.Email, user.Role, s.config.RefreshTokenTTL,
	)

	if err != nil {
		return resp, errors.New("failed to create token")
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
