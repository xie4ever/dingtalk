# dingtalk
dingtalk-sdk

## 1.dingtalk doc

base on https://developers.dingtalk.com/document/app/custom-robot-access?spm=ding_open_doc.document.0.0.6d9d28e1xy0yjw#topic-2026027

## 2.all you need

1. create your dingtalk group.
2. create a robot.
3. get your param:
    - web hook
    - secret key

## 3.example

```golang
package dingtalktest

import (
	"testing"

	"dingtalk"
)

// 钉钉开发文档：https://developers.dingtalk.com/document/app/custom-robot-access?spm=ding_open_doc.document.0.0.6d9d28e1xy0yjw#topic-2026027
const (
	webHook   = "your web hook"   // 分配给应用的WebHook（钉钉群机器人处获取）
	secretKey = "your secret key" // 分配给应用的SecretKey（钉钉群机器人处获取）
)

var g *dingtalk.Group

func init() {
	g, _ = dingtalk.InitDingTalkGroup(webHook, secretKey)
}

// TestSendText ...
func TestSendText(t *testing.T) {
	msg := dingtalk.NewText().SetText("哼哼哼，啊啊啊啊啊啊啊！啊啊啊啊啊啊啊啊啊啊啊啊！")
	err := g.SendMessage(msg)
	if err != nil {
		t.Fatal(err)
	}
}

// TestSendMarkdown ...
func TestSendMarkdown(t *testing.T) {
	title := `员工岗位变更通知`
	text := `
		# 世界级美声
		* 哼！
		* 哼！
		* 哼！
		* 啊啊啊啊啊啊！
		* 啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊！
	`
	msg := dingtalk.NewMarkdown().SetMarkdown(title, text)
	err := g.SendMessage(msg)
	if err != nil {
		t.Fatal(err)
	}
}
```
