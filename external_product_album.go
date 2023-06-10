package wework

import (
	"github.com/go-laoji/wecom-go-sdk/v2/internal"
)

type Product struct {
	Description string `json:"description" validate:"required,max=300"`
	Price       int    `json:"price" validate:"required,max=5000000"`
	ProductSn   string `json:"product_sn" validate:"omitempty,max=128,"`
	Attachments []struct {
		Type  string `json:"type"`
		Image struct {
			MediaID string `json:"media_id"`
		} `json:"image"`
	} `json:"attachments" validate:"required"`
}
type AddProductAlbumResponse struct {
	internal.BizResponse
	ProductId string `json:"product_id"`
}

// AddProductAlbum 创建商品图册
// https://open.work.weixin.qq.com/api/doc/90001/90143/95131#%E5%88%9B%E5%BB%BA%E5%95%86%E5%93%81%E5%9B%BE%E5%86%8C
func (ww *weWork) AddProductAlbum(corpId uint, product Product) (resp AddProductAlbumResponse) {
	if ok := validate.Struct(product); ok != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = ok.Error()
	}
	_, err := ww.getRequest(corpId).SetBody(product).SetResult(&resp).
		Post("/cgi-bin/externalcontact/add_product_album")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}

type GetProductAlbumResponse struct {
	internal.BizResponse
	Product struct {
		Product
		ProductId  string `json:"product_id"`
		CreateTime int    `json:"create_time"`
	} `json:"product"`
}

// GetProductAlbum 获取商品图册
// https://open.work.weixin.qq.com/api/doc/90001/90143/95131#%E8%8E%B7%E5%8F%96%E5%95%86%E5%93%81%E5%9B%BE%E5%86%8C
func (ww *weWork) GetProductAlbum(corpId uint, productId string) (resp GetProductAlbumResponse) {
	h := H{}
	h["product_id"] = productId
	_, err := ww.getRequest(corpId).SetBody(h).SetResult(&resp).
		Post("/cgi-bin/externalcontact/get_product_album")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}

type GetProductAlbumListResponse struct {
	internal.BizResponse
	NextCursor  string `json:"next_cursor"`
	ProductList []struct {
		ProductId string `json:"product_id"`
		Product
	} `json:"product_list"`
}

// GetProductAlbumList 获取商品图册列表
// https://open.work.weixin.qq.com/api/doc/90001/90143/95131#%E8%8E%B7%E5%8F%96%E5%95%86%E5%93%81%E5%9B%BE%E5%86%8C%E5%88%97%E8%A1%A8
func (ww *weWork) GetProductAlbumList(corpId uint, limit int, cursor string) (resp GetProductAlbumListResponse) {
	h := H{}
	h["limit"] = limit
	h["cursor"] = cursor
	_, err := ww.getRequest(corpId).SetBody(h).SetResult(&resp).
		Post("/cgi-bin/externalcontact/get_product_album_list")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}

type ProductUpdateRequest struct {
	ProductId   string `json:"product_id" validate:"required"`
	Description string `json:"description,omitempty" validate:"omitempty,max=300"`
	Price       int    `json:"price,omitempty" validate:"omitempty,max=5000000"`
	ProductSn   string `json:"product_sn,omitempty" validate:"omitempty,max=128,"`
	Attachments []struct {
		Type  string `json:"type"`
		Image struct {
			MediaID string `json:"media_id"`
		} `json:"image"`
	} `json:"attachments,omitempty" validate:"required"`
}

// UpdateProductAlbum 编辑商品图册
// https://open.work.weixin.qq.com/api/doc/90001/90143/95131#%E7%BC%96%E8%BE%91%E5%95%86%E5%93%81%E5%9B%BE%E5%86%8C
func (ww *weWork) UpdateProductAlbum(corpId uint, request ProductUpdateRequest) (resp internal.BizResponse) {
	if ok := validate.Struct(request); ok != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = ok.Error()
	}
	_, err := ww.getRequest(corpId).SetBody(request).SetResult(&resp).
		Post("/cgi-bin/externalcontact/update_product_album")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}

// DeleteProductAlbum 删除商品图册
// https://open.work.weixin.qq.com/api/doc/90001/90143/95131#%E5%88%A0%E9%99%A4%E5%95%86%E5%93%81%E5%9B%BE%E5%86%8C
func (ww *weWork) DeleteProductAlbum(corpId uint, productId string) (resp internal.BizResponse) {
	h := H{}
	h["product_id"] = productId
	_, err := ww.getRequest(corpId).SetBody(h).SetResult(&resp).
		Post("/cgi-bin/externalcontact/delete_product_album")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}
