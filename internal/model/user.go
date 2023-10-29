package model

type User struct {
	Id        string `json:"_id,omitempty"`        // ID
	UserId    int    `json:"user_id,omitempty"`    // 用户ID
	Mobile    string `json:"mobile,omitempty"`     // 手机号
	Nickname  string `json:"nickname,omitempty"`   // 用户昵称
	Avatar    string `json:"avatar,omitempty"`     // 用户头像地址
	Gender    int    `json:"gender,omitempty"`     // 用户性别  0:未知  1:男   2:女
	Password  string `json:"password,omitempty"`   // 用户密码
	Motto     string `json:"motto,omitempty"`      // 用户座右铭
	Email     string `json:"email,omitempty"`      // 用户邮箱
	Birthday  string `json:"birthday,omitempty"`   // 生日
	IsRobot   int    `json:"is_robot,omitempty"`   // 是否机器人[0:否;1:是;]
	CreatedAt int64  `json:"created_at,omitempty"` // 注册时间
	UpdatedAt int64  `json:"updated_at,omitempty"` // 更新时间
}
