package service

import (
	"context"
	"fmt"
	practics "practics/api/api"
	"practics/internal/mapping"
)

func (i *Implementation) CreateItem(ctx context.Context, request *practics.CreateItemRequest) (*practics.CreateItemResponse, error){
	res, err := i.itemsService.CreateItem(ctx, mapping.CreateItemRequestToItem(request))

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return  mapping.ItemToCreateItemResponse(res), nil
}


