package qyapi

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/cloverzrg/wechat-work-message-push-go/config"
	"github.com/cloverzrg/wechat-work-message-push-go/grafana"
	"github.com/cloverzrg/wechat-work-message-push-go/logger"
)

// SendMessage 给 /push 接口使用，发送普通文本消息
func SendMessage(content string, toUser string) (err error) {
	if content == "" {
		content = "content 为空"
	}
	logger.Infof("push message: %s\n", content)
	token, err := GetToken()
	if err != nil {
		logger.Error(err)
		return err
	}
	url := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=%s", token)
	m := TextMessage{
		Agentid: config.Config.WechatWork.AgentID,
		Msgtype: "text",
	}
	if len(toUser) == 0 {
		m.Touser = config.Config.WechatWork.DefaultReceiverUserID
	} else {
		m.Touser = toUser
	}

	m.Text.Content = content

	jsonStr, err := json.Marshal(m)
	if err != nil {
		logger.Error("sendMessage error:%s", err)
		return err
	}
	postJSON(url, jsonStr)
	return err
}

// SendTextGrafanaMessage 给/grafana 使用，发送grafana消息
func SendTextGrafanaMessage(noti grafana.Notification) (err error) {
	message, title, imageURL, toUser, ErrList := noti.Message, noti.Title, noti.ImageUrl, "", ""
	now := time.Now()
	CurrentTime := fmt.Sprintf("%02d:%02d:%02d", now.Hour(), now.Minute(), now.Second())

	if imageURL == "" {
		imageURL = config.Config.DefaultResutURL
	}

	if message == "" {
		message = "message 为空"
	}
	if title == "" {
		title = "title 为空"
	}

	if len(noti.EvalMatches) == 1 {
		ErrList = noti.EvalMatches[0].Metric
	} else {
		for _, v := range noti.EvalMatches {
			ErrList = ErrList + "\n" + v.Metric
		}
	}

	var content string
	if noti.State == "ok" {
		content = fmt.Sprintf(`%v
故障恢复详情
详  情: %v

恢复时间：%v

grafana页面：<a href="%v">点击</a>`, title, message, CurrentTime, imageURL)
	} else {
		content = fmt.Sprintf(`%v
故障详情
详  情: %v

故障时间：%v
故障服务列表：%v

grafana页面：<a href="%v">点击</a>`, title, message, CurrentTime, ErrList, imageURL)
	}

	// logger.Infof("push message: %s\n", content)
	token, err := GetToken()
	if err != nil {
		return err
	}
	url := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=%s", token)
	m := TextMessage{
		Msgtype: "text",
		Agentid: config.Config.WechatWork.AgentID,
	}

	if len(toUser) == 0 {
		m.Touser = config.Config.WechatWork.DefaultReceiverUserID
	} else {
		m.Touser = toUser
	}

	m.Text.Content = content

	jsonStr, err := json.Marshal(m)

	if err != nil {
		logger.Error("sendMessage error:%s", err)
		SendMessage(err.Error(), "")
		return err
	}
	postJSON(url, jsonStr)
	return err
}
