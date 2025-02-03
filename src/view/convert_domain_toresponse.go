package view

import (
	"github.com/giselescarvalho/CRUDemGo/src/controller/model"
	"github.com/giselescarvalho/CRUDemGo/src/controller/model/response"
)

func ConverDomainToResponse(
	userDomain model.UserDomainInterface,
) response.UserResponse {
	return response.UserResponse{
		ID:    "",
		Email: userDomain.GetEmail(),
		Name:  userDomain.GetName(),
		Age:   userDomain.GetAge(),
	}
}
