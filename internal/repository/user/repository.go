package user

import (
	"context"
	"errors"
	"fmt"

	domain "github.com/bulgil/im-proc-svc/internal/domain/user"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	db *pgxpool.Pool
}

func New(db *pgxpool.Pool) *Repository {
	return &Repository{db}
}

func (r *Repository) Get(ctx context.Context, id int64) (*domain.User, error) {
	const op = "repository.user.Repository.Get"

	stmt := `
SELECT id, username, passhash, created_at 
FROM users
WHERE id=$1;`

	var user domain.User
	err := r.db.QueryRow(ctx, stmt, id).Scan(&user)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, domain.ErrNoUser
		}

		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &user, nil
}

func (r *Repository) GetByUsername(ctx context.Context, username string) (*domain.User, error) {
	const op = "repository.user.Repository.Get"

	stmt := `
SELECT id, username, passhash, created_at 
FROM users
WHERE username=$1;`

	var user domain.User
	err := r.db.QueryRow(ctx, stmt, username).Scan(&user.ID, &user.Username, &user.Passhash, &user.CreatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, domain.ErrNoUser
		}

		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &user, nil
}

func (r *Repository) Create(ctx context.Context, user *domain.User) error {
	const op = "repository.user.Repository.Create"

	stmt := `
INSERT INTO users
(username, passhash)
VALUES ($1, $2)
RETURNING id, created_at;`

	// TODO constraint error handling
	err := r.db.QueryRow(ctx, stmt, user.Username, user.Passhash).Scan(&user.ID, &user.CreatedAt)
	if err != nil {
		var pgError *pgconn.PgError
		if errors.As(err, &pgError); pgError.Code == pgerrcode. {
			
		}

		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (r *Repository) Update(ctx context.Context, user *domain.User) error {
	panic("todo")
}

func (r *Repository) Delete(ctx context.Context, user *domain.User) error {
	panic("todo")
}

func (r *Repository) CheckUsername(ctx context.Context, username string) (bool, error) {
	const op = "repository.user.Repository.CheckUsername"

	stmt := `
SELECT id FROM users
WHERE username = $1;`

	var id int64
	err := r.db.QueryRow(ctx, stmt, username).Scan(&id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return false, nil
		}

		return false, fmt.Errorf("%s: %w", op, err)
	}

	return true, nil
}
