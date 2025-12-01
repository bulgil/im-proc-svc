package user

type Repository interface {
	Get(id int64) (*User, error)
	GetByUsername(username string) (*User, error)
	Create(user *User) error
	Update(user *User) error
	Delete(user *User) error
	CheckUsername(username string) bool
}
