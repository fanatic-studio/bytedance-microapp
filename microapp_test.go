package bytedance_microapp

import (
	"github.com/cute-angelia/go-utils/syntax/ijson"
	"github.com/cute-angelia/go-utils/utils/conf"
	"github.com/fanatic-studio/bytedance-microapp/microapp"
	"log"
	"testing"
)

func getApp() *microapp.Component {
	conf.LoadConfigFile("./config.toml")
	return microapp.Load("microapp")
}

func TestGetAccessToken(t *testing.T) {
	microApp := getApp()
	accessTokenInfo := microApp.GetAccessToken()
	log.Println(ijson.Pretty(accessTokenInfo))
}

func TestCreateQRCode(t *testing.T) {
	microApp := getApp()
	var accessToken = "080112184675337539513443336d324b7456305848776a4b38413d3d"
	accessTokenInfo, _ := microApp.CreateQRCode(accessToken, "douyin", "accountId:123123")
	log.Println(ijson.Pretty(accessTokenInfo))
}
