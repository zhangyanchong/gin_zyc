package tools

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Result struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func Success(c *gin.Context, msg string, data interface{}) {
	if msg == "" {
		msg = "ok"
	}

	c.JSON(http.StatusOK, Result{
		Code: 200,
		Msg:  msg,
		Data: data,
	})

}

//错误输出
func Error(c *gin.Context, msg string, data interface{}) {
	if msg == "" {
		msg = "ok"
	}

	c.JSON(http.StatusOK, Result{
		Code: 500,
		Msg:  msg,
		Data: data,
	})
}

/*页面输出**/
func HtmlOut(c *gin.Context, htmlName string, data interface{}) {
	c.HTML(http.StatusOK, htmlName+".tmpl", gin.H{
		"Data": data,
	})
}
