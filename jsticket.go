package wework

import (
	"encoding/json"
	"fmt"
	"github.com/dgraph-io/badger/v2"
	"github.com/go-laoji/wecom-go-sdk/internal"
	"time"
)

type TicketResponse struct {
	internal.BizResponse
	Ticket    string `json:"ticket"`
	ExpiresIn int    `json:"expires_in"`
}

func (ww weWork) GetJsApiTicket(corpId uint) (resp TicketResponse) {
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
		queryParams := ww.buildCorpQueryToken(corpId)
		body, err := internal.HttpGet(fmt.Sprintf("/cgi-bin/get_jsapi_ticket?%s", queryParams.Encode()))
		if err != nil {
			resp.ErrCode = 500
			resp.ErrorMsg = err.Error()
		} else {
			json.Unmarshal(body, &resp)
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

func (ww weWork) GetJsApiAgentTicket(corpId uint, agentId int) (resp TicketResponse) {
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
		queryParams := ww.buildCorpQueryToken(corpId)
		queryParams.Add("type", "agent_config")
		body, err := internal.HttpGet(fmt.Sprintf("/cgi-bin/ticket/get?%s", queryParams.Encode()))
		if err != nil {
			resp.ErrCode = 500
			resp.ErrorMsg = err.Error()
		} else {
			json.Unmarshal(body, &resp)
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
