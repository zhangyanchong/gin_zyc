package html

import (
	"gin_zyc/app/tools"
	"github.com/gin-gonic/gin"
)

type htmlController struct {
}

var HtmlController htmlController

func (a *htmlController) Login(c *gin.Context) {
	tools.HtmlOut(c, "login/login", nil)
}
