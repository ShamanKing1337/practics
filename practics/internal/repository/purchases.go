package repository

import (
	"github.com/jmoiron/sqlx"
	"practics/internal/datastruct"
	"strconv"
)

type PurchasesQuery interface {
	Create(purchase *datastruct.Purchases) (*datastruct.Purchases, error)
	Delete(id int64) (*datastruct.Purchases, error)
}

type purchasesQuery struct {
	db sqlx.DB
}

func (u *purchasesQuery) Create(purchase *datastruct.Purchases) (*datastruct.Purchases, error) {
	res, err := u.db.Queryx("insert into " + datastruct.PurchasesTableName + "(user_id, item_id, purchase_time) values ($1, $2, $3) returning *", purchase.UserId, purchase.ItemId, purchase.PurchaseTime)
	if err != nil {
		return nil,err
	}
	uuu := &datastruct.Purchases{}
	res.Next()
	err = res.StructScan(uuu)
	if err != nil {
		return nil, err
	}
	return uuu, nil
}

func (u *purchasesQuery) Delete(id int64) (*datastruct.Purchases,error) {
	res, err := u.db.Queryx("delete from " + datastruct.PurchasesTableName + " where id = " + strconv.FormatInt(id, 10) + " returning *")
	if err != nil {
		return nil,err
	}
	uuu := &datastruct.Purchases{}
	res.Next()
	err = res.StructScan(uuu)
	if err != nil {
		return nil, err
	}
	return uuu, nil
}
