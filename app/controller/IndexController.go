package controller

import (
	"gin_zyc/app/tools"
	"github.com/gin-gonic/gin"
)

type IndexController struct {
}

func (a *IndexController) Hello(c *gin.Context) {
	tools.Success(c, "hello", nil)
}
