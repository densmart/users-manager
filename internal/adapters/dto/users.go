package dto

type CreateUserDTO struct {
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Phone     string `json:"phone,omitempty"`
	IsActive  bool   `json:"is_active"`
	Is2fa     bool   `json:"is_2fa"`
	RoleID    uint64 `json:"role_id"`
	Password  string `json:"password"`
	Token2fa  string
}

type UpdateUserDTO struct {
	ID        uint64
	Email     *string `json:"email,omitempty"`
	FirstName *string `json:"first_name,omitempty"`
	LastName  *string `json:"last_name,omitempty"`
	Phone     *string `json:"phone,omitempty"`
	IsActive  *bool   `json:"is_active,omitempty"`
	Is2fa     *bool   `json:"is_2fa,omitempty"`
	RoleID    *uint64 `json:"role_id,omitempty"`
	Password  *string
}

type SearchUserDTO struct {
	BaseSearchRequestDto
	ID        *uint64 `form:"id"`
	Email     *string `form:"email"`
	FirstName *string `form:"first_name"`
	LastName  *string `form:"last_name"`
	Phone     *string `form:"phone"`
	IsActive  *bool   `form:"is_active"`
	Is2fa     *bool   `form:"is_2fa"`
	RoleID    *uint64 `form:"role_id"`
}

type UserDTO struct {
	ID        uint64 `json:"id"`
	CreatedAt string `json:"created_at"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Phone     string `json:"phone"`
	IsActive  bool   `json:"is_active"`
	Is2fa     bool   `json:"is_2fa"`
	RoleID    uint64 `json:"role_id"`
}

type UsersDTO struct {
	Pagination string    `json:"-"`
	Items      []UserDTO `json:"items"`
}
