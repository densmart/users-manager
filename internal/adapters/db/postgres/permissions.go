package postgres

import (
	"github.com/densmart/users-manager/internal/adapters/db"
	"github.com/densmart/users-manager/internal/adapters/dto"
	"github.com/densmart/users-manager/internal/domain/entities"
	"github.com/densmart/users-manager/internal/logger"
	"github.com/doug-martin/goqu/v9"
	"github.com/jackc/pgx/v4"
	"time"
)

const (
	permissionsTable = "permissions"
)

type PermissionsPostgres struct {
	db *db.WrapperDB
}

func NewPermissionsPostgres(db *db.WrapperDB) *PermissionsPostgres {
	return &PermissionsPostgres{db: db}
}

func (r *PermissionsPostgres) Create(data []dto.CreatePermissionDTO) ([]entities.Permission, error) {
	var permissions []entities.Permission

	conn, err := r.db.Pool.Acquire(r.db.Ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()

	query, err := createPermissionsQuery(data)
	if err != nil {
		return nil, err
	}
	rows, err := conn.Query(r.db.Ctx, query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var permission entities.Permission
		if err = rows.Scan(&permission.Id, &permission.CreatedAt, &permission.RoleID, &permission.ResourceID,
			&permission.MethodMask); err != nil {
			logger.Errorf(err.Error())
			return nil, err
		}
		permissions = append(permissions, permission)
	}
	return permissions, nil
}

func (r *PermissionsPostgres) Update(data []dto.UpdatePermissionDTO) ([]entities.Permission, error) {
	var permissions []entities.Permission
	conn, err := r.db.Pool.Acquire(r.db.Ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()

	// open DB transaction
	tx, err := conn.BeginTx(r.db.Ctx, pgx.TxOptions{})
	if err != nil {
		return nil, err
	}
	for _, dataPerm := range data {
		var permission entities.Permission
		query, err := updatePermissionsQuery(dataPerm)
		if err != nil {
			if rbErr := tx.Rollback(r.db.Ctx); rbErr != nil {
				logger.Errorf("error rollback tx: ", rbErr.Error())
			}
			return nil, err
		}
		row := tx.QueryRow(r.db.Ctx, query)
		if err = row.Scan(&permission.Id, &permission.CreatedAt, &permission.RoleID, &permission.ResourceID,
			&permission.MethodMask); err != nil {
			if rbErr := tx.Rollback(r.db.Ctx); rbErr != nil {
				logger.Errorf("error rollback tx: ", rbErr.Error())
			}
			return nil, err
		}
		permissions = append(permissions, permission)
	}
	if err = tx.Commit(r.db.Ctx); err != nil {
		return nil, err
	}
	return permissions, nil
}

func (r *PermissionsPostgres) Search(data dto.SearchPermissionDTO) ([]entities.Permission, error) {
	var permissions []entities.Permission

	conn, err := r.db.Pool.Acquire(r.db.Ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()

	query, err := searchPermissionsQuery(data)
	if err != nil {
		return nil, err
	}
	rows, err := conn.Query(r.db.Ctx, query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var permission entities.Permission
		if err = rows.Scan(&permission.Id, &permission.CreatedAt, &permission.RoleID, &permission.ResourceID,
			&permission.MethodMask); err != nil {
			return nil, err
		}
		permissions = append(permissions, permission)
	}
	return permissions, nil
}

func (r *PermissionsPostgres) Delete(ids []uint64) error {
	conn, err := r.db.Pool.Acquire(r.db.Ctx)
	if err != nil {
		return err
	}
	defer conn.Release()
	query, err := deletePermissionQuery(ids)
	if err != nil {
		return err
	}
	_, err = conn.Exec(r.db.Ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func createPermissionsQuery(data []dto.CreatePermissionDTO) (string, error) {
	dialect := goqu.Dialect("postgres")
	var data2insert []goqu.Record
	for _, item := range data {
		row := goqu.Record{
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"role_id":     item.RoleID,
			"resource_id": item.ResourceID,
			"method_mask": item.MethodMask,
		}
		data2insert = append(data2insert, row)
	}
	ds := dialect.Insert(permissionsTable).Rows(data2insert).Returning("id", "created_at", "role_id",
		"resource_id", "method_mask")
	return toSQL(ds)
}

func updatePermissionsQuery(data dto.UpdatePermissionDTO) (string, error) {
	dialect := goqu.Dialect("postgres")
	ds := dialect.Update(permissionsTable)
	record := goqu.Record{}
	if data.MethodMask != 0 {
		record["method_mask"] = data.MethodMask
	}
	ds = ds.Set(record).Where(goqu.Ex{"id": data.ID}).Returning("id", "created_at", "role_id",
		"resource_id", "method_mask")

	return toSQL(ds)
}

func searchPermissionsQuery(data dto.SearchPermissionDTO) (string, error) {
	dialect := goqu.Dialect("postgres")
	ds := dialect.Select("id", "created_at", "role_id",
		"resource_id", "method_mask").From(permissionsTable).Where(goqu.Ex{"role_id": data.RoleID})
	if data.ResourceID != nil {
		ds = ds.Where(goqu.Ex{"resource_id": *data.ResourceID})
	}

	ds = filterByCreatedAt(ds, data.CreatedAtFrom, data.CreatedAtTo)
	ds = ordering(ds, data.OrderBy)

	return toSQL(ds)
}

func deletePermissionQuery(ids []uint64) (string, error) {
	dialect := goqu.Dialect("postgres")
	ds := dialect.Delete(permissionsTable).Where(goqu.Ex{"id": ids})

	return toSQL(ds)
}
