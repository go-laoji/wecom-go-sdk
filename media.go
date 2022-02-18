package wework

import (
	"encoding/json"
	"fmt"
	"github.com/go-laoji/wecom-go-sdk/internal"
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
func (ww weWork) MediaUploadAttachment(corpId uint, attrs Media) (resp MediaUploadResponse) {
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
	queryParams := ww.buildCorpQueryToken(corpId)
	queryParams.Add("media_type", string(attrs.Type))
	queryParams.Add("attachment_type", fmt.Sprintf("%v", attrs.AttachmentType))
	body, err := internal.HttpUploadMedia(
		fmt.Sprintf("/cgi-bin/media/upload_attachment?%s",
			queryParams.Encode()), attrs.FilePath, "")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	} else {
		json.Unmarshal(body, &resp)
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
func (ww weWork) MediaUpload(corpId uint, fileType MediaType, filePath string, fileName string) (resp MediaUploadResponse) {
	if !isExists(filePath) {
		resp.ErrCode = 500
		resp.ErrorMsg = "文件路径不存在"
		return
	}
	queryParams := ww.buildCorpQueryToken(corpId)
	queryParams.Add("type", string(fileType))

	body, err := internal.HttpUploadMedia(
		fmt.Sprintf("/cgi-bin/media/upload?%s",
			queryParams.Encode()), filePath, fileName)
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	} else {
		json.Unmarshal(body, &resp)
	}
	return
}

type MediaUploadImgResponse struct {
	internal.BizResponse
	Url string `json:"url"`
}

// MediaUploadImg 上传图片
// https://open.work.weixin.qq.com/api/doc/90001/90143/90392
func (ww weWork) MediaUploadImg(corpId uint, filePath string, fileName string) (resp MediaUploadImgResponse) {
	if !isExists(filePath) {
		resp.ErrCode = 500
		resp.ErrorMsg = "文件路径不存在"
		return
	}
	queryParams := ww.buildCorpQueryToken(corpId)
	body, err := internal.HttpUploadMedia(
		fmt.Sprintf("/cgi-bin/media/uploadimg?%s",
			queryParams.Encode()), filePath, fileName)
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	} else {
		json.Unmarshal(body, &resp)
	}
	return
}
