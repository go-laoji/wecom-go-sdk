package models

type Agent struct {
	Id               uint   `gorm:"primaryKey;autoIncrement;column:fi_corp_access_token_id"`
	CorpId           uint   `json:"corp_id" gorm:"column:fi_corp_auth_id"`
	AgentId          int    `json:"agentid" gorm:"column:fi_agent_id"`
	AuthCorpId       string `json:"auth_corp_id" gorm:"column:fs_corp_app_id"`
	PermanentCode    string `json:"permanent_code" gorm:"column:fs_permanent_code;varchar(512)"`
	Name             string `json:"name" gorm:"column:fs_name;varchar(200)"`
	RoundLogoURL     string `json:"round_logo_url" gorm:"column:fs_round_logo_url;varchar(255)"`
	SquareLogoURL    string `json:"square_logo_url" gorm:"column:fs_square_logo_url;varchar(255)"`
	AuthMode         int    `json:"auth_mode,omitempty" gorm:"column:fi_auth_mode;int(2)"`
	IsCustomizedApp  bool   `json:"is_customized_app,omitempty" gorm:"column:fi_is_customized_app"`
	AuthFromThirdapp bool   `json:"auth_from_thirdapp,omitempty" gorm:"column:fi_auth_from_thirdapp"`
	Privilege        struct {
		Level      int      `json:"level"`
		AllowParty []int    `json:"allow_party"`
		AllowUser  []string `json:"allow_user"`
		AllowTag   []int    `json:"allow_tag"`
		ExtraParty []int    `json:"extra_party"`
		ExtraUser  []string `json:"extra_user"`
		ExtraTag   []int    `json:"extra_tag"`
	} `json:"privilege,omitempty" gorm:"-"` //TODO:授权信息处理
	SharedFrom struct {
		Corpid    string `json:"corpid"`
		ShareType int    `json:"share_type"`
	} `json:"shared_from" gorm:"-"` //TODO:授权信息处理
}

func (Agent) TableName() string {
	return "ts_suite_corp_agent"
}
