package wework

import (
	"crypto/sha1"
	"fmt"
	"github.com/dgraph-io/badger/v3"
	"github.com/go-laoji/wecom-go-sdk/v2/internal"
	"io"
	"math/rand"
	"sort"
	"strings"
	"time"
)

type TicketResponse struct {
	internal.BizResponse
	Ticket    string `json:"ticket"`
	ExpiresIn int    `json:"expires_in"`
}

func (ww *weWork) GetJsApiTicket(corpId uint) (resp TicketResponse) {
	var item *badger.Item
	var err error
	err = ww.cache.View(func(txn *badger.Txn) error {
		item, err = txn.Get([]byte(fmt.Sprintf("ticket-%v", corpId)))
		if err == badger.ErrKeyNotFound {
			return err
		}
		item.Value(func(val []byte) error {
			resp.Ticket = string(val)
			return nil
		})
		return err
	})
	if err == badger.ErrKeyNotFound {
		_, err := ww.getRequest(corpId).SetResult(&resp).Get("/cgi-bin/get_jsapi_ticket")
		if err != nil {
			resp.ErrCode = 500
			resp.ErrorMsg = err.Error()
		} else {
			ww.cache.Update(func(txn *badger.Txn) error {
				entry := badger.NewEntry([]byte(fmt.Sprintf("ticket-%v", corpId)), []byte(resp.Ticket)).
					WithTTL(time.Second * time.Duration(resp.ExpiresIn))
				err = txn.SetEntry(entry)
				return err
			})
		}
	}
	return
}

func (ww *weWork) GetJsApiAgentTicket(corpId uint, agentId int) (resp TicketResponse) {
	var item *badger.Item
	var err error
	err = ww.cache.View(func(txn *badger.Txn) error {
		item, err = txn.Get([]byte(fmt.Sprintf("ticket-%v-%v", corpId, agentId)))
		if err == badger.ErrKeyNotFound {
			return err
		}
		item.Value(func(val []byte) error {
			resp.Ticket = string(val)
			return nil
		})
		return err
	})
	if err == badger.ErrKeyNotFound {
		_, err := ww.getRequest(corpId).SetResult(&resp).
			SetQueryParam("type", "agent_config").Get("/cgi-bin/ticket/get")
		if err != nil {
			resp.ErrCode = 500
			resp.ErrorMsg = err.Error()
		} else {
			ww.cache.Update(func(txn *badger.Txn) error {
				entry := badger.NewEntry([]byte(fmt.Sprintf("ticket-%v-%v", corpId, agentId)), []byte(resp.Ticket)).
					WithTTL(time.Second * time.Duration(resp.ExpiresIn))
				err = txn.SetEntry(entry)
				return err
			})
		}
	}
	return
}

type JsTicketSignatureResponse struct {
	NonceStr string `json:"noncestr"`
	//JsapiTicket string `json:"jsapi_ticket,omitempty"`
	Timestamp int64  `json:"timestamp"`
	Signature string `json:"signature"`
}

const (
	letterBytes = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

func randString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}
	return string(b)
}

func (ww *weWork) GetConfigSignature(corpId uint, referer string) (resp JsTicketSignatureResponse) {
	noncestr := randString(16)
	timestamp := time.Now().Unix()
	sl := []string{fmt.Sprintf("noncestr=%s", noncestr),
		fmt.Sprintf("jsapi_ticket=%s", ww.GetJsApiTicket(corpId).Ticket),
		fmt.Sprintf("timestamp=%v", timestamp),
		fmt.Sprintf("url=%s", referer),
	}
	sort.Strings(sl)
	s := sha1.New()
	io.WriteString(s, strings.Join(sl, "&"))
	signature := fmt.Sprintf("%x", s.Sum(nil))
	//resp.JsapiTicket = mp.getJsTicket() 前端调用config时不需要此参数为安全考虑不输出到前端
	resp.NonceStr = noncestr
	resp.Timestamp = timestamp
	resp.Signature = signature
	return

}

func (ww *weWork) GetAgentConfigSignature(corpId uint, agentId int, referer string) (resp JsTicketSignatureResponse) {
	noncestr := randString(16)
	timestamp := time.Now().Unix()
	sl := []string{fmt.Sprintf("noncestr=%s", noncestr),
		fmt.Sprintf("jsapi_ticket=%s", ww.GetJsApiAgentTicket(corpId, agentId).Ticket),
		fmt.Sprintf("timestamp=%v", timestamp),
		fmt.Sprintf("url=%s", referer),
	}
	sort.Strings(sl)
	s := sha1.New()
	io.WriteString(s, strings.Join(sl, "&"))
	signature := fmt.Sprintf("%x", s.Sum(nil))
	//resp.JsapiTicket = mp.getJsTicket() 前端调用config时不需要此参数为安全考虑不输出到前端
	resp.NonceStr = noncestr
	resp.Timestamp = timestamp
	resp.Signature = signature
	return

}
