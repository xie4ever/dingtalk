package dingtalk

import (
	"strings"
	"time"

	"github.com/xiaojiaoyu100/cast"
)

// Group ...
type Group struct {
	webHook   string
	secretKey string
	cast      *cast.Cast
}

// InitDingTalkGroup 初始化服务
// webHook:		dingtalk自定义机器人的群webHook
// secretKey:	dingtalk的自定义机器人的群密钥
func InitDingTalkGroup(webHook, secretKey string) (*Group, error) {
	c, err := cast.New(
		cast.WithBaseURL(webHook),
		cast.AddCircuitConfig(defaultCircuitName),
		cast.WithDefaultCircuit(defaultCircuitName),
		cast.WithRetry(3),
		cast.WithExponentialBackoffDecorrelatedJitterStrategy(
			time.Millisecond*200,
			time.Millisecond*500,
		),
	)

	g := &Group{
		webHook:   webHook,
		secretKey: secretKey,
		cast:      c,
	}
	return g, err
}

// SendMessage 发送消息
// message:    消息模板
func (g *Group) SendMessage(message interface{}) error {
	switch message.(type) {
	case textMessage:
		// 校验content
		if strings.TrimSpace(message.(textMessage).Text.Content) == "" {
			return emptyContent
		}
		// 校验cast
		if g.cast == nil {
			return castIsNil
		}
		return g.post(g.cast, message)

	case markdownMessage:
		// 校验title
		if strings.TrimSpace(message.(markdownMessage).Markdown.Title) == "" {
			return emptyTittle
		}
		// 校验text
		if strings.TrimSpace(message.(markdownMessage).Markdown.Text) == "" {
			return emptyText
		}
		// 校验cast
		if g.cast == nil {
			return castIsNil
		}
		return g.post(g.cast, message)

	default:
		return unknownMessageType
	}
}
