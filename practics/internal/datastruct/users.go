package datastruct

import "database/sql"

const UsersTableName = "users"

type Users struct {
	ID	int64	`db:"id"`	//primary_key
	Login	string	`db:"login"`
	Password	string	`db:"password"`
	LastName	string	`db:"last_name"`
	FirstName	sql.NullString	`db:"first_name"`
}
