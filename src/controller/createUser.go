package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/giselescarvalho/CRUDemGo/src/configuration/logger"
	"github.com/giselescarvalho/CRUDemGo/src/configuration/validation"
	"github.com/giselescarvalho/CRUDemGo/src/controller/model/request"
	"github.com/giselescarvalho/CRUDemGo/src/controller/model/response"
	"go.uber.org/zap"
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

	response := response.UserResponse{
		ID:    "test",
		Email: userRequest.Email,
		Name:  userRequest.Name,
		Age:   userRequest.Age,
	}

	logger.Info("User created sucessfully", zap.String("journey", "createUser"))

	c.JSON(http.StatusOK, response)
}
