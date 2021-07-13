package datastruct

import "time"

const ReturnedPurchasesTableName = "returned_purchases"

type ReturnedPurchases struct {
	ID	int64	`db:"id"`
	UserId	int64	`db:"user_id"`
	ItemId	int64	`db:"item_id"`
	PurchaseTime	time.Time	`db:"purchase_time"`
	ReturnTime	time.Time	`db:"return_time"`
	Description	string	`db:"description"`
}
