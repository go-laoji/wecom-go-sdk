package models

// CorpAuthUserInfo 授权应用管理员信息
type CorpAuthUserInfo struct {
	ID         uint   `gorm:"primaryKey;autoIncrement;column:fi_corp_user_auth_id"`
	CorpId     uint   `json:"corp_id" gorm:"column:fi_corp_auth_id"`
	UserId     string `json:"userid" gorm:"column:fs_user_id;varchar(255)"`
	OpenUserId string `json:"open_userid" gorm:"column:fs_open_user_id;varchar(255)"`
	Name       string `json:"name" gorm:"column:fs_name;varchar(200)"`
	Avatar     string `json:"avatar" gorm:"column:fs_avatar;varchar(255)"`
	BizModel
}

func (CorpAuthUserInfo) TableName() string {
	return "ts_suite_corp_auth_user_info"
}
