package config

import "os"

type jsonConfig struct {
	Token      string
	WechatWork struct {
		DefaultReceiverUserID string // 默认发送到的user ID
		CorpID                string // 企业ID
		CorpSecret            string // 应用secret
		AgentID               string // 应用ID

	}
	GrafanaWebhookUser     string
	GrafanaWebhookPassword string
	DefaultResutURL        string // grafana不返回图片链接时的默认链接
}

// Config 用于全局共享配置
var Config jsonConfig

// LoadConfig 载入环境变量中的配置
func LoadConfig() (err error) {
	Config.Token = os.Getenv("Token")
	Config.WechatWork.CorpSecret = os.Getenv("WechatWorkCorpSecret")
	Config.WechatWork.CorpID = os.Getenv("WechatWorkCorpId")
	Config.WechatWork.DefaultReceiverUserID = os.Getenv("DefaultReceiverUserId")
	Config.WechatWork.AgentID = os.Getenv("WechatWorkAgentId")
	Config.DefaultResutURL = os.Getenv("GrafanaWebhookResutURL")
	Config.GrafanaWebhookUser = os.Getenv("GrafanaWebhookUser")
	Config.GrafanaWebhookPassword = os.Getenv("GrafanaWebhookPassword")

	// fmt.Printf("%+v\n", Config)
	return err
}
