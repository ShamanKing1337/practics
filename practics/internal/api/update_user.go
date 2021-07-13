package service

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	practics "practics/api/api"
)

func (i *Implementation) UpdateUser(ctx context.Context, request *practics.UpdateUserRequest) (*emptypb.Empty, error) {

	return &emptypb.Empty{}, nil
}

