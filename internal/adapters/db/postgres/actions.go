package postgres

import (
	"github.com/densmart/users-manager/internal/adapters/db"
	"github.com/densmart/users-manager/internal/adapters/dto"
	"github.com/densmart/users-manager/internal/domain/entities"
	"github.com/doug-martin/goqu/v9"
	"time"
)

const (
	actionsTable = "actions"
)

type ActionsPostgres struct {
	db *db.WrapperDB
}

func NewActionsPostgres(db *db.WrapperDB) *ActionsPostgres {
	return &ActionsPostgres{db: db}
}

func (r *ActionsPostgres) Create(data dto.CreateActionDTO) (entities.Action, error) {
	var action entities.Action

	conn, err := r.db.Pool.Acquire(r.db.Ctx)
	if err != nil {
		return action, err
	}
	defer conn.Release()

	query, err := createActionQuery(data)
	if err != nil {
		return action, err
	}
	row := conn.QueryRow(r.db.Ctx, query)
	if err = row.Scan(&action.Id, &action.CreatedAt, &action.Name, &action.Method); err != nil {
		return action, err
	}
	return action, nil
}

func (r *ActionsPostgres) Update(data dto.UpdateActionDTO) (entities.Action, error) {
	var action entities.Action
	conn, err := r.db.Pool.Acquire(r.db.Ctx)
	if err != nil {
		return action, err
	}
	defer conn.Release()

	query, err := updateActionQuery(data)
	if err != nil {
		return action, err
	}
	row := conn.QueryRow(r.db.Ctx, query)
	if err = row.Scan(&action.Id, &action.CreatedAt, &action.Name, &action.Method); err != nil {
		return action, err
	}
	return action, nil
}

func (r *ActionsPostgres) Retrieve(id uint64) (entities.Action, error) {
	var action entities.Action
	conn, err := r.db.Pool.Acquire(r.db.Ctx)
	if err != nil {
		return action, err
	}
	defer conn.Release()

	query, err := retrieveActionQuery(id)
	if err != nil {
		return action, err
	}
	row := conn.QueryRow(r.db.Ctx, query)
	if err = row.Scan(&action.Id, &action.CreatedAt, &action.UpdatedAt, &action.Name, &action.Method); err != nil {
		return action, err
	}
	return action, nil
}

func (r *ActionsPostgres) Search(data dto.SearchActionDTO) ([]entities.Action, error) {
	var actions []entities.Action
	conn, err := r.db.Pool.Acquire(r.db.Ctx)
	if err != nil {
		return actions, err
	}
	defer conn.Release()

	query, err := searchActionQuery(data)
	if err != nil {
		return actions, err
	}
	rows, err := conn.Query(r.db.Ctx, query)
	if err != nil {
		return actions, err
	}
	for rows.Next() {
		var action entities.Action
		if err = rows.Scan(&action.Id, &action.CreatedAt, &action.UpdatedAt, &action.Name, &action.Method); err != nil {
			return actions, err
		}
		actions = append(actions, action)
	}

	return actions, nil
}

func (r *ActionsPostgres) Delete(id uint64) error {
	conn, err := r.db.Pool.Acquire(r.db.Ctx)
	if err != nil {
		return err
	}
	defer conn.Release()
	query, err := deleteActionQuery(id)
	if err != nil {
		return err
	}
	_, err = conn.Exec(r.db.Ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func createActionQuery(data dto.CreateActionDTO) (string, error) {
	dialect := goqu.Dialect("postgres")
	data2insert := goqu.Record{
		"created_at": time.Now(),
		"updated_at": time.Now(),
		"name":       data.Name,
		"method":     data.Method,
	}
	ds := dialect.Insert(actionsTable).Rows(data2insert).Returning("id", "created_at", "name", "method")

	return toSQL(ds)
}

func updateActionQuery(data dto.UpdateActionDTO) (string, error) {
	dialect := goqu.Dialect("postgres")
	ds := dialect.Update(actionsTable)
	record := goqu.Record{}
	if data.Name != nil {
		record["name"] = data.Name
	}
	if data.Method != nil {
		record["method"] = data.Method
	}
	ds = ds.Set(record).Where(goqu.Ex{"id": data.ID}).Returning("id", "created_at", "name", "method")

	return toSQL(ds)
}

func retrieveActionQuery(id uint64) (string, error) {
	dialect := goqu.Dialect("postgres")
	ds := dialect.Select("id", "created_at", "updated_at", "name", "method").From(
		actionsTable).Where(goqu.Ex{"id": id})

	return toSQL(ds)
}

func searchActionQuery(data dto.SearchActionDTO) (string, error) {
	dialect := goqu.Dialect("postgres")
	ds := dialect.Select("id", "created_at", "updated_at", "name", "method").From(actionsTable)

	if data.Name != nil {
		ds = ds.Where(goqu.Ex{
			"name": goqu.Op{"like": *data.Name},
		})
	}
	if data.Method != nil {
		ds = ds.Where(goqu.Ex{"method": *data.Method})
	}
	if data.ID != nil {
		ds = ds.Where(goqu.Ex{"id": *data.ID})
	}

	ds = filterByCreatedAt(ds, data.CreatedAtFrom, data.CreatedAtTo)
	ds = ordering(ds, data.OrderBy)
	ds = slicing(ds, data.Offset, data.Limit)

	return toSQL(ds)
}

func deleteActionQuery(id uint64) (string, error) {
	dialect := goqu.Dialect("postgres")
	ds := dialect.Delete(actionsTable).Where(goqu.Ex{"id": id})

	return toSQL(ds)
}
