
package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/aksharma27/go-otp-auth/internal/models"
)

type UserRepository struct {
	DB *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (r *UserRepository) Create(
	ctx context.Context,
	phoneNumber string,
) (*models.User, error) {

	query := `
		INSERT INTO users (phone_number, is_verified)
		VALUES ($1, TRUE)
		RETURNING id, phone_number, is_verified, created_at, updated_at
	`

	var user models.User

	err := r.DB.QueryRow(
		ctx,
		query,
		phoneNumber,
	).Scan(
		&user.ID,
		&user.PhoneNumber,
		&user.IsVerified,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("create user: %w", err)
	}

	return &user, nil
}

func (r *UserRepository) FindByPhone(
	ctx context.Context,
	phoneNumber string,
) (*models.User, error) {

	query := `
		SELECT id, phone_number, is_verified, created_at, updated_at
		FROM users
		WHERE phone_number = $1
	`

	var user models.User

	err := r.DB.QueryRow(
		ctx,
		query,
		phoneNumber,
	).Scan(
		&user.ID,
		&user.PhoneNumber,
		&user.IsVerified,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}

	if err != nil {
		return nil, fmt.Errorf("find user by phone: %w", err)
	}

	return &user, nil
}

func (r *UserRepository) FindByID(
	ctx context.Context,
	id uuid.UUID,
) (*models.User, error) {

	query := `
		SELECT id, phone_number, is_verified, created_at, updated_at
		FROM users
		WHERE id = $1
	`

	var user models.User

	err := r.DB.QueryRow(
		ctx,
		query,
		id,
	).Scan(
		&user.ID,
		&user.PhoneNumber,
		&user.IsVerified,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}

	if err != nil {
		return nil, fmt.Errorf("find user by ID: %w", err)
	}

	return &user, nil
}
