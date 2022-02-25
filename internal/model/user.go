package model

type User struct {
	*Model
	Uid      string `json:"uid" column:"uid;index"`
	Username string `json:"username" gorm:"column:username;index"`
	Passport string `json:"passport" gorm:"column:passport"`
	Nickname string `json:"nickname" gorm:"column:nickname"`
	Role     string `json:"role" gorm:"column:role"`
	Avatar   string `json:"avatar" gorm:"column:avatar"`
}

const UserTableName = "user"

func (u User) TableName() string {
	return UserTableName
}
