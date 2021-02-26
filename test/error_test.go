package dingtalktest

import (
	"testing"

	"dingtalk"
)

// TestUnknownMessageType ...
func TestUnknownMessageType(t *testing.T) {
	if err := s.Send("unknown message"); err != nil {
		if dingtalk.IsSevereError(err) {
			return
		}
		t.Fatal(err)
	}
}

// TestEmptyText ...
func TestEmptyText(t *testing.T) {
	msg := dingtalk.NewText()
	if err := s.Send(msg); err != nil {
		if dingtalk.IsSevereError(err) {
			return
		}
		t.Fatal(err)
	}
}

// TestEmptyTittle ...
func TestEmptyTittle(t *testing.T) {
	text := `
		# 好时代！ 
		## 来临罢！
	`
	msg := dingtalk.NewMarkdown().SetMarkdown("", text)
	if err := s.Send(msg); err != nil {
		if dingtalk.IsSevereError(err) {
			return
		}
		t.Fatal(err)
	}
}

// TestEmptyContent ...
func TestEmptyContent(t *testing.T) {
	title := `homo兴，野兽王！`
	msg := dingtalk.NewMarkdown().SetMarkdown(title, "")
	if err := s.Send(msg); err != nil {
		if dingtalk.IsSevereError(err) {
			return
		}
		t.Fatal(err)
	}
}

// TestSendTooFast ...
func TestSendTooFast(t *testing.T) {
	msg := dingtalk.NewText().SetText("哼哼哼，啊啊啊啊啊啊啊！啊啊啊啊啊啊啊啊啊啊啊啊！")
	for i := 0; i < 30; i++ {
		if err := s.Send(msg); err != nil {
			if dingtalk.IsRepeatableError(err) {
				return
			}
			t.Fatal(err)
		}
	}
}

// TestTextTooLarge ...
func TestTextTooLarge(t *testing.T) {
	largeText := `
		哼哼哼，啊啊啊啊啊啊啊！啊啊啊啊啊啊啊啊啊啊啊啊！
	`
	for i := 0; i < 10000; i++ {
		largeText += `哼哼哼，啊啊啊啊啊啊啊！啊啊啊啊啊啊啊啊啊啊啊啊！`
	}
	msg := dingtalk.NewText().SetText(largeText)
	if err := s.Send(msg); err != nil {
		if dingtalk.IsSevereError(err) {
			return
		}
		t.Fatal(err)
	}
}
