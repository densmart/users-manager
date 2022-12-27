package dto

import "time"

type BaseSearchRequestDto struct {
	RawURL        string
	CreatedAtFrom time.Time `form:"created_at_from"`
	CreatedAtTo   time.Time `form:"created_at_to"`
	Page          *uint     `form:"page"`
	PerPage       *uint     `form:"per_page"`
	OrderBy       *string   `form:"order_by"`
	Offset        *uint
	Limit         *uint
}

type PaginationInfo struct {
	Total   uint   `json:"total"`
	Page    uint   `json:"page"`
	PerPage uint   `json:"per_page"`
	Pages   uint   `json:"pages"`
	Next    string `json:"next"`
	Prev    string `json:"prev"`
}
