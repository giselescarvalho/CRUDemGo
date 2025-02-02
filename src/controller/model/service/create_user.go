package service

import (
	"fmt"

	"github.com/giselescarvalho/CRUDemGo/src/configuration/apperror"
	"github.com/giselescarvalho/CRUDemGo/src/configuration/logger"
	"github.com/giselescarvalho/CRUDemGo/src/controller/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) CreateUser(userDomain model.UserDomainInterface) *apperror.RestErr {

	logger.Info("Init createUSer model", zap.String("journey", "createUser"))

	userDomain.EncryptPassword()

	fmt.Println(userDomain.GetPassword())

	return nil
}
