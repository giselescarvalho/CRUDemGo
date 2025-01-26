package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/giselescarvalho/CRUDemGo/src/configuration/apperror"
)

func FindUserById(c *gin.Context) {
	err := apperror.NewBadRequestError("chamou de forma errada")
	c.JSON(err.Code, err)
}

func FindUserByEmail(c *gin.Context) {}
