package mapping

import (
	"google.golang.org/protobuf/types/known/timestamppb"
	practics "practics/api/api"
	"practics/internal/datastruct"
	"time"
)

func CreatePurchaseRequestToPurchase(request *practics.CreatePurchaseRequest) *datastruct.Purchases {
	return &datastruct.Purchases{
		UserId:       request.GetUserId(),
		ItemId:       request.GetItemId(),
		PurchaseTime: time.Now(),
	}

}

func PurchaseToCreatePurchaseResponse(purchase *datastruct.Purchases) *practics.CreatePurchaseResponse {
	return &practics.CreatePurchaseResponse{
		Id:           purchase.ID,
		UserId:       purchase.UserId,
		ItemId:       purchase.ItemId,
		PurchaseTime: timestamppb.New(purchase.PurchaseTime),
	}
}

