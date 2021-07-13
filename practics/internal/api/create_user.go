package service

import (
	"context"
	"fmt"
	practics "practics/api/api"
	"practics/internal/mapping"
)

func (i *Implementation) CreateUser(ctx context.Context, request *practics.CreateUserRequest) (*practics.CreateUserResponse, error){
	res, err := i.usersService.CreateUser(ctx, mapping.CreateUserRequestToUsers(request))

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return  mapping.UsersToCreateUserResponse(res), nil
}

