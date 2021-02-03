package dingtalk

type dtError string

func (e dtError) Error() string {
	return string(e)
}

// DIY error
const (
	castRequestFailed  = dtError("cast request failed")  // http请求失败
	unknownMessageType = dtError("unknown message type") // 未知的模板类型
	emptyText          = dtError("empty text")           // 发送markdown空text
	emptyTittle        = dtError("empty title")          // 发送markdown空title
	emptyContent       = dtError("empty content")        // 发送文案空内容
	castIsNil          = dtError("cast is nil")          // 空cast
)

// DingTalk error
const (
	sendTooFast      = dtError("send too fast, exceed 20 times per minute") // 发送信息频率过快（1min > 20）
	textTooLong      = dtError("message too long, exceed 20000 bytes")      // 发送内容太长
	invalidTimestamp = dtError("invalid timestamp")                         // 非法时间戳
)

// IsSlightError 是否轻微错误（不建议报警）
func IsSlightError(err error) bool {
	switch err.Error() {
	case castRequestFailed.Error():
		return true
	default:
		return false
	}
}

// IsRepeatableError 是否可重试错误（不建议报警，建议重试）
func IsRepeatableError(err error) bool {
	switch err.Error() {
	case castRequestFailed.Error(), sendTooFast.Error():
		return true
	default:
		return false
	}
}

// IsSevereError 是否严重错误（建议报警）
func IsSevereError(err error) bool {
	return !IsSlightError(err) && !IsRepeatableError(err)
}
