package wework

import (
	"encoding/json"
	"fmt"
	"github.com/dgraph-io/badger/v2"
	"github.com/go-laoji/wecom-go-sdk/internal"
	"net/url"
	"time"
)

type providerAccessTokenResponse struct {
	internal.BizResponse
	ProviderAccessToken string `json:"provider_access_token"`
	ExpiresIn           int    `json:"expires_in"`
}

func (ww weWork) requestProviderToken() (resp providerAccessTokenResponse) {
	apiUrl := "/cgi-bin/service/get_provider_token"
	params := H{}
	params["corpid"] = ww.corpId
	params["provider_secret"] = ww.providerSecret
	var data []byte
	var err error
	if data, err = internal.HttpPost(apiUrl, params); err != nil {
		resp.ErrCode = 400
		resp.ErrorMsg = err.Error()
		return
	}
	if err = json.Unmarshal(data, &resp); err != nil {
		resp.ErrCode = 400
		resp.ErrorMsg = err.Error()
		return
	}
	return resp
}

func (ww weWork) getProviderToken() (token string) {
	var err error
	var item *badger.Item
	err = ww.cache.View(func(txn *badger.Txn) error {
		item, err = txn.Get([]byte("providerToken"))
		if err == badger.ErrKeyNotFound {
			return err
		}
		item.Value(func(val []byte) error {
			token = string(val)
			return nil
		})
		return err
	})
	if err == badger.ErrKeyNotFound {
		if resp := ww.requestProviderToken(); resp.ErrCode != 0 {
			panic(resp)
			//logger.Error(err.Error())
		} else {
			token = resp.ProviderAccessToken
			ww.cache.Update(func(txn *badger.Txn) error {
				entry := badger.NewEntry([]byte("providerToken"), []byte(token)).
					WithTTL(time.Second * time.Duration(resp.ExpiresIn))
				err = txn.SetEntry(entry)
				return err
			})
		}
	}
	return token
}

type GetLoginInfoResponse struct {
	internal.BizResponse
	UserType int `json:"usertype"`
	UserInfo struct {
	} `json:"user_info"`
	CorpInfo struct {
		CorpId string `json:"corpid"`
	} `json:"corp_info"`
	Agent []struct {
		AgentId  int `json:"agentid"`
		AuthType int `json:"auth_type"`
	} `json:"agent"`
	AuthInfo []struct {
		Department []struct {
			Id       int  `json:"id"`
			Writable bool `json:"writable"`
		} `json:"department"`
	} `json:"auth_info"`
}

// GetLoginInfo 获取登录用户信息
// https://open.work.weixin.qq.com/api/doc/90001/90143/91125
func (ww weWork) GetLoginInfo(authCode string) (resp GetLoginInfoResponse) {
	queryParams := url.Values{}
	queryParams.Add("access_token", ww.getProviderToken())
	apiUrl := fmt.Sprintf("/cgi-bin/service/get_login_info?%s", queryParams.Encode())
	h := H{}
	h["auth_code"] = authCode

	body, err := internal.HttpPost(apiUrl, h)
	if err != nil {
		logger.Sugar().Error(err)
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
		return
	}
	if err = json.Unmarshal(body, &resp); err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}
