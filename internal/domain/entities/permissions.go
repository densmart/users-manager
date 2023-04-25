package entities

type Permission struct {
	BaseEntity
	RoleID     uint64 `db:"role_id"`
	ResourceID uint64 `db:"resource_id"`
}
