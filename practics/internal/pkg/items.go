package pkg

import (
	"context"
	"practics/internal/datastruct"
	"practics/internal/repository"
)

type ItemsService interface {
	CreateItem(ctx context.Context, req *datastruct.Items) (*datastruct.Items, error)
}


type itemsService struct {
	dao repository.DAO
}

func (p *itemsService) CreateItem(ctx context.Context, req *datastruct.Items) (*datastruct.Items, error) {
	users, err := p.dao.NewItemsQuery().Create(req)
	if err != nil {
		return nil,err
	}
	return users, nil
}

func NewItemsService(dao repository.DAO) ItemsService {
	return &itemsService{dao: dao}
}
