package dto

type CreateRoleDTO struct {
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	IsPermitted bool   `json:"is_permitted"`
}

type UpdateRoleDTO struct {
	ID          uint64
	Name        *string `json:"name,omitempty"`
	Slug        *string `json:"slug,omitempty"`
	IsPermitted *bool   `json:"is_permitted,omitempty"`
}

type SearchRoleDTO struct {
	BaseSearchRequestDto
	ID          *uint64 `form:"id"`
	Name        *string `form:"name"`
	Slug        *string `form:"slug"`
	IsPermitted *bool   `form:"is_permitted"`
}

type RoleDTO struct {
	ID          uint64 `json:"id"`
	CreatedAt   string `json:"created_at"`
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	IsPermitted bool   `json:"is_permitted"`
}

type RolesDTO struct {
	Pagination PaginationInfo `json:"pagination"`
	Items      []RoleDTO      `json:"items"`
}
