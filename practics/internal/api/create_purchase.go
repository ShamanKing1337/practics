package service

import (
	"context"
	"fmt"
	practics "practics/api/api"
	"practics/internal/mapping"
)

func (i *Implementation) CreatePurchase(ctx context.Context, request *practics.CreatePurchaseRequest) (*practics.CreatePurchaseResponse, error) {
	res, err := i.purchasesService.CreatePurchase(ctx, mapping.CreatePurchaseRequestToPurchase(request))

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return mapping.PurchaseToCreatePurchaseResponse(res), nil
}
