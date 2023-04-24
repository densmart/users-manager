package postgres

import (
	"github.com/densmart/users-manager/internal/adapters/db"
	"github.com/densmart/users-manager/internal/adapters/dto"
	"github.com/densmart/users-manager/internal/domain/entities"
	"github.com/doug-martin/goqu/v9"
	"time"
)

const (
	resourcesTable = "resources"
)

type ResourcesPostgres struct {
	db *db.WrapperDB
}

func NewResourcesPostgres(db *db.WrapperDB) *ResourcesPostgres {
	return &ResourcesPostgres{db: db}
}

func (r *ResourcesPostgres) Create(data dto.CreateResourceDTO) (entities.Resource, error) {
	var resource entities.Resource

	conn, err := r.db.Pool.Acquire(r.db.Ctx)
	if err != nil {
		return resource, err
	}
	defer conn.Release()

	query, err := createResourceQuery(data)
	if err != nil {
		return resource, err
	}
	row := conn.QueryRow(r.db.Ctx, query)
	if err = row.Scan(&resource.Id, &resource.CreatedAt, &resource.Name, &resource.UriMask); err != nil {
		return resource, err
	}
	return resource, nil
}

func (r *ResourcesPostgres) Update(data dto.UpdateResourceDTO) (entities.Resource, error) {
	var resource entities.Resource
	conn, err := r.db.Pool.Acquire(r.db.Ctx)
	if err != nil {
		return resource, err
	}
	defer conn.Release()

	query, err := updateResourceQuery(data)
	if err != nil {
		return resource, err
	}
	row := conn.QueryRow(r.db.Ctx, query)
	if err = row.Scan(&resource.Id, &resource.CreatedAt, &resource.Name, &resource.UriMask); err != nil {
		return resource, err
	}
	return resource, nil
}

func (r *ResourcesPostgres) Retrieve(id uint64) (entities.Resource, error) {
	var resource entities.Resource
	conn, err := r.db.Pool.Acquire(r.db.Ctx)
	if err != nil {
		return resource, err
	}
	defer conn.Release()

	query, err := retrieveResourceQuery(id)
	if err != nil {
		return resource, err
	}
	row := conn.QueryRow(r.db.Ctx, query)
	if err = row.Scan(&resource.Id, &resource.CreatedAt, &resource.UpdatedAt, &resource.Name, &resource.UriMask); err != nil {
		return resource, err
	}
	return resource, nil
}

func (r *ResourcesPostgres) Search(data dto.SearchResourceDTO) ([]entities.Resource, error) {
	var resources []entities.Resource
	conn, err := r.db.Pool.Acquire(r.db.Ctx)
	if err != nil {
		return resources, err
	}
	defer conn.Release()

	query, err := searchResourceQuery(data)
	if err != nil {
		return resources, err
	}
	rows, err := conn.Query(r.db.Ctx, query)
	if err != nil {
		return resources, err
	}
	for rows.Next() {
		var resource entities.Resource
		if err = rows.Scan(&resource.Id, &resource.CreatedAt, &resource.UpdatedAt, &resource.Name, &resource.UriMask); err != nil {
			return resources, err
		}
		resources = append(resources, resource)
	}

	return resources, nil
}

func (r *ResourcesPostgres) Delete(id uint64) error {
	conn, err := r.db.Pool.Acquire(r.db.Ctx)
	if err != nil {
		return err
	}
	defer conn.Release()
	query, err := deleteResourceQuery(id)
	if err != nil {
		return err
	}
	_, err = conn.Exec(r.db.Ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func createResourceQuery(data dto.CreateResourceDTO) (string, error) {
	dialect := goqu.Dialect("postgres")
	data2insert := goqu.Record{
		"created_at": time.Now(),
		"updated_at": time.Now(),
		"name":       data.Name,
		"uri_mask":   data.UriMask,
	}
	ds := dialect.Insert(resourcesTable).Rows(data2insert).Returning("id", "created_at", "name", "uri_mask")

	return toSQL(ds)
}

func updateResourceQuery(data dto.UpdateResourceDTO) (string, error) {
	dialect := goqu.Dialect("postgres")
	ds := dialect.Update(resourcesTable)
	record := goqu.Record{}
	if data.Name != nil {
		record["name"] = data.Name
	}
	if data.UriMask != nil {
		record["uri_mask"] = data.UriMask
	}
	ds = ds.Set(record).Where(goqu.Ex{"id": data.ID}).Returning("id", "created_at", "name", "uri_mask")

	return toSQL(ds)
}

func retrieveResourceQuery(id uint64) (string, error) {
	dialect := goqu.Dialect("postgres")
	ds := dialect.Select("id", "created_at", "updated_at", "name", "uri_mask").From(
		resourcesTable).Where(goqu.Ex{"id": id})

	return toSQL(ds)
}

func searchResourceQuery(data dto.SearchResourceDTO) (string, error) {
	dialect := goqu.Dialect("postgres")
	ds := dialect.Select("id", "created_at", "updated_at", "name", "uri_mask").From(resourcesTable)

	if data.Name != nil {
		ds = ds.Where(goqu.Ex{
			"name": goqu.Op{"like": *data.Name},
		})
	}
	if data.UriMask != nil {
		ds = ds.Where(goqu.Ex{
			"name": goqu.Op{"uri_mask": *data.UriMask},
		})
	}
	if data.ID != nil {
		ds = ds.Where(goqu.Ex{"id": *data.ID})
	}

	ds = filterByCreatedAt(ds, data.CreatedAtFrom, data.CreatedAtTo)
	ds = ordering(ds, data.OrderBy)
	ds = slicing(ds, data.Offset, data.Limit)

	return toSQL(ds)
}

func deleteResourceQuery(id uint64) (string, error) {
	dialect := goqu.Dialect("postgres")
	ds := dialect.Delete(resourcesTable).Where(goqu.Ex{"id": id})

	return toSQL(ds)
}
