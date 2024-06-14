package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/mapacheco031975/meu-primeiro-cru-go/src/configuration/rest_err"
	"github.com/mapacheco031975/meu-primeiro-cru-go/src/controller/model/request"
)

func CreateUser(c *gin.Context) {

	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restErr := rest_err.NewBadResquestError(
			fmt.Sprintf("There are some incorrect filds, error=%s\n", err.Error()))

		c.JSON(restErr.Code, restErr)
		return
	}
}
