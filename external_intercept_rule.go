package wework

import (
	"github.com/go-laoji/wecom-go-sdk/v2/internal"
)

type InterceptRule struct {
	RuleName        string   `json:"rule_name" validate:"required,max=20"`
	WordList        []string `json:"word_list" validate:"required,max=300"`
	SemanticsList   []int    `json:"semantics_list"`
	InterceptType   int      `json:"intercept_type" validate:"required,oneof=1 2"`
	ApplicableRange struct {
		UserList       []string `json:"user_list" validate:"required_without=DepartmentList,max=1000"`
		DepartmentList []uint   `json:"department_list" validate:"required_without=UserList,max=1000"`
	} `json:"applicable_range" validate:"required"`
}

type AddInterceptRuleResponse struct {
	internal.BizResponse
	RuleId string `json:"rule_id"`
}

// AddInterceptRule 新建敏感词规则
// https://open.work.weixin.qq.com/api/doc/90001/90143/95130#%E6%96%B0%E5%BB%BA%E6%95%8F%E6%84%9F%E8%AF%8D%E8%A7%84%E5%88%99
func (ww *weWork) AddInterceptRule(corpId uint, interceptRule InterceptRule) (resp AddInterceptRuleResponse) {
	if ok := validate.Struct(interceptRule); ok != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = ok.Error()
	}
	_, err := ww.getRequest(corpId).SetBody(interceptRule).SetResult(&resp).
		Post("/cgi-bin/externalcontact/add_intercept_rule")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}

type GetInterceptRuleListResponse struct {
	internal.BizResponse
	RuleList []struct {
		RuleID     string `json:"rule_id"`
		RuleName   string `json:"rule_name"`
		CreateTime int    `json:"create_time"`
	} `json:"rule_list"`
}

// GetInterceptRuleList 获取敏感词规则列表
// https://open.work.weixin.qq.com/api/doc/90001/90143/95130#%E8%8E%B7%E5%8F%96%E6%95%8F%E6%84%9F%E8%AF%8D%E8%A7%84%E5%88%99%E5%88%97%E8%A1%A8
func (ww *weWork) GetInterceptRuleList(corpId uint) (resp GetInterceptRuleListResponse) {
	_, err := ww.getRequest(corpId).SetResult(&resp).
		Get("/cgi-bin/externalcontact/get_intercept_rule_list")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}

type GetInterceptRuleResponse struct {
	internal.BizResponse
	Rule struct {
		RuleId string
		InterceptRule
	}
}

// GetInterceptRule 获取敏感词规则详情
// https://open.work.weixin.qq.com/api/doc/90001/90143/95130#%E8%8E%B7%E5%8F%96%E6%95%8F%E6%84%9F%E8%AF%8D%E8%A7%84%E5%88%99%E8%AF%A6%E6%83%85
func (ww *weWork) GetInterceptRule(corpId uint, ruleId string) (resp GetInterceptRuleResponse) {
	h := H{}
	h["rule_id"] = ruleId
	_, err := ww.getRequest(corpId).SetBody(h).SetResult(&resp).
		Post("/cgi-bin/externalcontact/get_intercept_rule")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}

type UpdateInterceptRuleRequest struct {
	RuleId    string   `json:"rule_id" validate:"required"`
	RuleName  string   `json:"rule_name,omitempty"`
	WordList  []string `json:"word_list,omitempty"`
	ExtraRule struct {
		SemanticsList []int `json:"semantics_list"`
	} `json:"extra_rule,omitempty"`
	InterceptType      int `json:"intercept_type,omitempty" validate:"omitempty,oneof=1 2"`
	AddApplicableRange struct {
		UserList       []string `json:"user_list" validate:"required_without=DepartmentList,max=1000"`
		DepartmentList []uint   `json:"department_list" validate:"required_without=UserList,max=1000"`
	} `json:"add_applicable_range,omitempty"`
	RemoveApplicableRange struct {
		UserList       []string `json:"user_list" validate:"required_without=DepartmentList,max=1000"`
		DepartmentList []uint   `json:"department_list" validate:"required_without=UserList,max=1000"`
	} `json:"remove_applicable_range,omitempty"`
}

// UpdateInterceptRule 修改敏感词规则
// https://open.work.weixin.qq.com/api/doc/90001/90143/95130#%E4%BF%AE%E6%94%B9%E6%95%8F%E6%84%9F%E8%AF%8D%E8%A7%84%E5%88%99
func (ww *weWork) UpdateInterceptRule(corpId uint, request UpdateInterceptRuleRequest) (resp internal.BizResponse) {
	if ok := validate.Struct(request); ok != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = ok.Error()
	}
	_, err := ww.getRequest(corpId).SetBody(request).SetResult(&resp).
		Post("/cgi-bin/externalcontact/update_intercept_rule")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}

// DeleteInterceptRule 删除敏感词规则
// https://open.work.weixin.qq.com/api/doc/90001/90143/95130#%E5%88%A0%E9%99%A4%E6%95%8F%E6%84%9F%E8%AF%8D%E8%A7%84%E5%88%99
func (ww *weWork) DeleteInterceptRule(corpId uint, ruleId string) (resp internal.BizResponse) {
	h := H{}
	h["rule_id"] = ruleId
	_, err := ww.getRequest(corpId).SetBody(h).SetResult(&resp).
		Post("/cgi-bin/externalcontact/del_intercept_rule")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}
