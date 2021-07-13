package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"practics/internal/datastruct"


)

type UsersQuery interface {
	Create(users *datastruct.Users) (*datastruct.Users, error)

}

type usersQuery struct {
	db sqlx.DB
}

func (u *usersQuery) Create(users *datastruct.Users) (*datastruct.Users, error) {
	fmt.Println(users)
	res, err := u.db.Queryx("insert into " + datastruct.UsersTableName + "(login, password, last_name, first_name) values ($1, $2, $3, $4) returning *", users.Login, users.Password, users.LastName, users.FirstName)
	if err != nil {
		return nil,err
	}
	uuu := &datastruct.Users{}
	res.Next()
	err = res.StructScan(uuu)
	if err != nil {
		return nil, err
	}
	return uuu, nil
}

