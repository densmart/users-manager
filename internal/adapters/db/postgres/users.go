package postgres

import (
	"github.com/densmart/users-manager/internal/adapters/db"
	"github.com/densmart/users-manager/internal/adapters/dto"
	"github.com/densmart/users-manager/internal/domain/entities"
	"github.com/doug-martin/goqu/v9"
	"time"
)

const (
	usersTable = "users"
)

type UsersPostgres struct {
	db *db.WrapperDB
}

func NewUsersPostgres(db *db.WrapperDB) *UsersPostgres {
	return &UsersPostgres{db: db}
}

func (r *UsersPostgres) Create(data dto.CreateUserDTO) (entities.User, error) {
	var user entities.User

	conn, err := r.db.Pool.Acquire(r.db.Ctx)
	if err != nil {
		return user, err
	}
	defer conn.Release()

	query, err := createUserQuery(data)
	if err != nil {
		return user, err
	}
	row := conn.QueryRow(r.db.Ctx, query)
	if err = row.Scan(&user.Id, &user.CreatedAt, &user.UpdatedAt, &user.Email, &user.FirstName, &user.LastName, &user.Phone,
		&user.IsActive, &user.Is2fa, &user.RoleID, &user.Password, &user.Token2fa); err != nil {
		return user, err
	}
	return user, nil
}

func (r *UsersPostgres) Update(data dto.UpdateUserDTO) (entities.User, error) {
	var user entities.User
	conn, err := r.db.Pool.Acquire(r.db.Ctx)
	if err != nil {
		return user, err
	}
	defer conn.Release()

	query, err := updateUserQuery(data)
	if err != nil {
		return user, err
	}
	row := conn.QueryRow(r.db.Ctx, query)
	if err = row.Scan(&user.Id, &user.CreatedAt, &user.UpdatedAt, &user.Email, &user.FirstName, &user.LastName, &user.Phone,
		&user.IsActive, &user.Is2fa, &user.RoleID, &user.Password, &user.Token2fa); err != nil {
		return user, err
	}
	return user, nil
}

func (r *UsersPostgres) Retrieve(id uint64) (entities.User, error) {
	var user entities.User
	conn, err := r.db.Pool.Acquire(r.db.Ctx)
	if err != nil {
		return user, err
	}
	defer conn.Release()

	query, err := retrieveUserQuery(id)
	if err != nil {
		return user, err
	}
	row := conn.QueryRow(r.db.Ctx, query)
	if err = row.Scan(&user.Id, &user.CreatedAt, &user.UpdatedAt, &user.Email, &user.FirstName, &user.LastName, &user.Phone,
		&user.IsActive, &user.Is2fa, &user.RoleID, &user.Password, &user.Token2fa); err != nil {
		return user, err
	}
	return user, nil
}

func (r *UsersPostgres) Search(data dto.SearchUserDTO) ([]entities.User, error) {
	var users []entities.User
	conn, err := r.db.Pool.Acquire(r.db.Ctx)
	if err != nil {
		return users, err
	}
	defer conn.Release()

	query, err := searchUserQuery(data)
	if err != nil {
		return users, err
	}
	rows, err := conn.Query(r.db.Ctx, query)
	if err != nil {
		return users, err
	}
	for rows.Next() {
		var user entities.User
		if err = rows.Scan(&user.Id, &user.CreatedAt, &user.UpdatedAt, &user.Email, &user.FirstName, &user.LastName,
			&user.Phone, &user.IsActive, &user.Is2fa, &user.RoleID, &user.Password, &user.Token2fa); err != nil {
			return users, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (r *UsersPostgres) Delete(id uint64) error {
	conn, err := r.db.Pool.Acquire(r.db.Ctx)
	if err != nil {
		return err
	}
	defer conn.Release()
	query, err := deleteUserQuery(id)
	if err != nil {
		return err
	}
	_, err = conn.Exec(r.db.Ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func createUserQuery(data dto.CreateUserDTO) (string, error) {
	dialect := goqu.Dialect("postgres")
	data2insert := goqu.Record{
		"created_at": time.Now(),
		"updated_at": time.Now(),
		"email":      data.Email,
		"first_name": data.FirstName,
		"last_name":  data.LastName,
		"is_active":  data.IsActive,
		"is_2fa":     data.Is2fa,
		"role_id":    data.RoleID,
		"password":   data.Password,
		"token_2fa":  data.Token2fa,
	}
	if data.Phone != "" {
		data2insert["phone"] = data.Phone
	}
	ds := dialect.Insert(usersTable).Rows(data2insert).Returning("id", "created_at", "updated_at", "email",
		"first_name", "last_name", "phone", "is_active", "is_2fa", "role_id", "password", "token_2fa")

	return toSQL(ds)
}

func updateUserQuery(data dto.UpdateUserDTO) (string, error) {
	dialect := goqu.Dialect("postgres")
	ds := dialect.Update(usersTable)
	record := goqu.Record{}
	if data.Email != nil {
		record["email"] = data.Email
	}
	if data.FirstName != nil {
		record["first_name"] = data.FirstName
	}
	if data.LastName != nil {
		record["last_name"] = data.LastName
	}
	if data.Phone != nil {
		record["phone"] = data.Phone
	}
	if data.IsActive != nil {
		record["is_active"] = data.IsActive
	}
	if data.Is2fa != nil {
		record["is_2fa"] = data.Is2fa
	}
	if data.RoleID != nil {
		record["role_id"] = data.RoleID
	}
	if data.Password != nil {
		record["password"] = data.Password
	}
	ds = ds.Set(record).Where(goqu.Ex{"id": data.ID}).Returning("id", "created_at", "updated_at", "email",
		"first_name", "last_name", "phone", "is_active", "is_2fa", "role_id", "password", "token_2fa")

	return toSQL(ds)
}

func retrieveUserQuery(id uint64) (string, error) {
	dialect := goqu.Dialect("postgres")
	ds := dialect.Select("id", "created_at", "updated_at", "email",
		"first_name", "last_name", "phone", "is_active", "is_2fa", "role_id", "password", "token_2fa").From(
		usersTable).Where(goqu.Ex{"id": id})

	return toSQL(ds)
}

func searchUserQuery(data dto.SearchUserDTO) (string, error) {
	dialect := goqu.Dialect("postgres")
	ds := dialect.Select("id", "created_at", "updated_at", "email",
		"first_name", "last_name", "phone", "is_active", "is_2fa", "role_id", "password", "token_2fa").From(usersTable)
	// call filters
	if data.ID != nil {
		ds = ds.Where(goqu.Ex{"id": *data.ID})
	}
	if data.Email != nil {
		ds = ds.Where(goqu.Ex{"email": *data.Email})
	}
	if data.FirstName != nil {
		ds = ds.Where(goqu.Ex{
			"first_name": goqu.Op{"like": *data.FirstName},
		})
	}
	if data.LastName != nil {
		ds = ds.Where(goqu.Ex{
			"last_name": goqu.Op{"like": *data.LastName},
		})
	}
	if data.Phone != nil {
		ds = ds.Where(goqu.Ex{
			"phone": goqu.Op{"like": *data.Phone},
		})
	}
	if data.IsActive != nil {
		ds = ds.Where(goqu.Ex{"is_active": *data.IsActive})
	}
	if data.Is2fa != nil {
		ds = ds.Where(goqu.Ex{"is_2fa": *data.Is2fa})
	}
	if data.RoleID != nil {
		ds = ds.Where(goqu.Ex{"role_id": *data.RoleID})
	}

	ds = filterByCreatedAt(ds, data.CreatedAtFrom, data.CreatedAtTo)
	ds = ordering(ds, data.OrderBy)
	ds = slicing(ds, data.Offset, data.Limit)

	return toSQL(ds)
}

func deleteUserQuery(id uint64) (string, error) {
	dialect := goqu.Dialect("postgres")
	ds := dialect.Delete(usersTable).Where(goqu.Ex{"id": id})

	return toSQL(ds)
}
