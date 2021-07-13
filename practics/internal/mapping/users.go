package mapping

import (
	"database/sql"
	practics "practics/api/api"
	"practics/internal/datastruct"
)

func CreateUserRequestToUsers(request *practics.CreateUserRequest) *datastruct.Users {
	return &datastruct.Users{
		Login:     request.GetLogin(),
		Password:  request.GetPassword(),
		LastName:  request.GetLastname(),
		FirstName: sql.NullString{request.GetFirstname(), true},
	}
}
func UsersToCreateUserResponse(users *datastruct.Users) *practics.CreateUserResponse {
	return &practics.CreateUserResponse{
		Id: users.ID,
		Login:    users.Login,
		Password: users.Password,
		Lastname: users.LastName,
		Firstname: users.FirstName.String,
	}
}
