package entities

const (
	MethodGet    = "GET"
	MethodPost   = "POST"
	MethodPut    = "PUT"
	MethodPatch  = "PATCH"
	MethodDelete = "DELETE"
)

type Action struct {
	BaseEntity
	Name   string `db:"name"`
	Method string `db:"method"`
}
