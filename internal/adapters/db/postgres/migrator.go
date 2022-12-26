package postgres

import (
	"context"
	"github.com/densmart/users-manager/internal/adapters/db"
	"github.com/jackc/tern/migrate"
)

type MigratorPostgres struct {
	db *db.WrapperDB
}

func NewMigratorPostgres(db *db.WrapperDB) *MigratorPostgres {
	return &MigratorPostgres{db: db}
}

func (r *MigratorPostgres) Up() error {
	conn, err := r.db.Pool.Acquire(r.db.Ctx)
	defer conn.Release()
	if err != nil {
		return err
	}
	migrator, err := migrate.NewMigrator(context.Background(), conn.Conn(), "schema_version")
	if err != nil {
		return err
	}
	if err = migrator.LoadMigrations("./migrations"); err != nil {
		return err
	}
	if err = migrator.Migrate(context.Background()); err != nil {
		return err
	}
	return nil
}

func (r *MigratorPostgres) Down() error {
	return nil
}
