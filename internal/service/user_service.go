package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"ranggaAdiPratama/go-with-claude/internal/database"
	"ranggaAdiPratama/go-with-claude/internal/responses"
)

type UserService struct {
	store *database.Store
}

func NewUserService(store *database.Store) *UserService {
	return &UserService{
		store: store,
	}
}

func (s *UserService) Index(ctx context.Context, params database.UserListParams) ([]responses.UserResponse, error) {
	fmt.Println(params)
	users, err := s.store.UserList(ctx, params)

	if err != nil {
		return nil, err
	}

	resp := make([]responses.UserResponse, len(users))

	fmt.Println(users)

	for i, user := range users {
		resp[i] = responses.UserResponse{
			ID:        user.ID,
			Name:      user.Name,
			Username:  user.Username,
			Email:     user.Email,
			Role:      user.Role,
			CreatedAt: user.CreatedAt.Format("2006-01-02T15:04:07Z"),
			UpdatedAt: user.UpdatedAt.Format("2006-01-02T15:04:07Z"),
		}
	}

	return resp, nil
}

func (s *UserService) Store(ctx context.Context, body database.StoreUserParams) (resp *responses.UserResponse, err error) {
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

	emailTaken, err := s.store.GetUserByEmail(ctx, body.Email)

	if err != nil && err != sql.ErrNoRows {
		return resp, err
	}

	if emailTaken.ID != 0 {
		return resp, errors.New("username already taken")
	}

	usernameTaken, err := s.store.GetUserByUsername(ctx, body.Username)

	if err != nil && err != sql.ErrNoRows {
		return resp, err
	}

	if usernameTaken.ID != 0 {
		return resp, errors.New("username already taken")
	}

	user, err := q.StoreUser(ctx, body)

	if err != nil {
		fmt.Println(err)

		return resp, err
	}

	resp = &responses.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Username:  user.Username,
		Email:     user.Email,
		Role:      user.Role,
		CreatedAt: user.CreatedAt.Format("2006-01-02T15:04:05Z"),
		UpdatedAt: user.UpdatedAt.Format("2006-01-02T15:04:05Z"),
	}

	return resp, err
}
