package model

type Role string

const (
	None  Role = "none"
	Admin Role = "admin"
)

type User struct {
	ID    uint
	Name  string
	Role  Role
	Admin bool
}

func (u *User) IsAdmin() bool {
	return u.Admin
}
