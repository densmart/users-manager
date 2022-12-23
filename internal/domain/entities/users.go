package entities

import "time"

type User struct {
	BaseEntity
	Email       string    `db:"email"`
	Password    string    `db:"password"`
	FirstName   string    `db:"first_name"`
	LastName    string    `db:"last_name"`
	Phone       string    `db:"phone"`
	IsActive    bool      `db:"is_active"`
	Is2fa       bool      `db:"is_2fa"`
	Token2fa    string    `db:"token_2fa"`
	LastLoginAt time.Time `db:"last_login_at"`
	RoleID      uint64    `db:"role_id"`
}
