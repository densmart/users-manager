package entities

import "time"

type BaseEntity struct {
	Id        uint64    `db:"id"`         // table general primary key bigint, autoincrement
	CreatedAt time.Time `db:"created_at"` // record create datetime
	UpdatedAt time.Time `db:"updated_at"` // record update datetime
}
