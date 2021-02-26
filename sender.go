package dingtalk

import (
	"time"

	"github.com/xiaojiaoyu100/cast"
)

// Sender ...
type Sender struct {
	webHook   string
	secretKey string
	cast      *cast.Cast
}

// NewSender 初始化发送者
// webHook: dingtalk自定义机器人的群webHook
// secretKey: dingtalk的自定义机器人的群密钥
func NewSender(webHook, secretKey string) (*Sender, error) {
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

	g := &Sender{
		webHook:   webHook,
		secretKey: secretKey,
		cast:      c,
	}
	return g, err
}
