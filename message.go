package dingtalk

const (
	msgTypeText     = "text"
	msgTypeMarkdown = "markdown"
)

type textMessage struct {
	MsgType string `json:"msgtype"`
	Text    text   `json:"text"`
	At      at     `json:"at"`
}

type text struct {
	Content string `json:"content"`
}

type at struct {
	AtMobiles []string `json:"atMobiles"`
	IsAtAll   bool     `json:"isAtAll"`
}

// NewText ...
func NewText() textMessage {
	return textMessage{
		MsgType: msgTypeText,
		Text:    text{},
		At:      at{},
	}
}

// SetText ...
func (m textMessage) SetText(content string) textMessage {
	m.Text.Content = content
	return m
}

// SetAt ...
func (m textMessage) SetAt(all bool, mobiles ...string) textMessage {
	m.At = at{
		AtMobiles: mobiles,
		IsAtAll:   all,
	}
	return m
}

type markdown struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

type markdownMessage struct {
	MsgType  string   `json:"msgtype"`
	Markdown markdown `json:"markdown"`
	At       at       `json:"at"`
}

// NewMarkdown ...
func NewMarkdown() *markdownMessage {
	return &markdownMessage{
		MsgType:  msgTypeMarkdown,
		Markdown: markdown{},
		At:       at{},
	}
}

// SetMarkdown ...
func (m markdownMessage) SetMarkdown(title, text string) markdownMessage {
	m.Markdown = markdown{
		Title: title,
		Text:  text,
	}
	return m
}

// SetAt ...
func (m markdownMessage) SetAt(all bool, mobiles ...string) markdownMessage {
	m.At = at{
		AtMobiles: mobiles,
		IsAtAll:   all,
	}
	return m
}
