package controllers

import (
	"github.com/gin-gonic/gin"
	"go-delivery/app/services"
	"net/http"
)

type UserController struct {
}

var userService = new(services.UserService)

func (uc *UserController) List(c *gin.Context) {
	users, err := userService.FindAll()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": http.StatusText(http.StatusBadRequest)})
	}

	c.JSON(http.StatusOK, users)
}

func (uc *UserController) Detail(c *gin.Context) {

}

func (uc *UserController) Create(c *gin.Context) {

}

func (uc *UserController) Update(c *gin.Context) {

}

func (uc *UserController) Delete(c *gin.Context) {

}
