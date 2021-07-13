package datastruct

const ItemsTableName = "items"

type Items struct {
	ID	int64	`db:"id"`	//primary_key
	Title	string	`db:"title"`
	Description	string	`db:"description"`
	Cost	int64	`db:"cost"`
}
