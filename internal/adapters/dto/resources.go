package dto

type CreateResourceDTO struct {
	Name    string `json:"name"`
	UriMask string `json:"uri_mask"`
}

type UpdateResourceDTO struct {
	ID      uint64
	Name    *string `json:"name,omitempty"`
	UriMask *string `json:"uri_mask,omitempty"`
}

type SearchResourceDTO struct {
	BaseSearchRequestDto
	ID      *uint64 `form:"id"`
	Name    *string `form:"name"`
	UriMask *string `form:"uri_mask"`
}

type ResourceDTO struct {
	ID        uint64 `json:"id"`
	CreatedAt string `json:"created_at"`
	Name      string `json:"name"`
	UriMask   string `json:"uri_mask"`
}

type ResourcesDTO struct {
	Pagination string        `json:"-"`
	Items      []ResourceDTO `json:"items"`
}
