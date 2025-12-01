package user

import "context"

type Repository interface {
	Get(ctx context.Context, id int64) (*User, error)
	GetByUsername(ctx context.Context, username string) (*User, error)
	Create(ctx context.Context, user *User) error
	Update(ctx context.Context, user *User) error
	Delete(ctx context.Context, user *User) error
	CheckUsername(ctx context.Context, username string) (bool, error)
}
