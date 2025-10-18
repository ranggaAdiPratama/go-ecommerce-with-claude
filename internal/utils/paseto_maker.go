package utils

import (
	"fmt"
	"time"

	"aidanwoods.dev/go-paseto"
	"github.com/google/uuid"
)

type PasetoMaker struct {
	symmetricKey paseto.V4SymmetricKey
}

type TokenPayload struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

func NewPasetoMaker(symmetricKey string) (*PasetoMaker, error) {
	if len(symmetricKey) != 32 {
		return nil, fmt.Errorf("invalid key size: must be exactly 32 characters")
	}

	key, err := paseto.V4SymmetricKeyFromBytes([]byte(symmetricKey))

	if err != nil {
		return nil, fmt.Errorf("cannot create symmetric key: %w", err)
	}

	return &PasetoMaker{
		symmetricKey: key,
	}, nil
}

func (maker *PasetoMaker) CreateToken(
	userID uuid.UUID, username, email, role string, duration time.Duration) (string, *TokenPayload, error) {
	payload := &TokenPayload{
		ID:        userID,
		Username:  username,
		Email:     email,
		Role:      role,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}

	token := paseto.NewToken()

	token.SetIssuedAt(payload.IssuedAt)
	token.SetNotBefore(payload.IssuedAt)
	token.SetExpiration(payload.ExpiredAt)
	token.SetString("user_id", payload.ID.String())
	token.SetString("username", payload.Username)
	token.SetString("email", payload.Email)
	token.SetString("role", payload.Role)

	encrypted := token.V4Encrypt(maker.symmetricKey, nil)

	return encrypted, payload, nil
}

func (maker *PasetoMaker) VerifyToken(tokenString string) (*TokenPayload, error) {
	parser := paseto.NewParser()

	parser.AddRule(paseto.NotExpired())

	token, err := parser.ParseV4Local(maker.symmetricKey, tokenString, nil)

	if err != nil {
		return nil, fmt.Errorf("invalid token: %w", err)
	}

	userIDStr, err := token.GetString("user_id")

	if err != nil {
		return nil, fmt.Errorf("error getting user_id from token: %w", err)
	}

	userID, err := uuid.Parse(userIDStr)

	if err != nil {
		return nil, fmt.Errorf("invalid user_id in token: %w", err)
	}

	username, err := token.GetString("username")

	if err != nil {
		return nil, fmt.Errorf("error getting username from token: %w", err)
	}

	email, err := token.GetString("email")

	if err != nil {
		return nil, fmt.Errorf("error getting email from token: %w", err)
	}

	role, err := token.GetString("role")

	if err != nil {
		return nil, fmt.Errorf("error getting role from token: %w", err)
	}

	issuedAt, err := token.GetIssuedAt()

	if err != nil {
		return nil, fmt.Errorf("error getting issuedAt data from token: %w", err)
	}

	expiredAt, err := token.GetExpiration()

	if err != nil {
		return nil, fmt.Errorf("error getting expiration from token: %w", err)
	}

	payload := &TokenPayload{
		ID:        userID,
		Username:  username,
		Email:     email,
		Role:      role,
		IssuedAt:  issuedAt,
		ExpiredAt: expiredAt,
	}

	return payload, nil
}
