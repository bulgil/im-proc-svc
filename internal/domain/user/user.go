package user

import (
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Password []byte `json:"-"`
	Passhash []byte `json:"-"`

	CreatedAt time.Time `json:"created_at"`
}

func (u *User) HashPassword() error {
	op := "domain.user.User.HashPassword"

	hash, err := bcrypt.GenerateFromPassword(u.Password, bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	u.Passhash = hash
	return nil
}

func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword(u.Passhash, []byte(password))
	return err == nil
}
