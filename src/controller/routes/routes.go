package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/giselescarvalho/CRUDemGo/src/controller"
)

func InitRoutes(r *gin.RouterGroup) {

	//r.GET("/getUserById/:userId", func(c *gin.Context) {}, func(c *gin.Context) {})
	r.GET("/getUserById/:userId", controller.FindUserById)
	r.GET("/getUserByEmail/:userEmail", controller.FindUserByEmail)
	r.POST("/post/createUser", controller.CreateUser)
	r.PUT("/updateUser/:userId", controller.UpdateUser)
	r.DELETE("/deleteUser/:userId", controller.DeleteUser)

}
