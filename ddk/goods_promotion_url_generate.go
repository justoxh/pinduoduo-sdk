package ddk

import (
	"github.com/justoxh/pinduoduo-sdk/common"
	"github.com/justoxh/pinduoduo-sdk/util"
)

/**
生成多多进宝商品推广链接
*/
type GoodsPromotionUrlGenerateParams struct {
	PID                  string  `json:"p_id"`                             //推广位ID
	GoodsIdList          []int64 `json:"goods_id_list"`                    //商品ID，仅支持单个查询
	GenerateShortUrl     bool    `json:"generate_short_url"`               //是否生成短链接，true-是，false-否
	MultiGroup           *bool   `json:"multi_group,omitempty"`            //true--生成多人团推广链接 false--生成单人团推广链接（默认false）1、单人团推广链接：用户访问单人团推广链接，可直接购买商品无需拼团。2、多人团推广链接：用户访问双人团推广链接开团，若用户分享给他人参团，则开团者和参团者的佣金均结算给推手
	CustomParameters     *string `json:"custom_parameters,omitempty"`      //自定义参数，为链接打上自定义标签。自定义参数最长限制64个字节。
	PullNew              *bool   `json:"pull_new,omitempty"`               //是否开启订单拉新，true表示开启（订单拉新奖励特权仅支持白名单，请联系工作人员开通）
	GenerateWeappWebview *bool   `json:"generate_weapp_webview,omitempty"` //是否生成唤起微信客户端链接，true-是，false-否，默认false
	ZsDuoID              *int64  `json:"zs_duo_id,omitempty"`              //招商多多客ID
	GenerateWeApp        *bool   `json:"generate_we_app"`                  //是否生成小程序推广
}

type GoodsPromotionUrlGenerateResult struct {
	GoodsPromotionUrlGenerateResponse struct {
		GoodsPromotionUrlList []GoodsPromotionUrlGenerateInfo `json:"goods_promotion_url_list"`
	} `json:"goods_promotion_url_generate_response"`
	common.CommonResult
}

type GoodsPromotionUrlGenerateInfo struct {
	WeAppWebViewShortUrl string `json:"we_app_web_view_short_url"`
	WeAppWebViewUrl      string `json:"we_app_web_view_url"`
	MobileShortUrl       string `json:"mobile_short_url"`
	MobileUrl            string `json:"mobile_url"`
	ShortUrl             string `json:"short_url"`
	Url                  string `json:"url"`
	WeAppInfo            struct {
		WeAppIconUrl      string `json:"we_app_icon_url"`
		BannerUrl         string `json:"banner_url"`
		Desc              string `json:"desc"`
		SourceDisplayName string `json:"source_display_name"`
		PagePath          string `json:"page_path"`
		UserName          string `json:"user_name"`
		Title             string `json:"title"`
		AppId             string `json:"app_id"`
	} `json:"we_app_info"`
}

func (this *DuoduoKe) GoodsPromotionUrlGenerate(p *GoodsPromotionUrlGenerateParams) (*GoodsPromotionUrlGenerateResult, error) {
	apiType := `pdd.ddk.goods.promotion.url.generate`
	params, paramsURL := util.FormatURLParams(p)
	url := this.GetURL(apiType, "", params, paramsURL)

	var result GoodsPromotionUrlGenerateResult
	err := util.HttpPOST(url, nil, &result)
	if err != nil {
		return nil, err
	}
	err = common.CheckErrCode(result.CommonResult)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
