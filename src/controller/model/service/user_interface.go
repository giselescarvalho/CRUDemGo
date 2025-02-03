package service

import (
	"github.com/giselescarvalho/CRUDemGo/src/configuration/apperror"
	"github.com/giselescarvalho/CRUDemGo/src/controller/model"
)

type UserDomainService interface {
	CreateUser(model.UserDomainInterface) *apperror.RestErr
	UpdateUser(string, model.UserDomainInterface) *apperror.RestErr
	FindUser(string) (*model.UserDomainInterface, *apperror.RestErr)
	DeleteUser(string) *apperror.RestErr
}

type userDomainService struct {
}

func NewUserDomainService() UserDomainService {
	return &userDomainService{}
}
