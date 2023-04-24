package postgres

import (
	"github.com/densmart/users-manager/internal/logger"
	"github.com/doug-martin/goqu/v9"
	"github.com/doug-martin/goqu/v9/exp"
	"time"
)

func filterByCreatedAt(ds *goqu.SelectDataset, createdAtFrom time.Time, createdAtTo time.Time) *goqu.SelectDataset {
	if !createdAtFrom.IsZero() {
		ds = ds.Where(goqu.C("created_at").Gte(createdAtFrom))
	}
	if !createdAtTo.IsZero() {
		ds = ds.Where(goqu.C("created_at").Lte(createdAtTo))
	}
	return ds
}

func ordering(ds *goqu.SelectDataset, field *string) *goqu.SelectDataset {
	f := "id"
	d := "desc"

	if field != nil {
		customOrderBy := *field
		if customOrderBy[0:1] == "-" {
			f = customOrderBy[1:]
		} else {
			f = customOrderBy
			d = "asc"
		}
	}
	switch d {
	case "desc":
		ds = ds.Order(goqu.I(f).Desc())
	case "asc":
		ds = ds.Order(goqu.I(f).Asc())
	}
	return ds
}

func slicing(ds *goqu.SelectDataset, offset *uint, limit *uint) *goqu.SelectDataset {
	if offset != nil {
		ds = ds.Offset(*offset)
	}
	if limit != nil {
		ds = ds.Limit(*limit)
	}
	return ds
}

func toSQL(ds exp.SQLExpression) (string, error) {
	sql, _, err := ds.ToSQL()
	logger.Debugf(sql)
	if err != nil {
		return "", err
	}

	return sql, nil
}
