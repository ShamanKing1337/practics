package datastruct

import "time"

const PurchasesTableName = "purchases"

type Purchases struct {
	ID	int64	`db:"id"`	//primary_key
	UserId	int64	`db:"user_id"`
	ItemId	int64	`db:"item_id"`
	PurchaseTime	time.Time	`db:"purchase_time"`
}
