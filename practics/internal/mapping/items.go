package mapping

import (
	practics "practics/api/api"
	"practics/internal/datastruct"
)

func CreateItemRequestToItem(request *practics.CreateItemRequest) *datastruct.Items {
	return &datastruct.Items{
		Title:       request.GetTitle(),
		Description: request.GetDescription(),
		Cost:        request.GetCost(),
	}
}

func ItemToCreateItemResponse(items *datastruct.Items) *practics.CreateItemResponse {
	return &practics.CreateItemResponse{
		Id:          items.ID,
		Title:       items.Title,
		Description: items.Description,
		Cost:        items.Cost,
	}
}

