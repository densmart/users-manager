package entities

type Resource struct {
	BaseEntity
	Name    string `db:"name"`
	UriMask string `db:"uri_mask"`
}
