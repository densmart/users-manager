package entities

type Permission struct {
	BaseEntity
	RoleID             uint64 `db:"role_id"`
	ResourceID         uint64 `db:"resource_id"`
	ResourceURIMask    string `db:"resource_uri_mask"`
	ResourceMethodMask uint8  `db:"resource_method_mask"`
	MethodMask         uint8  `db:"method_mask"`
}
