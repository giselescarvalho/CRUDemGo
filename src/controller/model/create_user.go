package model

import (
	"fmt"

	"github.com/giselescarvalho/CRUDemGo/src/configuration/apperror"
	"github.com/giselescarvalho/CRUDemGo/src/configuration/logger"
	"go.uber.org/zap"
)

func (ud *UserDomain) CreateUser() *apperror.RestErr {

	logger.Info("Init createUSer model", zap.String("journey", "createUser"))

	ud.EncryptPassword()

	fmt.Println(ud)

	return nil
}
