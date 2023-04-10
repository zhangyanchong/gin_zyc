package model

type Content struct {
	Id         int    `json:"id"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	CreateTime int    `json:"create_time""`
}

func (Content) TableName() string {
	return "content"
}
