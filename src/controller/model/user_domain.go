package model

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/giselescarvalho/CRUDemGo/src/configuration/apperror"
)

type UserDomain struct {
	Email    string
	Password string
	Name     string
	Age      int
}

type UserDomainInterface interface {
	CreateUser() *apperror.RestErr
	UpdateUser(string) *apperror.RestErr
	FindUser(string) (*UserDomain, apperror.RestErr)
	DeleteUser(string) apperror.RestErr
}

func NewUserDomain(
	email, password, name string, age int,
) *UserDomain {
	return &UserDomain{
		email, password, name, age,
	}
}

func (ud *UserDomain) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(ud.Password))
	ud.Password = hex.EncodeToString(hash.Sum(nil))
}
