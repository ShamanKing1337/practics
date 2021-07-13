package repository

import (
	"github.com/jmoiron/sqlx"
	"practics/internal/datastruct"
)

type ReturnedPurchasesQuery interface {
	Create(returnedPurchases *datastruct.ReturnedPurchases) error

}

type returnedPurchasesQuery struct {
	db sqlx.DB
}

func (u *returnedPurchasesQuery) Create(returnedPurchases *datastruct.ReturnedPurchases) error {
	res, err := u.db.Queryx("insert into " + datastruct.ReturnedPurchasesTableName + "(id, user_id, item_id, purchase_time, return_time, description) values ($1, $2, $3, $4, $5, $6) returning *",
		returnedPurchases.ID, returnedPurchases.UserId, returnedPurchases.ItemId, returnedPurchases.PurchaseTime, returnedPurchases.ReturnTime, returnedPurchases.Description)
	if err != nil {
		return err
	}
	uuu := &datastruct.ReturnedPurchases{}
	res.Next()
	err = res.StructScan(uuu)
	if err != nil {
		return err
	}
	return nil
}
