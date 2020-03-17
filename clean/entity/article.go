package entity

import "time"

// Article 代表一篇文章
type Article struct {
	Slug          string    // 文章的唯一id，显示在浏览器地址栏
	Title         string    // 标题
	Description   string    // 简介
	Body          string    // 内容
	TagList       []string  // 标签
	FavoriteCount int       // 点赞数
	Author        User      // 作者
	Comments      []Comment // 评论
	CreatedAt     time.Time // 创建时间
	UpdatedAt     time.Time // 修改时间
}
