package models

// CorpPermanentCode 授权企业永久授权码
// Deprecated
type CorpPermanentCode struct {
	Id              uint   `gorm:"primaryKey;autoIncrement;column:fi_corp_access_token_id"`
	CorpId          uint   `json:"corp_id" gorm:"column:fi_corp_auth_id"`
	AuthCorpId      string `json:"auth_corp_id" gorm:"column:fs_corp_app_id"`
	PermanentCode   string `json:"permanent_code" gorm:"column:fs_permanent_code;varchar(512)"`
	IsCustomizedApp bool   `json:"is_customized_app" gorm:"column:fi_is_customized_app"`
	BizModel
}

func (CorpPermanentCode) TableName() string {
	return "ts_suite_corp_permanent_code"
}
