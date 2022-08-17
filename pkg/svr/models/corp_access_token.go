package models

// CorpAccessToken 授权企业 access token
// Deprecated: 由缓存接管，不存数据库
type CorpAccessToken struct {
	ID          uint   `gorm:"primaryKey;autoIncrement;column:fi_corp_access_token_id"`
	CorpId      uint   `json:"corp_id" gorm:"column:fi_corp_auth_id"`
	AccessToken string `json:"access_token" gorm:"column:fs_access_token"`
	ExpiresIn   int    `json:"expires_in" gorm:"column:fi_expires_in"`
	BizModel
}

func (CorpAccessToken) TableName() string {
	return "ts_suite_corp_access_token"
}
