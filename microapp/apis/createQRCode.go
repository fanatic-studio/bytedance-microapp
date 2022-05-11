package apis

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/fanatic-studio/bytedance-microapp/microapp/utils"
	"github.com/fanatic-studio/bytedance-microapp/microapp/utils/request"
	"io/ioutil"
	"strings"
)

const (
	createQRCodeURL    = "https://developer.toutiao.com/api/apps/qrcode"
	AppNameToutiao     = "toutiao"
	AppNameToutiaoLite = "toutiao_lite"
	AppNameDouyin      = "douyin"
	AppNameDouyinLite  = "douyin_lite"
	AppNamePipixia     = "pipixia"
	AppNameHuoshan     = "huoshan"
	AppNameXigua       = "xigua"
)

// QRCodeParams 获取小程序/小游戏的二维码的请求参数
type QRCodeParams struct {
	AccessToken string `json:"access_token"`

	// 是打开二维码的字节系 app 名称，默认为今日头条，取值如下表所示
	// | appname | 对应字节系 app |
	// | ------- | ------------  |
	// | toutiao | 今日头条       |
	// | douyin  | 抖音          |
	// | pipixia | 皮皮虾         |
	// | huoshan | 火山小视频      |
	Appname string `json:"appname,omitempty"`

	// 小程序/小游戏启动参数，小程序则格式为 encode({path}?{query})，
	// 小游戏则格式为 JSON 字符串，默认为空
	Path string `json:"path,omitempty"`

	// 是否展示小程序/小游戏 icon，默认不展示
	SetIcon bool `json:"set_icon,omitempty"`
}

// CreateQRCode 获取 AccessToken
// 获取小程序/小游戏的二维码。该二维码可通过任意 app 扫码打开，能跳转到开发者指定的对应字节系 app 内拉起小程序/小游戏， 并传入开发者指定的参数。通过该接口生成的二维码，永久有效，暂无数量限制。
// see https://microapp.bytedance.com/docs/zh-CN/mini-app/develop/server/qr-code/create-qr-code
func CreateQRCode(accessToken, appname, path string) (response string, err error) {
	params := QRCodeParams{
		AccessToken: accessToken,
		Appname:     appname,
		Path:        path,
		SetIcon:     true,
	}
	resp, err := request.PostJSON(createQRCodeURL, params)
	if err != nil {
		return
	}
	contentType := resp.Header.Get("Content-Type")
	if strings.HasPrefix(contentType, "application/json") {
		var body []byte
		body, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			return
		}

		result := &utils.Error{}
		err = json.Unmarshal(body, result)
		if err != nil {
			return
		}
		if result.Code != 0 {
			err = fmt.Errorf("fetchQrCode error : errcode=%v , errmsg=%v", result.Code, result.Msg)
			return
		}
	} else if contentType == "image/jpeg" || contentType == "image/png" {
		res, _ := ioutil.ReadAll(resp.Body)
		response = base64.StdEncoding.EncodeToString(res)
		return
	} else {
		err = fmt.Errorf("fetchQrCode error : unknown response content type - %v", contentType)
		return
	}

	return
}
