package repository

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

type DAO interface {
	NewUsersQuery() UsersQuery
	NewItemsQuery() ItemsQuery
	NewPurchasesQuery() PurchasesQuery
	NewReturnedPurchasesQuery() ReturnedPurchasesQuery
}

type dao struct {
	db sqlx.DB
}

func (d *dao)NewUsersQuery() UsersQuery{
	return &usersQuery{db: d.db}
}

func (d *dao)NewItemsQuery() ItemsQuery{
	return &itemsQuery{db: d.db}
}

func (d *dao)NewPurchasesQuery() PurchasesQuery{
	return &purchasesQuery{db: d.db}
}

func (d *dao)NewReturnedPurchasesQuery() ReturnedPurchasesQuery{
	return &returnedPurchasesQuery{db: d.db}
}


func NewDAO(db sqlx.DB) DAO {
	return &dao{db}
}

func NewDBBalancer(ctx context.Context)  *sqlx.DB {
	db, err := sqlx.Connect("postgres", "user=vasilyevm dbname=practics host=localhost port=5432 password=ozon sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("priv")

	return db
}

