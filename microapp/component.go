package microapp

import "github.com/fanatic-studio/bytedance-microapp/microapp/apis"

type Component struct {
	config *config
}

func newComponent(config *config) *Component {
	comp := &Component{}
	comp.config = config
	return comp
}

// GetAccessToken 获取 AccessToken 支付时使用
func (c *Component) GetAccessToken() apis.GetAccessTokenResp {
	return apis.GetAccessToken(c.config.AppId, c.config.AppSecret)
}

// CreateQRCode 创建跳转二维码
func (c *Component) CreateQRCode(accessToken, appname, path string) (response string, err error) {
	return apis.CreateQRCode(accessToken, appname, path)
}

// Code2Session 不需要授权获取用户 openid 和 sessionkey
func (c *Component) Code2Session(code string, anonymousCode string) apis.Code2SessionResp {
	return apis.Code2Session(c.config.AppId, c.config.AppSecret, code, anonymousCode)
}

// Decrypt 解密用户授权登陆的数据
// https://microapp.bytedance.com/docs/zh-CN/mini-app/develop/api/open-interface/user-information/sensitive-data-process/
func (c *Component) Decrypt(encryptedData, sessionKey, iv string) apis.DecryptUserInfoResp {
	return apis.DecryptUserInfo(encryptedData, sessionKey, iv)
}
