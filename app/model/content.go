package model

type Content struct {
	Id         int    `json:"id"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	CreateTime int    `json:"create_time""`
}

type ContentReq struct {
	Title   string `form:"title" json:"title" binding:"required"`
	Content string `form:"content" json:"content" binding:"required"`
}

func (Content) TableName() string {
	return "content"
}
