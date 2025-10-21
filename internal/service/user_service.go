package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"ranggaAdiPratama/go-with-claude/internal/database"
	"ranggaAdiPratama/go-with-claude/internal/responses"

	"github.com/google/uuid"
)

type UserService struct {
	store *database.Store
}

func NewUserService(store *database.Store) *UserService {
	return &UserService{
		store: store,
	}
}

func (s *UserService) Index(ctx context.Context, params database.UserListParams) (responses.UserListResponse, error) {
	var resp responses.UserListResponse

	users, err := s.store.UserList(ctx, params)

	if err != nil {
		return resp, err
	}

	countParam := database.UserListTotalParams{
		Role:   params.Role,
		Search: params.Search,
	}

	total, err := s.store.UserListTotal(ctx, countParam)

	if err != nil {
		fmt.Println(err)

		return resp, err
	}

	userData := make([]responses.UserResponse, len(users))

	for i, user := range users {
		userData[i] = responses.UserResponse{
			ID:        user.ID,
			Name:      user.Name,
			Username:  user.Username,
			Email:     user.Email,
			Role:      user.Role,
			CreatedAt: user.CreatedAt.Format("2006-01-02T15:04:07Z"),
			UpdatedAt: user.UpdatedAt.Format("2006-01-02T15:04:07Z"),
		}
	}

	totalPages := (total + int64(params.Till) - 1) / int64(params.Till)

	currentPage := int32((params.Page / params.Till) + 1)

	resp = responses.UserListResponse{
		Data: userData,
		Pagination: responses.PaginationResponse{
			Total:       int32(total),
			CurrentPage: currentPage,
			Pages:       int32(totalPages),
			Limit:       params.Till,
		}}

	return resp, nil
}

func (s *UserService) Destroy(ctx context.Context, id uuid.UUID) (err error) {
	tx, err := s.store.BeginTx(ctx, nil)

	if err != nil {
		return errors.New("failed to start transaction")
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

	_, err = s.store.GetUserById(ctx, id)

	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("user not found")
		}

		return err
	}

	err = q.DeleteUser(ctx, id)

	if err != nil {
		fmt.Println(err)

		return err
	}

	return nil
}

func (s *UserService) Show(ctx context.Context, id uuid.UUID) (*responses.UserResponse, error) {
	user, err := s.store.GetUserById(ctx, id)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}

		return nil, err
	}

	resp := &responses.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Username:  user.Username,
		Email:     user.Email,
		Role:      user.Role,
		CreatedAt: user.CreatedAt.Format("2006-01-02T15:04:05Z"),
		UpdatedAt: user.UpdatedAt.Format("2006-01-02T15:04:05Z"),
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

	if emailTaken.ID != uuid.Nil {
		return resp, errors.New("e-mail already taken")
	}

	usernameTaken, err := s.store.GetUserByUsername(ctx, body.Username)

	if err != nil && err != sql.ErrNoRows {
		return resp, err
	}

	if usernameTaken.ID != uuid.Nil {
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

func (s *UserService) Update(ctx context.Context, body database.UpdateUserParams) (resp *responses.UserResponse, err error) {
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

	_, err = s.store.GetUserById(ctx, body.ID)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}

		return nil, err
	}

	emailTaken, err := s.store.GetUserByEmail(ctx, body.Email)

	if err != nil && err != sql.ErrNoRows {
		return resp, err
	}

	if emailTaken.ID != uuid.Nil && emailTaken.ID != body.ID {
		return resp, errors.New("e-mail already taken")
	}

	usernameTaken, err := s.store.GetUserByUsername(ctx, body.Username)

	if err != nil && err != sql.ErrNoRows {
		return resp, err
	}

	if usernameTaken.ID != uuid.Nil && usernameTaken.ID != body.ID {
		return resp, errors.New("username already taken")
	}

	user, err := q.UpdateUser(ctx, body)

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
