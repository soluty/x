package entity

import "time"

// Comment 代表一个评论
type Comment struct {
	Id        int       // 唯一id
	Body      string    // 内容
	Author    User      // 作者
	CreatedAt time.Time // 创建时间
}