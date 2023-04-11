package html

import (
	"fmt"
	"gin_zyc/app/model"
	"gin_zyc/app/tools"
	"gin_zyc/app/tools/db"
	"github.com/gin-gonic/gin"
	"time"
	"unicode/utf8"
)

type htmlController struct {
}

var HtmlController htmlController

func (a *htmlController) Login(c *gin.Context) {
	tools.HtmlOut(c, "login/login", nil)
}

func (a *htmlController) BlogAdd(c *gin.Context) {
	tools.HtmlOut(c, "blog/blogadd", nil)
}

func (a *htmlController) BlogDelApi(c *gin.Context) {
	id := c.PostForm("id")

	rs := db.DB().Delete(&model.Content{}, id)
	if rs.Error != nil {
		tools.Success(c, "删除失败", nil)
	}
	//判断是否删除成功了
	if rs.RowsAffected > 0 {
		tools.Success(c, "删除成功", nil)
	} else {
		tools.Error(c, "删除失败", nil)
	}

}

func (a *htmlController) BlogAddApi(c *gin.Context) {
	var contentReq model.ContentReq
	if err := c.ShouldBind(&contentReq); err != nil {
		tools.Error(c, err.Error(), nil)
		return
	}
	if utf8.RuneCountInString(contentReq.Title) < 4 {
		tools.Error(c, "标题不能少于4个字", nil)
		return
	}
	One := model.Content{
		Title:      contentReq.Title,
		Content:    contentReq.Content,
		CreateTime: int(time.Now().Unix()),
	}
	rs := db.DB().Create(&One)
	if rs.Error != nil {
		tools.Success(c, "插入失败", nil)
	}
	tools.Success(c, "ok", One)
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

func (a *htmlController) BlogDetail(c *gin.Context) {
	id := c.Query("id")
	content := model.Content{}
	db.DB().Where("id = ?", id).Take(&content)
	fmt.Println(content)
	tools.HtmlOut(c, "blog/blogdetail", content)
}
