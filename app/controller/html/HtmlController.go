package html

import (
	"gin_zyc/app/model"
	"gin_zyc/app/tools"
	"gin_zyc/app/tools/db"
	"github.com/gin-gonic/gin"
	"time"
)

type htmlController struct {
}

var HtmlController htmlController

func (a *htmlController) Login(c *gin.Context) {
	tools.HtmlOut(c, "login/login", nil)
}

func (a *htmlController) BlogList(c *gin.Context) {
	//数据content获取
	content := []model.Content{}
	db.DB().Find(&content)
	//修改返回来的结果，换成新的切片操作
	type xinContent struct {
		model.Content
		CreateTimeStr string
	}
	var showContent []xinContent
	for _, v := range content {
		showContent = append(showContent, xinContent{Content: v, CreateTimeStr: time.Unix(int64(v.CreateTime), 0).Format("2006-01-02 11:11:11")})
	}
	tools.HtmlOut(c, "blog/bloglist", showContent)
}
