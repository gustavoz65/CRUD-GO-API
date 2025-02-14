package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/gustavoz65/CRUD-GO-API.git/src/controller"
)

func InitRoutes(r *gin.RouterGroup) {
	r.GET("/getusersById/:userid", controller.FindUsersByIdController)
	r.GET("/getusersByEmail/:userEmail", controller.FindUsersByEmailController)
	r.POST("/createUser", controller.CreateUsersController)
	r.PUT("/uptdateUser/:userid", controller.UpdateUsersController)
	r.PATCH("/patchUser/:userid", controller.PatchUsersController)
	r.DELETE("deletUser", controller.DeleteUsersController)
}
