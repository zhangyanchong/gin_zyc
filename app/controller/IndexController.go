package controller

import (
	"gin_zyc/app/tools"
	"github.com/gin-gonic/gin"
)

type indexController struct {
}

var IndexController indexController

func (a *indexController) Hello(c *gin.Context) {
	tools.Success(c, "hello", nil)
}
