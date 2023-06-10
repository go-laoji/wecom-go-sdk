package wework

import (
	"fmt"
	"github.com/go-laoji/wecom-go-sdk/v2/internal"
	"os"
)

type MediaType string

const (
	MediaImage MediaType = "image"
	MediaVoice MediaType = "voice"
	MediaVideo MediaType = "video"
	MediaFile  MediaType = "file"
)

type Media struct {
	Type           MediaType `json:"type" validate:"required,oneof=image voice video file"`
	AttachmentType int       `json:"attachment_type" validate:"required,oneof=1 2"`
	FilePath       string    `json:"file_path" validate:"required"`
	FileName       string    `json:"file_name"`
}

type MediaUploadResponse struct {
	internal.BizResponse
	Type     string `json:"type"`
	MediaId  string `json:"media_id"`
	CreateAt uint64 `json:"create_at"`
}

// MediaUploadAttachment 上传附件资源
// 不同的附件类型用于不同的场景。1：朋友圈；2:商品图册
// https://open.work.weixin.qq.com/api/doc/90001/90143/95178
func (ww *weWork) MediaUploadAttachment(corpId uint, attrs Media) (resp MediaUploadResponse) {
	if ok := validate.Struct(attrs); ok != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = ok.Error()
		return
	}
	if !isExists(attrs.FilePath) {
		resp.ErrCode = 500
		resp.ErrorMsg = fmt.Sprintf("%s 文件不存在！", attrs.FilePath)
		return
	}
	_, err := ww.getRequest(corpId).SetQueryParam("media_type", string(attrs.Type)).
		SetQueryParam("attachment_type", fmt.Sprintf("%v", attrs.AttachmentType)).
		SetFile("media", attrs.FilePath).SetResult(&resp).
		Post("/cgi-bin/media/upload_attachment")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}

func isExists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

// MediaUpload 上传临时素材
// https://open.work.weixin.qq.com/api/doc/90001/90143/90389
func (ww *weWork) MediaUpload(corpId uint, fileType MediaType, filePath string) (resp MediaUploadResponse) {
	if !isExists(filePath) {
		resp.ErrCode = 500
		resp.ErrorMsg = "文件路径不存在"
		return
	}
	_, err := ww.getRequest(corpId).SetQueryParam("type", string(fileType)).
		SetFile("media", filePath).SetResult(&resp).
		Post("/cgi-bin/media/upload")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}

type MediaUploadImgResponse struct {
	internal.BizResponse
	Url string `json:"url"`
}

// MediaUploadImg 上传图片
// https://open.work.weixin.qq.com/api/doc/90001/90143/90392
func (ww *weWork) MediaUploadImg(corpId uint, filePath string) (resp MediaUploadImgResponse) {
	if !isExists(filePath) {
		resp.ErrCode = 500
		resp.ErrorMsg = "文件路径不存在"
		return
	}
	_, err := ww.getRequest(corpId).
		SetFile("media", filePath).SetResult(&resp).
		Post("/cgi-bin/media/uploadimg")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}
