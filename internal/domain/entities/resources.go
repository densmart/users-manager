package entities

import "database/sql"

type Resource struct {
	BaseEntity
	Name       string         `db:"name"`
	UriMask    string         `db:"uri_mask"`
	MethodMask int8           `db:"method_mask"`
	IsActive   bool           `db:"is_active"`
	ResGroup   sql.NullString `db:"res_group"` // Resource grouping for correct list representation
}
