package entity

import "time"

// User 代表一个用户
type User struct {
	Id        int       // 唯一id
	Name      string    // 昵称
	Email     string    // 邮箱
	Password  string    // 密码 不能明文存储
	Bio       *string   // 简介
	ImageLink *string    // 头像链接
	Follows   []int     // 关注了哪些用户的id
	Favorites []Article // 喜欢的文章
	CreatedAt time.Time // 注册时间
	UpdatedAt time.Time // 更新时间
}

func (u *User) Update(user *User) {
	if user.Name != "" {
		u.Name = user.Name
	}
	if user.Email != "" {
		u.Email = user.Email
	}
	if user.Password != "" {
		u.Password = user.Password
	}
	if user.Bio != nil {
		u.Bio = user.Bio
	}
	if user.ImageLink != nil {
		u.ImageLink = user.ImageLink
	}
}

// u是否关注了id的用户
func (u User) IsFollow(id int) bool {
	for _, value := range u.Follows {
		if value == id {
			return true
		}
	}
	return false
}

// followeeId关注或者取消关注u
func (u *User) UpdateFollowees(followeeId int, follow bool) {
	if follow {
		u.Follows = append(u.Follows, followeeId)
		return
	}

	for i := len(u.Follows) - 1; i >= 0; i-- {
		if u.Follows[i] == followeeId {
			u.Follows = append(u.Follows[:i], u.Follows[i+1:]...)
		}
	}

	if len(u.Follows) == 0 {
		u.Follows = nil
	}
}
