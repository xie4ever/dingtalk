package dingtalk

import "strings"

// Send 发送消息
// message: 消息模板
func (s *Sender) Send(message interface{}) error {
	switch message.(type) {
	case textMessage:
		// 校验content
		if strings.TrimSpace(message.(textMessage).Text.Content) == "" {
			return emptyContent
		}
		// 校验cast
		if s.cast == nil {
			return castIsNil
		}
		return s.post(s.cast, message)

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
		if s.cast == nil {
			return castIsNil
		}
		return s.post(s.cast, message)

	default:
		return unknownMessageType
	}
}
