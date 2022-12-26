package repo

import (
	"github.com/densmart/users-manager/internal/adapters/db"
	"github.com/densmart/users-manager/internal/adapters/db/postgres"
)

type Migrator interface {
	Up() error
	Down() error
}

type Repo struct {
	Migrator
}

func NewRepo(db *db.WrapperDB) *Repo {
	return &Repo{
		Migrator: postgres.NewMigratorPostgres(db),
	}
}
