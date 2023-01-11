package postgres

import (
	"errors"
	"github.com/densmart/users-manager/internal/adapters/db"
	"github.com/densmart/users-manager/internal/adapters/dto"
	"github.com/densmart/users-manager/internal/domain/entities"
	"github.com/doug-martin/goqu/v9"
	"github.com/jackc/pgconn"
	"net/http"
	"time"
)

const (
	rolesTable = "roles"
)

type RolesPostgres struct {
	db *db.WrapperDB
}

func NewRolesPostgres(db *db.WrapperDB) *RolesPostgres {
	return &RolesPostgres{db: db}
}

func (r *RolesPostgres) Create(data dto.CreateRoleDTO) (entities.Role, error) {
	var role entities.Role
	var pgErr *pgconn.PgError

	conn, err := r.db.Pool.Acquire(r.db.Ctx)
	if err != nil {
		return role, err
	}
	defer conn.Release()

	query, err := createRoleQuery(data)
	if err != nil {
		return role, err
	}
	row := conn.QueryRow(r.db.Ctx, query)
	if err = row.Scan(&role.Id, &role.CreatedAt, &role.Name, &role.Slug, &role.IsPermitted); err != nil {
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return role, &dto.APIError{
				HttpCode: http.StatusConflict,
				Message:  "role slug already exists",
			}
		}
		return role, err
	}

	return role, nil
}

func (r *RolesPostgres) Update(data dto.UpdateRoleDTO) (entities.Role, error) {
	var role entities.Role
	conn, err := r.db.Pool.Acquire(r.db.Ctx)
	if err != nil {
		return role, err
	}
	defer conn.Release()

	query, err := updateRoleQuery(data)
	if err != nil {
		return role, err
	}
	row := conn.QueryRow(r.db.Ctx, query)
	if err = row.Scan(&role.Id, &role.CreatedAt, &role.Name, &role.Slug, &role.IsPermitted); err != nil {
		return role, err
	}
	return role, nil
}

func (r *RolesPostgres) Retrieve(id uint64) (entities.Role, error) {
	var role entities.Role
	conn, err := r.db.Pool.Acquire(r.db.Ctx)
	if err != nil {
		return role, err
	}
	defer conn.Release()

	query, err := retrieveRoleQuery(id)
	if err != nil {
		return role, err
	}
	row := conn.QueryRow(r.db.Ctx, query)
	if err = row.Scan(&role.Id, &role.CreatedAt, &role.Name, &role.Slug, &role.IsPermitted); err != nil {
		return role, err
	}
	return role, nil
}

func (r *RolesPostgres) Search(data dto.SearchRoleDTO) ([]entities.Role, error) {
	var roles []entities.Role
	conn, err := r.db.Pool.Acquire(r.db.Ctx)
	if err != nil {
		return roles, err
	}
	defer conn.Release()

	query, err := searchRoleQuery(data)
	if err != nil {
		return roles, err
	}
	rows, err := conn.Query(r.db.Ctx, query)
	if err != nil {
		return roles, err
	}
	for rows.Next() {
		var role entities.Role
		if err = rows.Scan(&role.Id, &role.CreatedAt, &role.Name, &role.Slug, &role.IsPermitted); err != nil {
			return roles, err
		}
		roles = append(roles, role)
	}

	return roles, nil
}

func (r *RolesPostgres) Delete(id uint64) error {
	conn, err := r.db.Pool.Acquire(r.db.Ctx)
	if err != nil {
		return err
	}
	defer conn.Release()
	query, err := deleteRoleQuery(id)
	if err != nil {
		return err
	}
	_, err = conn.Exec(r.db.Ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func createRoleQuery(data dto.CreateRoleDTO) (string, error) {
	dialect := goqu.Dialect("postgres")
	data2insert := goqu.Record{
		"created_at":   time.Now(),
		"updated_at":   time.Now(),
		"name":         data.Name,
		"slug":         data.Slug,
		"is_permitted": data.IsPermitted,
	}
	ds := dialect.Insert(rolesTable).Rows(data2insert).Returning("id", "created_at", "name", "slug",
		"is_permitted")

	return toSQL(ds)
}

func updateRoleQuery(data dto.UpdateRoleDTO) (string, error) {
	dialect := goqu.Dialect("postgres")
	ds := dialect.Update(rolesTable)
	record := goqu.Record{}
	if data.Name != nil {
		record["name"] = data.Name
	}
	if data.Slug != nil {
		record["slug"] = data.Slug
	}
	if data.IsPermitted != nil {
		record["is_permitted"] = data.IsPermitted
	}
	ds = ds.Set(record).Where(goqu.Ex{"id": data.ID}).Returning("id", "created_at", "name", "slug",
		"is_permitted")

	return toSQL(ds)
}

func retrieveRoleQuery(id uint64) (string, error) {
	dialect := goqu.Dialect("postgres")
	ds := dialect.Select("id", "created_at", "updated_at", "name", "slug",
		"is_permitted").From(rolesTable).Where(goqu.Ex{"id": id})

	return toSQL(ds)
}

func searchRoleQuery(data dto.SearchRoleDTO) (string, error) {
	dialect := goqu.Dialect("postgres")
	ds := dialect.Select("id", "created_at", "updated_at", "name", "slug",
		"is_permitted").From(rolesTable)

	if data.Name != nil {
		ds = ds.Where(goqu.Ex{
			"name": goqu.Op{"like": *data.Name},
		})
	}
	if data.Slug != nil {
		ds = ds.Where(goqu.Ex{"slug": *data.Slug})
	}
	if data.ID != nil {
		ds = ds.Where(goqu.Ex{"id": *data.ID})
	}

	ds = filterByCreatedAt(ds, data.CreatedAtFrom, data.CreatedAtTo)
	ds = ordering(ds, data.OrderBy)
	ds = slicing(ds, data.Offset, data.Limit)

	return toSQL(ds)
}

func deleteRoleQuery(id uint64) (string, error) {
	dialect := goqu.Dialect("postgres")
	ds := dialect.Delete(rolesTable).Where(goqu.Ex{"id": id})

	return toSQL(ds)
}
