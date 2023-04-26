package dto

type CreatePermissionDTO struct {
	RoleID     uint64 `json:"role_id"`
	ResourceID uint64 `json:"resource_id"`
	ActionMask string `json:"action_mask"`
	MethodMask uint8
}

type UpdatePermissionDTO struct {
	ID         uint64
	ActionMask *string `json:"action_mask"`
	MethodMask *uint8
}

type SearchPermissionDTO struct {
	BaseSearchRequestDto
	RoleID     uint64
	ResourceID *uint64 `json:"resource_id,omitempty"`
}

type PermissionDTO struct {
	Name        string  `json:"name"`
	UriMask     string  `json:"uri_mask"`
	ID          *uint64 `json:"id,omitempty"`
	ActionsMask *string `json:"actions_mask,omitempty"`
}

type PermissionGroupDTO struct {
	Name        string          `json:"name"`
	Permissions []PermissionDTO `json:"permissions"`
}

type PermissionsDTO struct {
	Items []PermissionGroupDTO
}
