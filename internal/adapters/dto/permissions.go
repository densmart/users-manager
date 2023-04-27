package dto

type CreatePermissionDTO struct {
	RoleID      uint64
	ResourceID  uint64 `json:"resource_id"`
	ActionsMask string `json:"action_mask"`
	MethodMask  uint8
}

type UpdatePermissionDTO struct {
	ID          uint64  `json:"id"`
	ActionsMask *string `json:"action_mask"`
	MethodMask  uint8
}

type SearchPermissionDTO struct {
	BaseSearchRequestDto
	RoleID     uint64
	ResourceID *uint64 `form:"resource_id"`
}

type PermissionDTO struct {
	ID          uint64 `json:"id"`
	CreatedAt   string `json:"created_at"`
	RoleID      uint64 `json:"role_id"`
	ResourceID  uint64 `json:"resource_id"`
	ActionsMask string `json:"actions_mask"`
}

type PermissionsDTO struct {
	Pagination string `json:"-"`
	Items      []PermissionDTO
}
