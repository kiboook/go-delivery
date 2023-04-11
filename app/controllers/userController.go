package controllers

import (
	"github.com/gin-gonic/gin"
	"go-delivery/app/services"
	"net/http"
)

type UserController struct {
}

var userService = new(services.UserService)

// List userController godoc
// @tags 회원
// @Summary 회원 목록 조회
// @Description 회원 목록 조회
// @Accept  json
// @Produce  json
// @Router /users [get]
// @Success 200
// @Failure 400
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
