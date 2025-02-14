package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gustavoz65/CRUD-GO-API.git/src/configuration/rest_err"
	"github.com/gustavoz65/CRUD-GO-API.git/src/controller/model/request"
)

func CreateUsersController(c *gin.Context) {
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restErr := rest_err.NewBadRequestError(
			fmt.Sprint("There are some incorrect filds, error%\n", err.Error()))

		c.JSON(restErr.Code, restErr)
		return
	}

	fmt.Println(userRequest)
}
