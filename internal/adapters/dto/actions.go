package dto

type CreateActionDTO struct {
	Name   string `json:"name"`
	Method string `json:"method"`
}

type UpdateActionDTO struct {
	ID     uint64
	Name   *string `json:"name,omitempty"`
	Method *string `json:"method,omitempty"`
}

type SearchActionDTO struct {
	BaseSearchRequestDto
	ID     *uint64 `form:"id"`
	Name   *string `form:"name"`
	Method *string `form:"method"`
}

type ActionDTO struct {
	ID        uint64 `json:"id"`
	CreatedAt string `json:"created_at"`
	Name      string `json:"name"`
	Method    string `json:"method"`
}

type ActionsDTO struct {
	Pagination string      `json:"-"`
	Items      []ActionDTO `json:"items"`
}
