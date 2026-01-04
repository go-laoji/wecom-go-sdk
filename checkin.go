package wework

import (
	"github.com/go-laoji/wecom-go-sdk/v2/internal"
)

type CheckInRecord struct {
	Userid         string   `json:"userid" validate:"required"`
	CheckinTime    int64    `json:"checkin_time" validate:"required"`
	LocationTitle  string   `json:"location_title" validate:"required"`
	LocationDetail string   `json:"location_detail" validate:"required"`
	Mediaids       []string `json:"mediaids,omitempty"`
	Notes          string   `json:"notes,omitempty"`
	DeviceType     int      `json:"device_type" validate:"required"`
	Lat            int      `json:"lat,omitempty"`
	Lng            int      `json:"lng,omitempty"`
	DeviceDetail   string   `json:"device_detail" validate:"required"`
	Wifiname       string   `json:"wifiname,omitempty"`
	Wifimac        string   `json:"wifimac,omitempty"`
}

type Records struct {
	Records []CheckInRecord `json:"records" validate:"min=1,max=200"`
}

func (w *weWork) AddCheckInRecord(corpId uint, records Records) (resp internal.BizResponse) {
	if ok := validate.Struct(records); ok != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = ok.Error()
		return
	}
	_, err := w.getRequest(corpId).SetBody(records).
		SetResult(&resp).Post("/cgi-bin/checkin/add_checkin_record")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}
