package service

import (
	"github.com/giselescarvalho/CRUDemGo/src/configuration/apperror"
	"github.com/giselescarvalho/CRUDemGo/src/controller/model"
)

func (ud *userDomainService) FindUser(string) (*model.UserDomainInterface, *apperror.RestErr) {

	//	logger.Info("Init createUSer model", zap.String("journey", "createUser"))

	//	ud.EncryptPassword()

	//	fmt.Println(ud)

	return nil, nil
}
