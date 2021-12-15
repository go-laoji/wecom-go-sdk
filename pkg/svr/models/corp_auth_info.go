package models

// CorpAuthInfo 授权企业信息表
type CorpAuthInfo struct {
	ID                uint   `gorm:"primaryKey;autoIncrement;column:fi_corp_auth_id"`
	CorpId            string `json:"corp_id" gorm:"column:fs_corp_app_id"`
	CorpType          string `json:"corp_type" gorm:"column:fs_corp_type;varchar(20)"`
	CorpSquareLogoUrl string `json:"corp_square_logo_url" gorm:"column:fs_corp_square_logo_url;varchar(255)"`
	CorpUserMax       uint   `json:"corp_user_max" gorm:"column:fi_corp_user_max"`
	CorpAgentMax      uint   `json:"corp_agent_max" gorm:"column:fi_corp_agent_max"`
	CorpFullName      string `json:"corp_full_name" gorm:"column:fs_corp_full_name;varchar(255)"`
	VerifiedEndTime   uint64 `json:"verified_end_time" gorm:"column:fi_verified_end_time"`
	SubjectType       int    `json:"subject_type" gorm:"column:fi_subject_type"`
	CorpWxQrcode      string `json:"corp_wxqrcode" gorm:"column:fs_corp_wxqrcode;varchar(255)"`
	CorpScale         string `json:"corp_scale" gorm:"column:fs_corp_scale;varchar(200)"`
	CorpIndustry      string `json:"corp_industry" gorm:"column:fs_corp_industry;varchar(200)"`
	CorpSubIndustry   string `json:"corp_sub_industry" gorm:"column:fs_corp_sub_industry;varchar(200)"`
	BizModel
}

func (CorpAuthInfo) TableName() string {
	return "ts_suite_corp_auth_info"
}
