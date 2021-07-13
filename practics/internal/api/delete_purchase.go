package service

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	practics "practics/api/api"
)

func (i *Implementation) DeletePurchase(ctx context.Context, request *practics.DeletePurchaseRequest) (*emptypb.Empty, error) {
	err := i.purchasesService.DeletePurchase(ctx, request.GetId())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
