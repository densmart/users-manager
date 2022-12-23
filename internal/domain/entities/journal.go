package entities

type Journal struct {
	BaseEntity
	UserID       uint64 `db:"user_id"`
	ActionID     uint64 `db:"action_id"`
	ResourceID   uint64 `db:"resource_id"`
	RequestData  string `db:"request_data"`
	ResponseData string `db:"response_data"`
}
