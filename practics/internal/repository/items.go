package repository

import (
	"github.com/jmoiron/sqlx"
	"practics/internal/datastruct"
)

type ItemsQuery interface {
	Create(items *datastruct.Items) (*datastruct.Items, error)

}

type itemsQuery struct {
	db sqlx.DB
}

func (u *itemsQuery) Create(items *datastruct.Items) (*datastruct.Items, error) {
	res, err := u.db.Queryx("insert into " + datastruct.ItemsTableName + "(title, description, cost) values ($1, $2, $3) returning *", items.Title, items.Description, items.Cost)
	if err != nil {
		return nil,err
	}
	uuu := &datastruct.Items{}
	res.Next()
	err = res.StructScan(uuu)
	if err != nil {
		return nil, err
	}
	return uuu, nil
}
