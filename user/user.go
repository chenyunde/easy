package user

import "easy/mysql/userInfo"

type UInfo struct {
	UserName string
	Role     string
	PassWord string
	UserID   string
	RoleID   string
}

func (u *UInfo) UInfoCheck() bool {
	password, _ := userInfo.GetUserInfoByName(u.UserName)
	if u.PassWord == password {
		return true
	}
	return false
}

// 初始化
func (u *UInfo) Init() UInfo {
	return UInfo{}
}
