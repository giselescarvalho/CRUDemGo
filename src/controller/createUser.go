package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/giselescarvalho/CRUDemGo/src/configuration/logger"
	"github.com/giselescarvalho/CRUDemGo/src/configuration/validation"
	"github.com/giselescarvalho/CRUDemGo/src/controller/model"
	"github.com/giselescarvalho/CRUDemGo/src/controller/model/request"
	"go.uber.org/zap"
)

var (
	UserDomainInterface model.UserDomainInterface
)

func CreateUser(c *gin.Context) {

	logger.Info("Init CreateUser controller",
		zap.String("journey", "createUser"),
	)
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to marashal object, error=%s", err, zap.String("journey", "createUser"))
		errRest := validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	fmt.Println("#######\n", userRequest)

	domain := model.NewUserDomain(userRequest.Email, userRequest.Password, userRequest.Name, int(userRequest.Age))

	if err := domain.CreateUser(); err != nil {
		c.JSON(http.StatusOK, err)
		return
	}

	logger.Info("User created sucessfully", zap.String("journey", "createUser"))

	c.String(http.StatusOK, "")
}
