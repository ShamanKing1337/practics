package pkg

import (
	"context"
	"practics/internal/datastruct"
	"practics/internal/repository"
	"time"
)

type PurchasesService interface {
	CreatePurchase(ctx context.Context, req *datastruct.Purchases) (*datastruct.Purchases, error)
	DeletePurchase(ctx context.Context, id int64)  error
}


type purchasesService struct {
	dao repository.DAO
}

func (p *purchasesService) CreatePurchase(ctx context.Context, req *datastruct.Purchases) (*datastruct.Purchases, error) {
	users, err := p.dao.NewPurchasesQuery().Create(req)
	if err != nil {
		return nil,err
	}
	return users, nil
}

func (p *purchasesService) DeletePurchase(ctx context.Context, id int64)  error {
	purchase, err := p.dao.NewPurchasesQuery().Delete(id)
	if err != nil {
		return err
	}
	returnedPurchase := &datastruct.ReturnedPurchases{
		ID:           purchase.ID,
		UserId:       purchase.UserId,
		ItemId:       purchase.ItemId,
		PurchaseTime: purchase.PurchaseTime,
		ReturnTime:   time.Now(),
	}
	err = p.dao.NewReturnedPurchasesQuery().Create(returnedPurchase)
	if err != nil {
		return err
	}
	return  nil
}


func NewPurchasesService(dao repository.DAO) PurchasesService {
	return &purchasesService{dao: dao}
}


