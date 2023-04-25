package dto

type CreateResourceDTO struct {
	Name        string  `json:"name"`
	UriMask     string  `json:"uri_mask"`
	ActionsMask string  `json:"actions_mask"`
	IsActive    bool    `json:"is_active"`
	ResGroup    *string `json:"res_group,omitempty"`
	MethodMask  uint8
}

type UpdateResourceDTO struct {
	ID          uint64
	Name        *string `json:"name,omitempty"`
	UriMask     *string `json:"uri_mask,omitempty"`
	ActionsMask *string `json:"actions_mask"`
	IsActive    *bool   `json:"is_active,omitempty"`
	ResGroup    *string `json:"res_group,omitempty"`
	MethodMask  *uint8
}

type SearchResourceDTO struct {
	BaseSearchRequestDto
	ID       *uint64 `form:"id"`
	Name     *string `form:"name"`
	UriMask  *string `form:"uri_mask"`
	IsActive *bool   `form:"is_active"`
}

type ResourceDTO struct {
	ID          uint64 `json:"id"`
	CreatedAt   string `json:"created_at"`
	Name        string `json:"name"`
	UriMask     string `json:"uri_mask"`
	ActionsMask string `json:"actions_mask"`
	IsActive    bool   `json:"is_active"`
}

type ResourceGroupDTO struct {
	Name      string        `json:"name"`
	Resources []ResourceDTO `json:"resources"`
}

type ResourcesDTO struct {
	Items []ResourceGroupDTO
}
