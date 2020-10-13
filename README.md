# wechat-work-message-push-go

### 构建镜像
1. 编译生成可执行文件
2. 构建
Dockfile
```
FROM alpine:latest
ADD wechat-work-message-push-go /app/wechat-work-message-push-go
CMD ["/app/wechat-work-message-push-go"]
```

### 使用：
1.创建企业号（200人以下不需要认证），获取 `WechatWorkCorpId` https://work.weixin.qq.com/

2.创建自建应用，获取 `WechatWorkAgentId` 和 `WechatWorkCorpSecret`

3.到通讯录查看自己的账号，获取 `DefaultReceiverUserId`

3.启动容器
环境变量的说明
```
DefaultReceiverUserId 默认发送到的user ID
Token 发送POST请求到/push 时用于验证身份
GrafanaWebhookUser与GrafanaWebhookPassword 用于验证，grafana页面中的用户密码项
GrafanaWebhookResutURL grafana不返回图片地址时用以替代返回图片地址（grafana里面需要开启provider = local让grafana生成图片存储在本地）
```

```
docker run -d \
  --name webhook \
  --restart always \
  -p 192.168.1.120:5555:5555 \
  -v /usr/share/zoneinfo/Asia/Shanghai:/etc/localtime \
  -e Token="Token" \
  -e WechatWorkCorpSecret="WechatWorkCorpSecret" \
  -e WechatWorkCorpId="WechatWorkCorpId" \
  -e DefaultReceiverUserId="@all" \
  -e WechatWorkAgentId="WechatWorkAgentId" \
  -e GrafanaWebhookUser="admin" \
  -e GrafanaWebhookPassword="admin" \
  -e GrafanaWebhookResutURL="http://192.168.1.120:3000/" \
  wechat-work-message-push-go
```

4.到 企业微信->我的企业->微信插件，扫二维码关注企业微信 https://work.weixin.qq.com/wework_admin/frame#profile/wxPlugin

5.发送以下请求
```shell
curl -X POST \
  http://192.168.1.120:5555/push \
  -H 'Content-Type: application/x-www-form-urlencoded' \
  -H 'token: Token' \
  -d 'message=1234'
```

6.刚才在微信上的关注的微工作台应收到第五步发送的消息
![](https://github.com/cloverzrg/wechat-work-message-push-go/raw/master/IMG_8017.jpg)


7.grafana 报警通知功能
设置GrafanaWebhookUser和GrafanaWebhookPassword两个环境变量就可以用了
![](https://github.com/cloverzrg/wechat-work-message-push-go/raw/master/grafana_webhook.png)