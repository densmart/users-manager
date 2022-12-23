package entities

const (
	SuperUserRoleSlug = "superuser"
)

type Role struct {
	BaseEntity
	Name        string `db:"name"`
	Slug        string `db:"slug"`
	IsPermitted bool   `db:"is_permitted"`
}
