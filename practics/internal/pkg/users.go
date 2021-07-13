package pkg

import (
	"context"
	"fmt"
	"practics/internal/datastruct"
	"practics/internal/repository"
)

type UsersService interface {
	CreateUser(ctx context.Context, req *datastruct.Users) (*datastruct.Users, error)
}


type usersService struct {
	dao repository.DAO
}


func (p *usersService) CreateUser(ctx context.Context, req *datastruct.Users) (*datastruct.Users, error) {
	fmt.Println("dasdas")
	users, err := p.dao.NewUsersQuery().Create(req)
	if err != nil {
		return nil,err
	}
	return users, nil
}

func NewUsersService(dao repository.DAO) UsersService {
	return &usersService{dao: dao}
}