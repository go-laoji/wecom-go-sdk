package logic

import (
	"encoding/xml"
	"github.com/go-laoji/wecom-go-sdk/v2"
)

type SuiteTicketEvent struct {
	BizEvent
	SuiteTicket string `xml:"SuiteTicket"`
}

func SuiteTicketEventLogic(data []byte, ww wework.IWeWork) {
	var suiteEvent SuiteTicketEvent
	if e := xml.Unmarshal(data, &suiteEvent); e != nil {
		ww.Logger().Sugar().Error(e)
	} else {
		ww.UpdateSuiteTicket(suiteEvent.SuiteTicket)
	}
}
